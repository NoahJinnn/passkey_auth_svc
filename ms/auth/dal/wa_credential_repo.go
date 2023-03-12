package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
)

type IWebauthnCredentialRepo interface {
	GetById(ctx Ctx, id string) (*ent.WebauthnCredential, error)
	Create(ctx Ctx, credential ent.WebauthnCredential) error
	Update(ctx Ctx, credential ent.WebauthnCredential) error
	Delete(ctx Ctx, credential ent.WebauthnCredential) error
	GetFromUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error)
}

type webauthnRepo struct {
	db *ent.Client
}

func NewWebauthnCredentialRepo(db *ent.Client) IWebauthnCredentialRepo {
	return &webauthnRepo{db: db}
}

func (r *webauthnRepo) GetById(ctx Ctx, id string) (*ent.WebauthnCredential, error) {
	panic("implement me")
}

func (r *webauthnRepo) GetFromUser(ctx Ctx, userId uuid.UUID) ([]*ent.WebauthnCredential, error) {

	// Query all ent.WebauthnCredential by ent.User id and sort by created at return them
	credentials, err := r.db.WebauthnCredential.
		Query().
		Where(webauthncredential.HasUserWith(user.ID(userId))).
		Order(ent.Asc(webauthncredential.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (r *webauthnRepo) Create(ctx Ctx, credential ent.WebauthnCredential) error {
	panic("implement me")
}

func (r *webauthnRepo) Update(ctx Ctx, credential ent.WebauthnCredential) error {
	panic("implement me")
}

func (r *webauthnRepo) Delete(ctx Ctx, credential ent.WebauthnCredential) error {
	panic("implement me")
}
