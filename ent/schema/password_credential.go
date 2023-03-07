// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type PasswordCredential struct {
	ent.Schema
}

func (PasswordCredential) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.UUID("user_id").Optional().Unique(uuid.UUID{}), field.String("password"), field.Time("created_at"), field.Time("updated_at")}
}
func (PasswordCredential) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("password_credential").Unique().Field("user_id")}
}
func (PasswordCredential) Annotations() []schema.Annotation {
	return nil
}
