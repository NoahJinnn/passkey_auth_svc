package dal

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
)

type IUserRepo interface {
	All(ctx Ctx) ([]*ent.User, error)
	GetById(ctx Ctx, id uuid.UUID) (*ent.User, error)
	Create(ctx Ctx, u *ent.User) (*ent.User, error)
	Count(ctx Ctx, id uuid.UUID) (int, error)
}

type userRepo struct {
	pgsql *ent.Client
}

func NewUserRepo(pgsql *ent.Client) *userRepo {
	return &userRepo{pgsql: pgsql}
}

func (r *userRepo) All(ctx Ctx) ([]*ent.User, error) {
	us, err := r.pgsql.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (r *userRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	u, err := r.pgsql.User.
		Query().
		Where(user.ID(id)).
		WithPrimaryEmail(
			func(q *ent.PrimaryEmailQuery) {
				q.Limit(1)
				q.WithEmail()
			},
		).
		WithWebauthnCredentials().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return u, nil
}

func (r *userRepo) Create(ctx Ctx, u *ent.User) (*ent.User, error) {
	newu, err := r.pgsql.User.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return newu, nil
}

func (r *userRepo) Count(ctx Ctx, id uuid.UUID) (int, error) {
	count, err := r.pgsql.User.
		Query().
		Where(user.ID(id)).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
