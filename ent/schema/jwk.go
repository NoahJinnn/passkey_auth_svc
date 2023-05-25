package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Jwk struct {
	ent.Schema
}

func (Jwk) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}),
		field.String("key_data"),
		field.Time("created_at").Default(time.Now).Immutable().Default(time.Now).Immutable(),
	}
}

func (Jwk) Edges() []ent.Edge {
	return nil
}

func (Jwk) Annotations() []schema.Annotation {
	return nil
}
