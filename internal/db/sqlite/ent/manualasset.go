// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/manualasset"
)

// ManualAsset is the model entity for the ManualAsset schema.
type ManualAsset struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ProviderName holds the value of the "provider_name" field.
	ProviderName string `json:"provider_name,omitempty"`
	// AssetTableID holds the value of the "asset_table_id" field.
	AssetTableID string `json:"asset_table_id,omitempty"`
	// AssetType holds the value of the "asset_type" field.
	AssetType string `json:"asset_type,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ManualAsset) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case manualasset.FieldValue:
			values[i] = new(sql.NullFloat64)
		case manualasset.FieldProviderName, manualasset.FieldAssetTableID, manualasset.FieldAssetType:
			values[i] = new(sql.NullString)
		case manualasset.FieldCreatedAt, manualasset.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case manualasset.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ManualAsset fields.
func (ma *ManualAsset) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case manualasset.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ma.ID = *value
			}
		case manualasset.FieldProviderName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_name", values[i])
			} else if value.Valid {
				ma.ProviderName = value.String
			}
		case manualasset.FieldAssetTableID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field asset_table_id", values[i])
			} else if value.Valid {
				ma.AssetTableID = value.String
			}
		case manualasset.FieldAssetType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field asset_type", values[i])
			} else if value.Valid {
				ma.AssetType = value.String
			}
		case manualasset.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ma.Value = value.Float64
			}
		case manualasset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ma.CreatedAt = value.Time
			}
		case manualasset.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ma.UpdatedAt = value.Time
			}
		default:
			ma.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the ManualAsset.
// This includes values selected through modifiers, order, etc.
func (ma *ManualAsset) GetValue(name string) (ent.Value, error) {
	return ma.selectValues.Get(name)
}

// Update returns a builder for updating this ManualAsset.
// Note that you need to call ManualAsset.Unwrap() before calling this method if this ManualAsset
// was returned from a transaction, and the transaction was committed or rolled back.
func (ma *ManualAsset) Update() *ManualAssetUpdateOne {
	return NewManualAssetClient(ma.config).UpdateOne(ma)
}

// Unwrap unwraps the ManualAsset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ma *ManualAsset) Unwrap() *ManualAsset {
	_tx, ok := ma.config.driver.(*txDriver)
	if !ok {
		panic("ent: ManualAsset is not a transactional entity")
	}
	ma.config.driver = _tx.drv
	return ma
}

// String implements the fmt.Stringer.
func (ma *ManualAsset) String() string {
	var builder strings.Builder
	builder.WriteString("ManualAsset(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ma.ID))
	builder.WriteString("provider_name=")
	builder.WriteString(ma.ProviderName)
	builder.WriteString(", ")
	builder.WriteString("asset_table_id=")
	builder.WriteString(ma.AssetTableID)
	builder.WriteString(", ")
	builder.WriteString("asset_type=")
	builder.WriteString(ma.AssetType)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", ma.Value))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ma.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ma.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ManualAssets is a parsable slice of ManualAsset.
type ManualAssets []*ManualAsset
