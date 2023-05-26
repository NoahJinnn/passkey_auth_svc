package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type Identity struct {
	ent.Schema
}

func (Identity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("provider_id"),
		field.String("provider_name"),
		field.String("data").Optional(),
		field.UUID("email_id", uuid.UUID{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Identity) Edges() []ent.Edge {
	return []ent.Edge{edge.From("email", Email.Type).Ref("identities").Unique().Field("email_id")}
}

func (Identity) Annotations() []schema.Annotation {
	return nil
}
