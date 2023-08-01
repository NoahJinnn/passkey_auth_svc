package schema

import (
	"errors"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type ManualItem struct {
	ent.Schema
}

func (ManualItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("provider_name").Default("manual").Validate(func(s string) error {
			if s != "manual" && s != "finverse" {
				return errors.New("provider name must be manual|finverse")
			}
			return nil
		}),
		field.String("item_table_id").Default("asset"),
		field.String("type").Default("cash"),
		field.String("category").Default("asset").Validate(func(s string) error {
			if s != "asset" && s != "debt" {
				return errors.New("category must be asset|debt")
			}
			return nil
		}),
		field.String("description").Default("").Optional(),
		field.Float("value").Default(math.SmallestNonzeroFloat32),
	}
}
