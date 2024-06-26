package test

import (
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
)

type waCredentialRepo struct {
	init []*ent.WebauthnCredential
}

func NewWebauthnCredentialRepo(init []*ent.WebauthnCredential) dal.IWebauthnCredentialRepo {
	return &waCredentialRepo{append([]*ent.WebauthnCredential{}, init...)}
}

func (r *waCredentialRepo) GetById(ctx Ctx, id string) (*ent.WebauthnCredential, error) {
	panic("implement me")
}

func (r *waCredentialRepo) ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error) {
	var found []*ent.WebauthnCredential
	for _, cre := range r.init {
		if cre.UserID == userId {
			found = append(found, cre)
		}
	}

	return found, nil
}

func (r *waCredentialRepo) Create(ctx Ctx, credential ent.WebauthnCredential, transports []protocol.AuthenticatorTransport) error {
	r.init = append(r.init, &credential)
	return nil
}

func (r *waCredentialRepo) Update(ctx Ctx, credential ent.WebauthnCredential) error {
	for i, data := range r.init {
		if data.ID == credential.ID {
			r.init[i] = &credential
		}
	}
	return nil
}

func (r *waCredentialRepo) Delete(ctx Ctx, credential ent.WebauthnCredential) error {
	panic("implement me")
}
