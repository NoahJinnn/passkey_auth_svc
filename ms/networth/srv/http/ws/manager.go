package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/srv/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	clients ClientList

	// Using a syncMutex here to be able to lcok state before editing clients
	// Could also use Channels to block
	sync.RWMutex

	// handlers are functions that are used to handle Events
	handlers map[string]EventHandler
	srv      *handlers.HttpDeps
}

// NewManager is used to initalize all the values inside the manager
func NewManager(srv *handlers.HttpDeps) *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		srv:      srv,
	}
	m.setupEventHandlers()
	return m
}

// setupEventHandlers configures and adds all handlers
func (m *Manager) setupEventHandlers() {
	m.handlers[SyncType] = func(e Event, c *Client) error {
		fmt.Println("handle event type: ", e.Type)
		var syncP SyncPayload
		if err := json.Unmarshal(e.Payload, &syncP); err != nil {
			return fmt.Errorf("bad payload in ws msg: %v", err)
		}

		ctx := context.Background()
		csSvc := m.srv.Appl.GetChangesetSvc()
		csSvc.Upsert(ctx, c.userId, &ent.Changeset{
			SiteID:      syncP.SiteId,
			DbVersion:   syncP.DbVersion,
			FirstLaunch: true, // the first connection will change this to true
		})

		fmt.Printf("[WS-SYNC] saved changeset: %v\n", e.Type)
		for otherC := range m.clients[c.userId] {
			if c != otherC {
				otherC.egress <- e
			}
		}

		return nil
	}

	m.handlers[QueryCsType] = func(e Event, c *Client) error {
		fmt.Println("handle event type: ", e.Type)
		for otherC := range m.clients[c.userId] {
			if c != otherC {
				otherC.egress <- e
			}
		}
		return nil
	}

	m.handlers[ResponseCsType] = func(e Event, c *Client) error {
		fmt.Println("handle event type: ", e.Type)
		for otherC := range m.clients[c.userId] {
			if c != otherC {
				otherC.egress <- e
			}
		}
		return nil
	}
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (m *Manager) routeEvent(event Event, c *Client) error {
	// Check if Handler is present in Map
	if handler, ok := m.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) SyncBetweenUserDevices(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create New Client
	client := NewClient(userId, ws, m)
	m.addClient(userId, client)

	go client.readMessages()
	go client.writeMessages()

	return nil
}

// addClient will add clients to our clientList
func (m *Manager) addClient(usedId uuid.UUID, client *Client) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	if m.clients[usedId] == nil {
		m.clients[usedId] = make(map[*Client]bool)
	}
	m.clients[usedId][client] = true
}

// removeClient will remove the client and clean up
func (m *Manager) removeClient(usedId uuid.UUID, client *Client) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[usedId]; ok {
		clientSet := m.clients[usedId]
		if _, ok := clientSet[client]; ok {
			// close connection
			client.connection.Close()
			// remove
			delete(clientSet, client)
		}
	}
}
