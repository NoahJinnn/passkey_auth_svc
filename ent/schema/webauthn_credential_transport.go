package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type WebauthnCredentialTransport struct {
	ent.Schema
}

func (WebauthnCredentialTransport) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("webauthn_credential_id").Optional(),
	}
}
func (WebauthnCredentialTransport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("webauthn_credential", WebauthnCredential.Type).Ref("webauthn_credential_transports").Unique().Field("webauthn_credential_id"),
	}
}
func (WebauthnCredentialTransport) Annotations() []schema.Annotation {
	return nil
}
