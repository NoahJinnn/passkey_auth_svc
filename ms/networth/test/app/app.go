package test

import (
	"github.com/hellohq/hqservice/ms/networth/app/svcs"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// App implements interface Appl.
type app struct {
	cfg  *config.Config
	repo dal.INwRepo
}

// New creates and returns new App.
func NewApp(cfg *config.Config, repo dal.INwRepo) app {
	return app{
		cfg:  cfg,
		repo: repo,
	}
}

func (a app) GetSeAccountInfoSvc() svcs.ISeAccountInfoSvc {
	return svcs.NewSeAccountInfoSvc(a.cfg)
}
