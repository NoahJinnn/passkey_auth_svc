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

const (
	SyncType          = "sync"
	SaveCsFailedType  = "save_cs_failed"
	QueryCsType       = "query_cs"
	QueryCsFailedType = "query_cs_failed"
)

type SyncPayload struct {
	DbVersion int32  `json:"db_version"`
	SiteId    string `json:"site_id"`
	ChangeSet string `json:"change_set"`
}

type QueryCsPayload struct {
	DbVersion int32  `json:"db_version"`
	SiteId    string `json:"site_id"`
}
