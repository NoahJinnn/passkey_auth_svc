package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type FvSession struct {
	ent.Schema
}

func (FvSession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Optional().Nillable().Unique(),
		field.String("access_token"),
		field.Int("expires_in"),
		field.String("issued_at"),
		field.String("token_type"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (FvSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("fv_session").Unique().Field("user_id"),
	}
}

func (FvSession) Annotations() []schema.Annotation {
	return nil
}
