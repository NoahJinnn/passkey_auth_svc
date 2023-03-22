package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type WebauthnSessionData struct {
	ent.Schema
}

func (WebauthnSessionData) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("challenge").Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.String("user_verification"),
		field.String("operation"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Default(time.Now).UpdateDefault(time.Now),
	}
}
func (WebauthnSessionData) Edges() []ent.Edge {
	return []ent.Edge{edge.To("webauthn_session_data_allowed_credentials", WebauthnSessionDataAllowedCredential.Type)}
}
func (WebauthnSessionData) Annotations() []schema.Annotation {
	return nil
}
