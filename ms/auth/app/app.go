//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"context"
	"errors"

	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrAccessDenied  = errors.New("access denied")
	ErrAlreadyExist  = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrValidate      = errors.New("validate")
	ErrWrongPassword = errors.New("wrong password")
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetWebauthnSvc() svcs.IWebauthnSvc
	GetUserSvc() svcs.IUserSvc
}

// App implements interface Appl.
type App struct {
	cfg  *config.Config
	repo *dal.Repo
}

// New creates and returns new App.
func New(cfg *config.Config, repo *dal.Repo) App {

	return App{
		cfg:  cfg,
		repo: repo,
	}
}

func (a App) GetWebauthnSvc() svcs.IWebauthnSvc {
	return svcs.NewWebAuthn(a.cfg, a.repo)
}

func (a App) GetUserSvc() svcs.IUserSvc {
	return svcs.NewUserSvc(a.cfg, a.repo)
}
