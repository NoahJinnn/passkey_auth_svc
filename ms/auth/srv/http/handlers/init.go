package handlers

import (
	"github.com/hellohq/hqservice/ms/auth/app"
	"github.com/hellohq/hqservice/ms/auth/config"
)

type HttpDeps struct {
	App app.Appl
	Cfg *config.Config
}
