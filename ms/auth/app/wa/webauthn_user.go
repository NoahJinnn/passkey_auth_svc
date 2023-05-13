package wa

import (
	"context"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
)

func NewWebauthnUser(ctx context.Context, user ent.User, credentials []*ent.WebauthnCredential) (*WebauthnUser, error) {
	emails := user.Edges.Emails
	primEmail := user.Edges.PrimaryEmail
	var email ent.Email
	for _, m := range emails {
		if m.UserID == primEmail.UserID {
			email = *m
			break
		}
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
