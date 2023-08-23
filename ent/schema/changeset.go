package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

// Changeset holds the schema definition for the Changeset entity.
type Changeset struct {
	ent.Schema
}

// Fields of the Changeset.
func (Changeset) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("site_id"),
		field.Int32("db_version"),
		field.Bool("first_launch").Default(false),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Changeset.
func (Changeset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("changesets").Unique().Field("user_id"),
	}
}
