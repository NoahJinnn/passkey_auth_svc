package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type Asset struct {
	ent.Schema
}

func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.Int32("sheet").Optional().Default(0),
		field.Int32("section").Optional().Default(0),
		field.String("type").Default(""),
		field.String("provider_name").Default(""),
		field.String("description").Default("").Optional(),
		field.String("currency").Default(""),
		field.Float("value").SchemaType(map[string]string{dialect.Postgres: "numeric(19,6)"}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("assets").Unique().Field("user_id"),
	}
}
