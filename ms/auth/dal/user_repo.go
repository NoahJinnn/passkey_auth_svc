package dal

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
)

func (r *Repo) GetAllUsers(ctx Ctx) ([]*ent.User, error) {
	us, err := r.Db.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all users: %w", err)
	}

	return us, nil
}

func (r *Repo) GetUserById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	u, err := r.Db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	return u, nil
}

func (r *Repo) CreateUser(ctx Ctx, u *ent.User) (*ent.User, error) {
	u, err := r.Db.User.
		Create().
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user by id: %w", err)
	}

	return u, nil
}
