package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type FinItemTable struct {
	ent.Schema
}

func (FinItemTable) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.Int32("sheet"),
		field.Int32("section"),
		field.String("category"),
		field.String("description").Default("").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (FinItemTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("fin_item_tables").Unique().Field("user_id"),
	}
}
