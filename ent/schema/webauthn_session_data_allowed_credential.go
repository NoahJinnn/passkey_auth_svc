package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

type WebauthnSessionDataAllowedCredential struct {
	ent.Schema
}

func (WebauthnSessionDataAllowedCredential) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.String("credential_id"), field.UUID("webauthn_session_data_id", uuid.UUID{}).Optional(), field.Time("created_at").Default(time.Now).Immutable(), field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)}
}
func (WebauthnSessionDataAllowedCredential) Edges() []ent.Edge {
	return []ent.Edge{edge.From("webauthn_session_data", WebauthnSessionData.Type).Ref("webauthn_session_data_allowed_credentials").Unique().Field("webauthn_session_data_id")}
}
func (WebauthnSessionDataAllowedCredential) Annotations() []schema.Annotation {
	return nil
}
