package ws

import "encoding/json"

// Event is the Messages sent over the websocket
// Used to differ between different actions
type Event struct {
	// Type is the message type sent
	Type string `json:"type"`
	// Payload is the data Based on the Type
	Payload json.RawMessage `json:"payload"`
}

// EventHandler is a function signature that is used to affect messages on the socket and triggered
// depending on the type
type EventHandler func(event Event, c *Client) error

// AckPayload is the payload for the "ack" event
type AckPayload struct {
	// Message is the message sent
	Message string `json:"message"`
}

type SyncPayload struct {
	SiteId    string `json:"site_id"`
	ChangeSet string `json:"change_set"`
}
