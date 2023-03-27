package dal

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
)

type IUserRepo interface {
	GetAll(ctx Ctx) ([]*ent.User, error)
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

func (r *userRepo) GetAll(ctx Ctx) ([]*ent.User, error) {
	us, err := r.db.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all users: %w", err)
	}

	return us, nil
}

func (r *userRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	u, err := r.db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return u, nil
}

func (r *userRepo) Create(ctx Ctx, u *ent.User) (*ent.User, error) {
	newu, err := r.db.User.
		Create().
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user by id: %w", err)
	}

	return newu, nil
}

func (r *userRepo) Count(ctx Ctx, id uuid.UUID) (int, error) {
	count, err := r.db.User.
		Query().
		Where(user.ID(id)).
		Count(ctx)

	if err != nil {
		return 0, fmt.Errorf("failed counting users: %w", err)
	}

	return count, nil
}
