// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/passwordcredential"
	"github.com/hellohq/hqservice/ent/user"
)

// PasswordCredential is the model entity for the PasswordCredential schema.
type PasswordCredential struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PasswordCredentialQuery when eager-loading is set.
	Edges PasswordCredentialEdges `json:"edges"`
}

// PasswordCredentialEdges holds the relations/edges for other nodes in the graph.
type PasswordCredentialEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PasswordCredentialEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PasswordCredential) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case passwordcredential.FieldPassword:
			values[i] = new(sql.NullString)
		case passwordcredential.FieldCreatedAt, passwordcredential.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case passwordcredential.FieldID, passwordcredential.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PasswordCredential", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PasswordCredential fields.
func (pc *PasswordCredential) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case passwordcredential.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pc.ID = *value
			}
		case passwordcredential.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				pc.UserID = *value
			}
		case passwordcredential.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				pc.Password = value.String
			}
		case passwordcredential.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case passwordcredential.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the PasswordCredential entity.
func (pc *PasswordCredential) QueryUser() *UserQuery {
	return NewPasswordCredentialClient(pc.config).QueryUser(pc)
}

// Update returns a builder for updating this PasswordCredential.
// Note that you need to call PasswordCredential.Unwrap() before calling this method if this PasswordCredential
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PasswordCredential) Update() *PasswordCredentialUpdateOne {
	return NewPasswordCredentialClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PasswordCredential entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PasswordCredential) Unwrap() *PasswordCredential {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PasswordCredential is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PasswordCredential) String() string {
	var builder strings.Builder
	builder.WriteString("PasswordCredential(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.UserID))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(pc.Password)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PasswordCredentials is a parsable slice of PasswordCredential.
type PasswordCredentials []*PasswordCredential
