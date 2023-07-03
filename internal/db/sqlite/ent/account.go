// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/institution"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// InstitutionID holds the value of the "institution_id" field.
	InstitutionID uuid.UUID `json:"institution_id,omitempty"`
	// Data holds the value of the "data" field.
	Data *string `json:"data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Institution holds the value of the institution edge.
	Institution *Institution `json:"institution,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstitutionOrErr returns the Institution value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) InstitutionOrErr() (*Institution, error) {
	if e.loadedTypes[0] {
		if e.Institution == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: institution.Label}
		}
		return e.Institution, nil
	}
	return nil, &NotLoadedError{edge: "institution"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldData:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case account.FieldID, account.FieldInstitutionID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldInstitutionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field institution_id", values[i])
			} else if value != nil {
				a.InstitutionID = *value
			}
		case account.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				a.Data = new(string)
				*a.Data = value.String
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryInstitution queries the "institution" edge of the Account entity.
func (a *Account) QueryInstitution() *InstitutionQuery {
	return NewAccountClient(a.config).QueryInstitution(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("institution_id=")
	builder.WriteString(fmt.Sprintf("%v", a.InstitutionID))
	builder.WriteString(", ")
	if v := a.Data; v != nil {
		builder.WriteString("data=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
