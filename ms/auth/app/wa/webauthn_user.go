package wa

import (
	"context"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
)

func NewWebauthnUser(ctx context.Context, user ent.User, credentials []*ent.WebauthnCredential) (*WebauthnUser, error) {
	email, err := user.Edges.PrimaryEmail.QueryEmail().First(ctx)
	if err != nil {
		return nil, err
	}

	return &WebauthnUser{
		UserId:              user.ID,
		Email:               email.Address,
		WebauthnCredentials: credentials,
	}, nil
}

type WebauthnUser struct {
	UserId              uuid.UUID
	Email               string
	WebauthnCredentials []*ent.WebauthnCredential
}

func (u *WebauthnUser) WebAuthnID() []byte {
	return u.UserId.Bytes()
}

func (u *WebauthnUser) WebAuthnName() string {
	return u.Email
}

func (u *WebauthnUser) WebAuthnDisplayName() string {
	return u.Email
}

func (u *WebauthnUser) WebAuthnIcon() string {
	return ""
}

func (u *WebauthnUser) WebAuthnCredentials() []webauthn.Credential {
	var credentials []webauthn.Credential
	for _, credential := range u.WebauthnCredentials {
		cred := credential
		c := WebauthnCredentialFromModel(cred)
		credentials = append(credentials, *c)
	}

	return credentials
}
