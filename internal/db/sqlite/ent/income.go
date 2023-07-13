// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/income"
)

// Income is the model entity for the Income schema.
type Income struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ProviderName holds the value of the "provider_name" field.
	ProviderName string `json:"provider_name,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	institution_incomes *uuid.UUID
	selectValues        sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Income) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case income.FieldProviderName, income.FieldData:
			values[i] = new(sql.NullString)
		case income.FieldCreatedAt, income.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case income.FieldID:
			values[i] = new(uuid.UUID)
		case income.ForeignKeys[0]: // institution_incomes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Income fields.
func (i *Income) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case income.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case income.FieldProviderName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_name", values[j])
			} else if value.Valid {
				i.ProviderName = value.String
			}
		case income.FieldData:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[j])
			} else if value.Valid {
				i.Data = value.String
			}
		case income.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case income.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case income.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field institution_incomes", values[j])
			} else if value.Valid {
				i.institution_incomes = new(uuid.UUID)
				*i.institution_incomes = *value.S.(*uuid.UUID)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Income.
// This includes values selected through modifiers, order, etc.
func (i *Income) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Income.
// Note that you need to call Income.Unwrap() before calling this method if this Income
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Income) Update() *IncomeUpdateOne {
	return NewIncomeClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Income entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Income) Unwrap() *Income {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Income is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Income) String() string {
	var builder strings.Builder
	builder.WriteString("Income(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("provider_name=")
	builder.WriteString(i.ProviderName)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(i.Data)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Incomes is a parsable slice of Income.
type Incomes []*Income