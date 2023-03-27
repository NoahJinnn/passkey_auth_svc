package dal

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
)

type IWebauthnSessionRepo interface {
	Get(ctx Ctx, id uuid.UUID) (*ent.WebauthnSessionData, error)
	GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error)
	Create(ctx Ctx, sessionData ent.WebauthnSessionData) error
	Update(ctx Ctx, sessionData ent.WebauthnSessionData) error
	Delete(ctx Ctx, sessionData ent.WebauthnSessionData) error
}

type webauthnSessionRepo struct {
	db *ent.Client
}

func NewWebauthnSessionRepo(db *ent.Client) IWebauthnSessionRepo {
	return &webauthnSessionRepo{db: db}
}

func (r *webauthnSessionRepo) Get(ctx Ctx, id uuid.UUID) (*ent.WebauthnSessionData, error) {
	panic("implement me")
}

func (r *webauthnSessionRepo) GetByChallenge(ctx Ctx, challenge string) (*ent.WebauthnSessionData, error) {
	panic("implement me")
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

func (r *webauthnSessionRepo) Update(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	panic("implement me")
}

func (r *webauthnSessionRepo) Delete(ctx Ctx, sessionData ent.WebauthnSessionData) error {
	panic("implement me")
}
