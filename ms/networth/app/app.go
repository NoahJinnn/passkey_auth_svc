// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/app/ws"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetFvAuthSvc() *finverse.FvAuthSvc
	GetFvDataSvc() *finverse.FvDataSvc
	GetSeAccountInfoSvc() *saltedge.SeAccountInfoSvc
	GetChangesetSvc() *ws.ChangesetSvc
}

// App implements interface Appl.
type App struct {
	cfg              *config.Config
	seAccountInfoSvc *saltedge.SeAccountInfoSvc
	fvAuthSvc        *finverse.FvAuthSvc
	fvDataSvc        *finverse.FvDataSvc
	csSvc            *ws.ChangesetSvc
}

// New creates and returns new App.
func New(cfg *config.Config, repo dal.INwRepo) *App {
	return &App{
		cfg:              cfg,
		seAccountInfoSvc: saltedge.NewSeAccountInfoSvc(cfg),
		fvAuthSvc:        finverse.NewFvAuthSvc(cfg, repo),
		fvDataSvc:        finverse.NewFvDataSvc(cfg, repo),
		csSvc:            ws.NewChangesetSvc(repo),
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

func (a App) GetChangesetSvc() *ws.ChangesetSvc {
	return a.csSvc
}
