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
	GetFvAuthSvc() *finverse.FvAuthSvc
	GetFvDataSvc() *finverse.FvDataSvc
	GetSeAccountInfoSvc() *saltedge.SeAccountInfoSvc
}

// App implements interface Appl.
type App struct {
	cfg              *config.Config
	seAccountInfoSvc *saltedge.SeAccountInfoSvc
	fvAuthSvc        *finverse.FvAuthSvc
	fvDataSvc        *finverse.FvDataSvc
}

// New creates and returns new App.
func New(cfg *config.Config, repo dal.INwRepo) *App {
	return &App{
		cfg:              cfg,
		seAccountInfoSvc: saltedge.NewSeAccountInfoSvc(cfg),
		fvAuthSvc:        finverse.NewFvAuthSvc(cfg, repo),
		fvDataSvc:        finverse.NewFvDataSvc(cfg, repo),
	}
}

func (a App) GetSeAccountInfoSvc() *saltedge.SeAccountInfoSvc {
	return a.seAccountInfoSvc
}

func (a App) GetFvAuthSvc() *finverse.FvAuthSvc {
	return a.fvAuthSvc
}

func (a App) GetFvDataSvc() *finverse.FvDataSvc {
	return a.fvDataSvc
}
