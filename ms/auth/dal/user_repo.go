package dal

import (
	"fmt"
	"log"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/user"
	dom "github.com/hellohq/hqservice/ms/auth/app"
)

func (repo *Repo) GetAllUsers(ctx Ctx) ([]*dom.User, error) {
	eus, err := repo.Db.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying all users: %w", err)
	}

	us := dom.UserListFromEnt(eus)

	log.Println("users returned: ", us)
	return us, nil
}

func (repo *Repo) GetUserById(ctx Ctx, id uint) (*dom.User, error) {
	eu, err := repo.Db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	u := &dom.User{}
	u = u.FromEnt(eu)
	log.Println("users returned: ", u)
	return u, nil
}

func (repo *Repo) CreateUser(ctx Ctx, u *dom.User) (*ent.User, error) {
	eu, err := repo.Db.User.
		Create().
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPhoneNumber(u.PhoneNumber).
		SetAddress(u.Address).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	log.Println("users returned: ", u)
	return eu, nil
}

func (repo *Repo) UpdateUser(ctx Ctx, u *dom.User) (*ent.User, error) {
	eu, err := repo.Db.User.
		UpdateOneID(u.ID).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPhoneNumber(u.PhoneNumber).
		SetAddress(u.Address).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user by id: %w", err)
	}

	log.Println("users returned: ", u)
	return eu, nil
}
