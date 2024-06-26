package session

import (
	"context"
	"fmt"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ent/jwk"
	"github.com/NoahJinnn/passkey_auth_svc/internal/db"
)

type IJwkRepo interface {
	Jwk(ctx context.Context, id uint) (*ent.Jwk, error)
	All(ctx context.Context) ([]*ent.Jwk, error)
	Last(ctx context.Context) (*ent.Jwk, error)
	Create(ctx context.Context, jwk ent.Jwk) error
}

type jwkRepo struct {
	db *db.Db
}

func NewJwkRepo(db *db.Db) IJwkRepo {
	return &jwkRepo{db: db}
}

func (r *jwkRepo) Jwk(ctx context.Context, id uint) (*ent.Jwk, error) {
	jwk, err := r.db.PgEnt.Jwk.
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

func (r *jwkRepo) All(ctx context.Context) ([]*ent.Jwk, error) {
	jwks, err := r.db.PgEnt.Jwk.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying all jwks: %w", err)
	}

	return jwks, nil
}

func (r *jwkRepo) Last(ctx context.Context) (*ent.Jwk, error) {
	jwk, err := r.db.PgEnt.Jwk.
		Query().
		Order(ent.Desc(jwk.FieldCreatedAt, jwk.FieldID)).
		Limit(1).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying jwk by id: %w", err)
	}

	return jwk, nil
}

func (r *jwkRepo) Create(ctx context.Context, jwk ent.Jwk) error {
	_, err := r.db.PgEnt.Jwk.
		Create().
		SetID(jwk.ID).
		SetKeyData(jwk.KeyData).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating jwk by id: %w", err)
	}

	return nil
}
