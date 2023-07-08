package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type Institution struct {
	ent.Schema
}

func (Institution) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("provider_name"),
		field.String("data"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connection", Connection.Type).Unique(),
		edge.To("accounts", Account.Type),
		edge.To("assets", Asset.Type),
	}
}

func (Institution) Annotations() []schema.Annotation {
	return nil
}
