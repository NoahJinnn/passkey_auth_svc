//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/nw_track/config"
	"github.com/hellohq/hqservice/ms/nw_track/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	// GetUserSvc() svcs.IUserSvc
	// GetPasscodeSvc() svcs.IPasscodeSvc
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

// func (a App) GetUserSvc() svcs.IUserSvc {
// 	return svcs.NewUserSvc(a.cfg, a.repo)
// }

// func (a App) GetPasscodeSvc() svcs.IPasscodeSvc {
// 	return svcs.NewPasscodeSvc(a.cfg, a.repo)
// }
