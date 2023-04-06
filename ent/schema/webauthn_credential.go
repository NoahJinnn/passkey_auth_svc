package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type WebauthnCredential struct {
	ent.Schema
}

func (WebauthnCredential) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.String("public_key"),
		field.String("attestation_type"),
		field.UUID("aaguid", uuid.UUID{}),
		field.Int32("sign_count"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("name").Optional(),
		field.Bool("backup_eligible"),
		field.Bool("backup_state"),
		field.Time("last_used_at").Optional(),
	}
}
func (WebauthnCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("webauthn_credential_transports", WebauthnCredentialTransport.Type),
		edge.From("user", User.Type).Ref("webauthn_credentials").Unique().Field("user_id"),
	}
}
func (WebauthnCredential) Annotations() []schema.Annotation {
	return nil
}
