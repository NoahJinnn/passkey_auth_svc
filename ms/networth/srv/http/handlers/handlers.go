package handlers

import (
	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
)

type HttpDeps struct {
	app.Appl
	Cfg       *config.Config
	SharedCfg *sharedConfig.Shared
}
