package test

import (
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

// App implements interface Appl.
type appT struct {
	cfg  *config.Config
	repo dal.INwRepo
}

// New creates and returns new App.
func NewApp(cfg *config.Config, repo dal.INwRepo) appT {
	return appT{
		cfg:  cfg,
		repo: repo,
	}
}

func (a appT) GetSeAccountInfoSvc() saltedge.ISeAccountInfoSvc {
	return saltedge.NewSeAccountInfoSvc(a.cfg)
}
