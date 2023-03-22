package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type Passcode struct {
	ent.Schema
}

func (Passcode) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.UUID("user_id", uuid.UUID{}).Optional(), field.Int32("ttl"), field.String("code"), field.Int32("try_count"), field.Time("created_at").Default(time.Now).Immutable(), field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), field.UUID("email_id", uuid.UUID{}).Optional()}
}
func (Passcode) Edges() []ent.Edge {
	return []ent.Edge{edge.From("email", Email.Type).Ref("passcodes").Unique().Field("email_id"), edge.From("user", User.Type).Ref("passcodes").Unique().Field("user_id")}
}
func (Passcode) Annotations() []schema.Annotation {
	return nil
}
