package app

import (
	"fmt"

	"github.com/hellohq/hqservice/ms/auth/app/dom"
)

func (app *App) GetAllUsers(ctx Ctx) ([]*dom.User, error) {
	return app.repo.GetAllUsers(ctx)
}

func (app *App) GetUserById(ctx Ctx, id uint) (*dom.User, error) {
	return app.repo.GetUserById(ctx, id)
}

func (app *App) CreateUser(ctx Ctx, u *dom.User) (*dom.User, error) {

	u, err := app.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create new user failed: %w", err)
	}
	return u, nil
}

func (app *App) UpdateUser(ctx Ctx, u *dom.User) (*dom.User, error) {

	_, err := app.repo.UpdateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}
	return u, nil
}
