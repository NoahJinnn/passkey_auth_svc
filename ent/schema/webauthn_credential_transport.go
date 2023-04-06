package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type WebauthnCredentialTransport struct {
	ent.Schema
}

func (WebauthnCredentialTransport) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			id, _ := uuid.NewV4()
			return id
		}).Immutable(),
		field.String("name"),
		field.String("webauthn_credential_id").Optional(),
	}
}
func (WebauthnCredentialTransport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("webauthn_credential", WebauthnCredential.Type).
			Ref("webauthn_credential_transports").
			Unique().
			Field("webauthn_credential_id"),
	}
}
func (WebauthnCredentialTransport) Annotations() []schema.Annotation {
	return nil
}
