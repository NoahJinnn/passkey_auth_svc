package test

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type webauthnSessionRepo struct {
	init []*ent.WebauthnSessionData
}

func NewWebauthnSessionRepo(init []*ent.WebauthnSessionData) dal.IWebauthnSessionRepo {
	return &webauthnSessionRepo{append([]*ent.WebauthnSessionData{}, init...)}
}

func (r *webauthnSessionRepo) Get(ctx Ctx, id uuid.UUID) (*ent.WebauthnSessionData, error) {
	panic("implement me")
}

func (r *webauthnSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	panic("implement me")
}

func (r *webauthnSessionRepo) Create(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	return nil
}

func (r *webauthnSessionRepo) Update(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	panic("implement me")
}

func (r *webauthnSessionRepo) Delete(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	panic("implement me")
}
