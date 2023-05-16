//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetFvSvc() finverse.IFvAuthSvc
	GetSeAccountInfoSvc() saltedge.ISeAccountInfoSvc
}

// App implements interface Appl.
type App struct {
	cfg              *config.Config
	repo             *dal.NwRepo
	seAccountInfoSvc saltedge.ISeAccountInfoSvc
}

// New creates and returns new App.
func New(cfg *config.Config, repo *dal.NwRepo) App {
	seAccountInfoSvc := saltedge.NewSeAccountInfoSvc(cfg)

	return App{
		cfg:              cfg,
		repo:             repo,
		seAccountInfoSvc: seAccountInfoSvc,
	}
}

func (a App) GetSeAccountInfoSvc() saltedge.ISeAccountInfoSvc {
	return a.seAccountInfoSvc
}
