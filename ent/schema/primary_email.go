package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type PrimaryEmail struct {
	ent.Schema
}

func (PrimaryEmail) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("email_id", uuid.UUID{}).Optional().Unique(),
		field.UUID("user_id", uuid.UUID{}).Optional().Nillable().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (PrimaryEmail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("email", Email.Type).Ref("primary_email").Unique().Field("email_id"),
		edge.From("user", User.Type).Ref("primary_email").Unique().Field("user_id"),
	}
}

func (PrimaryEmail) Annotations() []schema.Annotation {
	return nil
}
