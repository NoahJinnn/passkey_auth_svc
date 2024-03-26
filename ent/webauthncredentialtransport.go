// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredential"
	"github.com/NoahJinnn/passkey_auth_svc/ent/webauthncredentialtransport"
)

// WebauthnCredentialTransport is the model entity for the WebauthnCredentialTransport schema.
type WebauthnCredentialTransport struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// WebauthnCredentialID holds the value of the "webauthn_credential_id" field.
	WebauthnCredentialID string `json:"webauthn_credential_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebauthnCredentialTransportQuery when eager-loading is set.
	Edges        WebauthnCredentialTransportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WebauthnCredentialTransportEdges holds the relations/edges for other nodes in the graph.
type WebauthnCredentialTransportEdges struct {
	// WebauthnCredential holds the value of the webauthn_credential edge.
	WebauthnCredential *WebauthnCredential `json:"webauthn_credential,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WebauthnCredentialOrErr returns the WebauthnCredential value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebauthnCredentialTransportEdges) WebauthnCredentialOrErr() (*WebauthnCredential, error) {
	if e.loadedTypes[0] {
		if e.WebauthnCredential == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: webauthncredential.Label}
		}
		return e.WebauthnCredential, nil
	}
	return nil, &NotLoadedError{edge: "webauthn_credential"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WebauthnCredentialTransport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case webauthncredentialtransport.FieldName, webauthncredentialtransport.FieldWebauthnCredentialID:
			values[i] = new(sql.NullString)
		case webauthncredentialtransport.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WebauthnCredentialTransport fields.
func (wct *WebauthnCredentialTransport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case webauthncredentialtransport.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wct.ID = *value
			}
		case webauthncredentialtransport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wct.Name = value.String
			}
		case webauthncredentialtransport.FieldWebauthnCredentialID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field webauthn_credential_id", values[i])
			} else if value.Valid {
				wct.WebauthnCredentialID = value.String
			}
		default:
			wct.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WebauthnCredentialTransport.
// This includes values selected through modifiers, order, etc.
func (wct *WebauthnCredentialTransport) Value(name string) (ent.Value, error) {
	return wct.selectValues.Get(name)
}

// QueryWebauthnCredential queries the "webauthn_credential" edge of the WebauthnCredentialTransport entity.
func (wct *WebauthnCredentialTransport) QueryWebauthnCredential() *WebauthnCredentialQuery {
	return NewWebauthnCredentialTransportClient(wct.config).QueryWebauthnCredential(wct)
}

// Update returns a builder for updating this WebauthnCredentialTransport.
// Note that you need to call WebauthnCredentialTransport.Unwrap() before calling this method if this WebauthnCredentialTransport
// was returned from a transaction, and the transaction was committed or rolled back.
func (wct *WebauthnCredentialTransport) Update() *WebauthnCredentialTransportUpdateOne {
	return NewWebauthnCredentialTransportClient(wct.config).UpdateOne(wct)
}

// Unwrap unwraps the WebauthnCredentialTransport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wct *WebauthnCredentialTransport) Unwrap() *WebauthnCredentialTransport {
	_tx, ok := wct.config.driver.(*txDriver)
	if !ok {
		panic("ent: WebauthnCredentialTransport is not a transactional entity")
	}
	wct.config.driver = _tx.drv
	return wct
}

// String implements the fmt.Stringer.
func (wct *WebauthnCredentialTransport) String() string {
	var builder strings.Builder
	builder.WriteString("WebauthnCredentialTransport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wct.ID))
	builder.WriteString("name=")
	builder.WriteString(wct.Name)
	builder.WriteString(", ")
	builder.WriteString("webauthn_credential_id=")
	builder.WriteString(wct.WebauthnCredentialID)
	builder.WriteByte(')')
	return builder.String()
}

// WebauthnCredentialTransports is a parsable slice of WebauthnCredentialTransport.
type WebauthnCredentialTransports []*WebauthnCredentialTransport
