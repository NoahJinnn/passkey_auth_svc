// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/fvsession"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Emails holds the value of the emails edge.
	Emails []*Email `json:"emails,omitempty"`
	// Passcodes holds the value of the passcodes edge.
	Passcodes []*Passcode `json:"passcodes,omitempty"`
	// WebauthnCredentials holds the value of the webauthn_credentials edge.
	WebauthnCredentials []*WebauthnCredential `json:"webauthn_credentials,omitempty"`
	// ItemTables holds the value of the item_tables edge.
	ItemTables []*ItemTable `json:"item_tables,omitempty"`
	// PrimaryEmail holds the value of the primary_email edge.
	PrimaryEmail *PrimaryEmail `json:"primary_email,omitempty"`
	// FvSession holds the value of the fv_session edge.
	FvSession *FvSession `json:"fv_session,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// EmailsOrErr returns the Emails value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EmailsOrErr() ([]*Email, error) {
	if e.loadedTypes[0] {
		return e.Emails, nil
	}
	return nil, &NotLoadedError{edge: "emails"}
}

// PasscodesOrErr returns the Passcodes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PasscodesOrErr() ([]*Passcode, error) {
	if e.loadedTypes[1] {
		return e.Passcodes, nil
	}
	return nil, &NotLoadedError{edge: "passcodes"}
}

// WebauthnCredentialsOrErr returns the WebauthnCredentials value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WebauthnCredentialsOrErr() ([]*WebauthnCredential, error) {
	if e.loadedTypes[2] {
		return e.WebauthnCredentials, nil
	}
	return nil, &NotLoadedError{edge: "webauthn_credentials"}
}

// ItemTablesOrErr returns the ItemTables value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ItemTablesOrErr() ([]*ItemTable, error) {
	if e.loadedTypes[3] {
		return e.ItemTables, nil
	}
	return nil, &NotLoadedError{edge: "item_tables"}
}

// PrimaryEmailOrErr returns the PrimaryEmail value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) PrimaryEmailOrErr() (*PrimaryEmail, error) {
	if e.loadedTypes[4] {
		if e.PrimaryEmail == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: primaryemail.Label}
		}
		return e.PrimaryEmail, nil
	}
	return nil, &NotLoadedError{edge: "primary_email"}
}

// FvSessionOrErr returns the FvSession value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) FvSessionOrErr() (*FvSession, error) {
	if e.loadedTypes[5] {
		if e.FvSession == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: fvsession.Label}
		}
		return e.FvSession, nil
	}
	return nil, &NotLoadedError{edge: "fv_session"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryEmails queries the "emails" edge of the User entity.
func (u *User) QueryEmails() *EmailQuery {
	return NewUserClient(u.config).QueryEmails(u)
}

// QueryPasscodes queries the "passcodes" edge of the User entity.
func (u *User) QueryPasscodes() *PasscodeQuery {
	return NewUserClient(u.config).QueryPasscodes(u)
}

// QueryWebauthnCredentials queries the "webauthn_credentials" edge of the User entity.
func (u *User) QueryWebauthnCredentials() *WebauthnCredentialQuery {
	return NewUserClient(u.config).QueryWebauthnCredentials(u)
}

// QueryItemTables queries the "item_tables" edge of the User entity.
func (u *User) QueryItemTables() *ItemTableQuery {
	return NewUserClient(u.config).QueryItemTables(u)
}

// QueryPrimaryEmail queries the "primary_email" edge of the User entity.
func (u *User) QueryPrimaryEmail() *PrimaryEmailQuery {
	return NewUserClient(u.config).QueryPrimaryEmail(u)
}

// QueryFvSession queries the "fv_session" edge of the User entity.
func (u *User) QueryFvSession() *FvSessionQuery {
	return NewUserClient(u.config).QueryFvSession(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
