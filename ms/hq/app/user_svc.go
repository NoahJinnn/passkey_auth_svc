package app

import (
	"fmt"

	"github.com/hellohq/hqservice/api/openapi/model"
)

func (app *App) GetAllUsers(ctx Ctx) ([]*User, error) {
	return app.repo.GetAllUsers(ctx)
}

func (app *App) GetUserById(ctx Ctx, id uint) (*User, error) {
	return app.repo.GetUserById(ctx, id)
}

func (app *App) CreateUser(ctx Ctx, u *model.User) (*User, error) {
	domU := &User{}
	domU = domU.FromOAIReq(u)
	eu, err := app.repo.CreateUser(ctx, domU)
	if err != nil {
		return nil, fmt.Errorf("create new user failed: %w", err)
	}
	domU = domU.FromEnt(eu)
	return domU, nil
}

func (app *App) UpdateUser(ctx Ctx, u *model.User) (*User, error) {
	domU := &User{}
	domU = domU.FromOAIReq(u)
	eu, err := app.repo.UpdateUser(ctx, domU)
	if err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}
	domU = domU.FromEnt(eu)
	return domU, nil
}
