package dal

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
)

type IUserRepo interface {
	GetAllUsers(ctx Ctx) ([]*ent.User, error)
	GetUserById(ctx Ctx, id uuid.UUID) (*ent.User, error)
	CreateUser(ctx Ctx, u *ent.User) (*ent.User, error)
}

type userRepo struct {
	db *ent.Client
}

func NewUserRepo(db *ent.Client) IUserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetAllUsers(ctx Ctx) ([]*ent.User, error) {
	us, err := r.db.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all users: %w", err)
	}

	return us, nil
}

func (r *userRepo) GetUserById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	u, err := r.db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return u, nil
}

func (r *userRepo) CreateUser(ctx Ctx, u *ent.User) (*ent.User, error) {
	newu, err := r.db.User.
		Create().
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user by id: %w", err)
	}

	return newu, nil
}
