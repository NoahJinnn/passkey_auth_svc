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
	db *ent.Client
}

func NewUserRepo(db *ent.Client) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) All(ctx Ctx) ([]*ent.User, error) {
	us, err := r.db.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (r *userRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	u, err := r.db.User.
		Query().
		Where(user.ID(id)).
		WithEmails().
		WithPrimaryEmail().
		WithWebauthnCredentials().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepo) Create(ctx Ctx, u *ent.User) (*ent.User, error) {
	newu, err := r.db.User.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return newu, nil
}

func (r *userRepo) Count(ctx Ctx, id uuid.UUID) (int, error) {
	count, err := r.db.User.
		Query().
		Where(user.ID(id)).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
