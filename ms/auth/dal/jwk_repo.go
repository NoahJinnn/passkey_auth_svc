package dal

import (
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/jwk"
)

func (r *Repo) GetJwk(ctx Ctx, id uint) (*ent.Jwk, error) {
	jwk, err := r.Db.Jwk.
		Query().
		Where(jwk.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return jwk, nil

}

func (r *Repo) GetAllJwk(ctx Ctx) ([]*ent.Jwk, error) {
	jwks, err := r.Db.Jwk.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all users: %w", err)
	}

	return jwks, nil
}

func (r *Repo) GetLastJwk(ctx Ctx) (*ent.Jwk, error) {
	jwk, err := r.Db.Jwk.
		Query().
		Order(ent.Desc(jwk.FieldCreatedAt, jwk.FieldID)).
		Limit(1).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return jwk, nil
}

func (r *Repo) Create(ctx Ctx, jwk ent.Jwk) error {
	_, err := r.Db.Jwk.
		Create().
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed creating jwk by id: %w", err)
	}

	return nil
}
