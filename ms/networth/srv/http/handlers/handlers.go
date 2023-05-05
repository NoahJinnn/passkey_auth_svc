package handlers

import (
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
)

type HttpDeps struct {
	app.Appl
	Cfg *config.Config
}
