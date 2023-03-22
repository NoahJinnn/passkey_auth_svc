package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type Email struct {
	ent.Schema
}

func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.String("address").Unique(),
		field.Bool("verified"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
func (Email) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("emails").Unique().Field("user_id"), edge.To("identities", Identity.Type), edge.To("passcodes", Passcode.Type), edge.To("primary_email", PrimaryEmail.Type).Unique()}
}
func (Email) Annotations() []schema.Annotation {
	return nil
}
