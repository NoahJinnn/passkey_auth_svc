package dal

import (
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/webauthnsessiondata"
)

type IWebauthnSessionRepo interface {
	GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error)
	Create(ctx Ctx, sessionData ent.WebauthnSessionData) error
	Delete(ctx Ctx, session ent.WebauthnSessionData) error
}

type webauthnSessionRepo struct {
	db *ent.Client
}

func NewWebauthnSessionRepo(db *ent.Client) IWebauthnSessionRepo {
	return &webauthnSessionRepo{db: db}
}

func (r *webauthnSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	var sessionData []*ent.WebauthnSessionData
	sessionData, err := r.db.WebauthnSessionData.Query().Where(webauthnsessiondata.Challenge(challenge)).All(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get sessionData: %w", err)
	}

	if len(sessionData) <= 0 {
		return nil, nil
	}

	return sessionData[0], nil
}

func (r *webauthnSessionRepo) Create(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	_, err := r.db.WebauthnSessionData.Create().
		SetUserID(sessionData.ID).
		SetChallenge(sessionData.Challenge).
		SetOperation(sessionData.Operation).
		SetUserVerification(sessionData.UserVerification).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed to store sessionData: %w", err)
	}

	return nil
}

func (r *webauthnSessionRepo) Delete(ctx Ctx, session ent.WebauthnSessionData) error {
	err := r.db.WebauthnSessionData.DeleteOne(&session).Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to delete sessionData: %w", err)
	}

	return nil
}
