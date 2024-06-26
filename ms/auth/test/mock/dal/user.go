package test

import (
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/gofrs/uuid"
)

func NewUserRepo(init []*ent.User) dal.IUserRepo {
	return &userRepo{append([]*ent.User{}, init...)}
}

type userRepo struct {
	users []*ent.User
}

func (r *userRepo) All(ctx Ctx) ([]*ent.User, error) {
	return r.users, nil
}

func (r *userRepo) GetById(ctx Ctx, id uuid.UUID) (*ent.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (r *userRepo) Create(ctx Ctx, u *ent.User) (*ent.User, error) {
	r.users = append(r.users, u)

	return u, nil
}

func (r *userRepo) Count(ctx Ctx, id uuid.UUID) (int, error) {
	for _, u := range r.users {
		if u.ID == id {
			return 1, nil
		}
	}
	return 0, nil
}
