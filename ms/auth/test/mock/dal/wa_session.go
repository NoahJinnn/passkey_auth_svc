package test

import (
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
)

type waSessionRepo struct {
	init []*ent.WebauthnSessionData
}

func NewWebauthnSessionRepo(init []*ent.WebauthnSessionData) dal.IWebauthnSessionRepo {
	return &waSessionRepo{append([]*ent.WebauthnSessionData{}, init...)}
}

func (r *waSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	var session *ent.WebauthnSessionData
	for _, s := range r.init {
		if s.Challenge == challenge {
			return s, nil
		}
	}
	return session, nil
}

func (r *waSessionRepo) Create(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	r.init = append(r.init, &sessionData)
	return nil
}

func (r *waSessionRepo) Delete(ctx Ctx, sessionData ent.WebauthnSessionData) error {
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
