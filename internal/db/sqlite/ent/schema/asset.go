package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.UUID("institution_id", uuid.UUID{}).Optional(),
		field.String("data"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).Ref("assets").Unique().Field("institution_id"),
	}
}

func (Asset) Annotations() []schema.Annotation {
	return nil
}
