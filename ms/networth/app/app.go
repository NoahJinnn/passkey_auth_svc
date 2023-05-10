//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/networth/app/svcs"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	// GetUserSvc() svcs.IUserSvc
	GetSeSvc() svcs.ISeSvc
}

// App implements interface Appl.
type App struct {
	cfg  *config.Config
	repo *dal.NwRepo
}

// New creates and returns new App.
func New(cfg *config.Config, repo *dal.NwRepo) App {

	return App{
		cfg:  cfg,
		repo: repo,
	}
}

// func (a App) GetUserSvc() svcs.IUserSvc {
// 	return svcs.NewUserSvc(a.cfg, a.repo)
// }

func (a App) GetSeSvc() svcs.ISeSvc {
	return svcs.NewSeSvc(a.cfg)
}
