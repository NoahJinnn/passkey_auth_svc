package dto

import (
	"time"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/gofrs/uuid"
)

type CreateUserResponse struct {
	ID      uuid.UUID `json:"id"` // deprecated
	UserID  uuid.UUID `json:"user_id"`
	EmailID uuid.UUID `json:"email_id"`
}

type GetUserResponse struct {
	ID                  uuid.UUID                 `json:"id"`
	Email               *string                   `json:"email,omitempty"`
	WebauthnCredentials []*ent.WebauthnCredential `json:"webauthn_credentials"` // deprecated
	UpdatedAt           time.Time                 `json:"updated_at"`
	CreatedAt           time.Time                 `json:"created_at"`
}

type UserInfoResponse struct {
	ID                    uuid.UUID `json:"id"`
	EmailID               uuid.UUID `json:"email_id"`
	Verified              bool      `json:"verified"`
	HasWebauthnCredential bool      `json:"has_webauthn_credential"`
}
