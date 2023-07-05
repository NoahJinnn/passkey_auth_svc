package dal

import (
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/webauthnsessiondata"
)

type IWebauthnSessionRepo interface {
	GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error)
	Create(ctx Ctx, sessionData ent.WebauthnSessionData) error
	Delete(ctx Ctx, session ent.WebauthnSessionData) error
}

type waSessionRepo struct {
	pgsql *ent.Client
}

func NewWebauthnSessionRepo(pgsql *ent.Client) *waSessionRepo {
	return &waSessionRepo{pgsql: pgsql}
}

func (r *waSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	var sessionData []*ent.WebauthnSessionData
	sessionData, err := r.pgsql.WebauthnSessionData.Query().Where(webauthnsessiondata.Challenge(challenge)).All(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if len(sessionData) <= 0 {
		return nil, nil
	}

	return sessionData[0], nil
}

func (r *waSessionRepo) Create(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	_, err := r.pgsql.WebauthnSessionData.Create().
		SetUserID(sessionData.UserID).
		SetChallenge(sessionData.Challenge).
		SetOperation(sessionData.Operation).
		SetUserVerification(sessionData.UserVerification).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *waSessionRepo) Delete(ctx Ctx, session ent.WebauthnSessionData) error {
	err := r.pgsql.WebauthnSessionData.DeleteOne(&session).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
