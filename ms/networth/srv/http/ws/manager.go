package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
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
}

// NewManager is used to initalize all the values inside the manager
func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()
	return m
}

// setupEventHandlers configures and adds all handlers
func (m *Manager) setupEventHandlers() {
	m.handlers["sync"] = func(e Event, c *Client) error {
		fmt.Println("handle event type: ", e.Type)

		// Marshal Payload into wanted format
		var syncP SyncPayload
		if err := json.Unmarshal(e.Payload, &syncP); err != nil {
			return fmt.Errorf("bad payload in request: %v", err)
		}
		fmt.Println("sync payload: ", syncP)

		var ackP AckPayload
		ackP.Message = "ack success"
		data, err := json.Marshal(ackP)
		if err != nil {
			return fmt.Errorf("failed to marshal broadcast message: %v", err)
		}

		var outgoingEvent Event
		outgoingEvent.Type = "ack"
		outgoingEvent.Payload = data
		c.egress <- outgoingEvent

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

func (m *Manager) Sync(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	fmt.Println("New connection")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create New Client
	client := NewClient(userId.String(), ws, m)
	// Add the newly created client to the manager
	m.addClient(userId.String(), client)

	go client.readMessages()
	go client.writeMessages()

	return nil
}

// addClient will add clients to our clientList
func (m *Manager) addClient(usedId string, client *Client) {
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
func (m *Manager) removeClient(usedId string, client *Client) {
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
