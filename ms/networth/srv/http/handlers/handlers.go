package handlers

import (
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
)

type HttpDeps struct {
	app.Appl
	Cfg       *config.Config
	SharedCfg *sharedconfig.Shared
}
