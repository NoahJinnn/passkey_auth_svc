package dal

import (
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/jwk"
)

type IJwkRepo interface {
	GetJwk(ctx Ctx, id uint) (*ent.Jwk, error)
	GetAllJwk(ctx Ctx) ([]*ent.Jwk, error)
	GetLastJwk(ctx Ctx) (*ent.Jwk, error)
	Create(ctx Ctx, jwk ent.Jwk) error
}

type jwkRepo struct {
	db *ent.Client
}

func NewJwkRepo(db *ent.Client) IJwkRepo {
	return &jwkRepo{db: db}
}

func (r *jwkRepo) GetJwk(ctx Ctx, id uint) (*ent.Jwk, error) {
	jwk, err := r.db.Jwk.
		Query().
		Where(jwk.ID(id)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed querying jwk by id: %w", err)
	}

	return jwk, nil

}

func (r *jwkRepo) GetAllJwk(ctx Ctx) ([]*ent.Jwk, error) {
	jwks, err := r.db.Jwk.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all jwks: %w", err)
	}

	return jwks, nil
}

func (r *jwkRepo) GetLastJwk(ctx Ctx) (*ent.Jwk, error) {
	jwk, err := r.db.Jwk.
		Query().
		Order(ent.Desc(jwk.FieldCreatedAt, jwk.FieldID)).
		Limit(1).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying jwk by id: %w", err)
	}

	return jwk, nil
}

func (r *jwkRepo) Create(ctx Ctx, jwk ent.Jwk) error {
	_, err := r.db.Jwk.
		Create().
		SetID(jwk.ID).
		SetKeyData(jwk.KeyData).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed creating jwk by id: %w", err)
	}

	return nil
}
