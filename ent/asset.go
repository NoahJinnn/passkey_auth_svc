// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/asset"
	"github.com/hellohq/hqservice/ent/user"
)

// Asset is the model entity for the Asset schema.
type Asset struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Sheet holds the value of the "sheet" field.
	Sheet int32 `json:"sheet,omitempty"`
	// Section holds the value of the "section" field.
	Section int32 `json:"section,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// ProviderName holds the value of the "provider_name" field.
	ProviderName string `json:"provider_name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AssetQuery when eager-loading is set.
	Edges        AssetEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AssetEdges holds the relations/edges for other nodes in the graph.
type AssetEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AssetEdges) UserOrErr() (*User, error) {
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
func (*Asset) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case asset.FieldValue:
			values[i] = new(sql.NullFloat64)
		case asset.FieldSheet, asset.FieldSection:
			values[i] = new(sql.NullInt64)
		case asset.FieldType, asset.FieldProviderName, asset.FieldDescription, asset.FieldCurrency:
			values[i] = new(sql.NullString)
		case asset.FieldCreatedAt, asset.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case asset.FieldID, asset.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Asset fields.
func (a *Asset) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asset.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case asset.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				a.UserID = *value
			}
		case asset.FieldSheet:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sheet", values[i])
			} else if value.Valid {
				a.Sheet = int32(value.Int64)
			}
		case asset.FieldSection:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field section", values[i])
			} else if value.Valid {
				a.Section = int32(value.Int64)
			}
		case asset.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case asset.FieldProviderName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_name", values[i])
			} else if value.Valid {
				a.ProviderName = value.String
			}
		case asset.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case asset.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				a.Currency = value.String
			}
		case asset.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				a.Value = value.Float64
			}
		case asset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case asset.FieldUpdatedAt:
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

// GetValue returns the ent.Value that was dynamically selected and assigned to the Asset.
// This includes values selected through modifiers, order, etc.
func (a *Asset) GetValue(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Asset entity.
func (a *Asset) QueryUser() *UserQuery {
	return NewAssetClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Asset.
// Note that you need to call Asset.Unwrap() before calling this method if this Asset
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Asset) Update() *AssetUpdateOne {
	return NewAssetClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Asset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Asset) Unwrap() *Asset {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Asset is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Asset) String() string {
	var builder strings.Builder
	builder.WriteString("Asset(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	builder.WriteString("sheet=")
	builder.WriteString(fmt.Sprintf("%v", a.Sheet))
	builder.WriteString(", ")
	builder.WriteString("section=")
	builder.WriteString(fmt.Sprintf("%v", a.Section))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(a.Type)
	builder.WriteString(", ")
	builder.WriteString("provider_name=")
	builder.WriteString(a.ProviderName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(a.Currency)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", a.Value))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Assets is a parsable slice of Asset.
type Assets []*Asset
