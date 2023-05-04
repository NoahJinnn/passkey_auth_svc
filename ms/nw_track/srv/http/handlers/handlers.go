package handlers

import (
	"github.com/hellohq/hqservice/ms/nw_track/app"
	"github.com/hellohq/hqservice/ms/nw_track/config"
)

type HttpDeps struct {
	app.Appl
	Cfg *config.Config
}
