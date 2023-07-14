// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/hellohq/hqservice/ms/networth/app/item_table"
	"github.com/hellohq/hqservice/ms/networth/app/provider"
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetItemTableSvc() *item_table.ItemTableSvc
	GetProviderSvc() *provider.ProviderSvc
	GetFvAuthSvc() *finverse.FvAuthSvc
	GetFvDataSvc() *finverse.FvDataSvc
	GetSeAccountInfoSvc() *saltedge.SeAccountInfoSvc
}

// App implements interface Appl.
type App struct {
	cfg              *config.Config
	itemTableSvc     *item_table.ItemTableSvc
	providerSvc      *provider.ProviderSvc
	seAccountInfoSvc *saltedge.SeAccountInfoSvc
	fvAuthSvc        *finverse.FvAuthSvc
	fvDataSvc        *finverse.FvDataSvc
}

// New creates and returns new App.
func New(cfg *config.Config, repo dal.INwRepo) *App {
	providerSvc := provider.NewProviderSvc()
	itemTableSvc := item_table.NewItemTableSvc(cfg, repo)
	seAccountInfoSvc := saltedge.NewSeAccountInfoSvc(cfg)
	fvAuthSvc := finverse.NewFvAuthSvc(cfg, providerSvc, repo)
	fvDataSvc := finverse.NewFvDataSvc(cfg, providerSvc, repo)

	return &App{
		cfg:              cfg,
		itemTableSvc:     itemTableSvc,
		providerSvc:      providerSvc,
		seAccountInfoSvc: seAccountInfoSvc,
		fvAuthSvc:        fvAuthSvc,
		fvDataSvc:        fvDataSvc,
	}
}

func (a App) GetItemTableSvc() *item_table.ItemTableSvc {
	return a.itemTableSvc
}

func (a App) GetProviderSvc() *provider.ProviderSvc {
	return a.providerSvc
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
