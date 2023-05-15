package handlers

import (
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
)

type HttpDeps struct {
	app.Appl
	Cfg       *config.Config
	SharedCfg *sharedconfig.Shared
}
