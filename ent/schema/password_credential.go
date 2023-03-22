package schema

import (
	"time"

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
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.UUID("user_id", uuid.UUID{}).Optional().Unique(), field.String("password"), field.Time("created_at").Default(time.Now).Immutable(), field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)}
}
func (PasswordCredential) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("password_credential").Unique().Field("user_id")}
}
func (PasswordCredential) Annotations() []schema.Annotation {
	return nil
}
