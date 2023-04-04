package test

import (
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type webauthnSessionRepo struct {
	init []*ent.WebauthnSessionData
}

func NewWebauthnSessionRepo(init []*ent.WebauthnSessionData) dal.IWebauthnSessionRepo {
	return &webauthnSessionRepo{append([]*ent.WebauthnSessionData{}, init...)}
}

func (r *webauthnSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	var session *ent.WebauthnSessionData
	for _, s := range r.init {
		if s.Challenge == challenge {
			return s, nil
		}
	}
	return session, nil
}

func (r *webauthnSessionRepo) Create(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	r.init = append(r.init, &sessionData)
	return nil
}

func (r *webauthnSessionRepo) Delete(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	index := -1
	for i, data := range r.init {
		if data.ID == sessionData.ID {
			index = i
		}
	}
	if index > -1 {
		r.init = append(r.init[:index], r.init[index+1:]...)
	}

	return nil
}
