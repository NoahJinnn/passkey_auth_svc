// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Collectible struct {
	ent.Schema
}

func (Collectible) Fields() []ent.Field {
	return []ent.Field{field.Uint("id").SchemaType(map[string]string{"postgres": "serial"}), field.Int32("user_id").Optional(), field.String("name"), field.String("description"), field.Time("created_at"), field.Time("updated_at")}
}
func (Collectible) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("collectibles").Unique().Field("user_id")}
}
func (Collectible) Annotations() []schema.Annotation {
	return nil
}
