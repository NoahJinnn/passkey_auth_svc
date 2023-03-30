package test

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type webauthnRepo struct {
	init []*ent.WebauthnCredential
}

func NewWebauthnCredentialRepo(init []*ent.WebauthnCredential) dal.IWebauthnCredentialRepo {
	return &webauthnRepo{append([]*ent.WebauthnCredential{}, init...)}
}

func (r *webauthnRepo) GetById(ctx Ctx, id string) (*ent.WebauthnCredential, error) {
	panic("implement me")
}

func (r *webauthnRepo) GetFromUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error) {
	var found []*ent.WebauthnCredential
	for _, cre := range r.init {
		if cre.UserID == userId {
			found = append(found, cre)
		}
	}

	return found, nil
}

func (r *webauthnRepo) Create(ctx Ctx, credential ent.WebauthnCredential) error {
	r.init = append(r.init, &credential)
	return nil
}
