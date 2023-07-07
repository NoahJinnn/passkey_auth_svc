// Package app provides business logic.
package app

import (
	"github.com/hellohq/hqservice/ms/networth/app/asset"
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/hellohq/hqservice/ms/networth/app/provider"
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// Appl provides application features (use cases) service.
type Appl interface {
	GetAssetSvc() *asset.AssetSvc
	GetProviderSvc() *provider.ProviderSvc
	GetFvAuthSvc() *finverse.FvAuthSvc
	GetFvDataSvc() *finverse.FvDataSvc
	GetSeAccountInfoSvc() *saltedge.SeAccountInfoSvc
}

// App implements interface Appl.
type App struct {
	cfg              *config.Config
	repo             *dal.NwRepo
	assetSvc         *asset.AssetSvc
	providerSvc      *provider.ProviderSvc
	seAccountInfoSvc *saltedge.SeAccountInfoSvc
	fvAuthSvc        *finverse.FvAuthSvc
	fvDataSvc        *finverse.FvDataSvc
}

// New creates and returns new App.
func New(cfg *config.Config, repo *dal.NwRepo) *App {
	assetSvc := asset.NewAssetSvc(cfg, repo)
	providerSvc := provider.NewProviderSvc()
	seAccountInfoSvc := saltedge.NewSeAccountInfoSvc(cfg)
	fvAuthSvc := finverse.NewFvAuthSvc(cfg, repo)
	fvDataSvc := finverse.NewFvDataSvc(cfg, repo)

	return &App{
		cfg:              cfg,
		repo:             repo,
		assetSvc:         assetSvc,
		providerSvc:      providerSvc,
		seAccountInfoSvc: seAccountInfoSvc,
		fvAuthSvc:        fvAuthSvc,
		fvDataSvc:        fvDataSvc,
	}
}

func (a App) GetAssetSvc() *asset.AssetSvc {
	return a.assetSvc
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
