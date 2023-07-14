package schema

import (
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type ItemTable struct {
	ent.Schema
}

func (ItemTable) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.Int32("sheet"),
		field.Int32("section"),
		field.String("category").Validate(func(s string) error {
			if s != "asset" && s != "debt" {
				return errors.New("category must be asset|debt")
			}
			return nil
		}),
		field.String("description").Default("").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (ItemTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("item_tables").Unique().Field("user_id"),
	}
}
