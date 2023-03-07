// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Jwk struct {
	ent.Schema
}

func (Jwk) Fields() []ent.Field {
	return []ent.Field{field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}), field.String("key_data"), field.Time("created_at")}
}
func (Jwk) Edges() []ent.Edge {
	return nil
}
func (Jwk) Annotations() []schema.Annotation {
	return nil
}
