package handlers

import (
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
)

type HttpDeps struct {
	app.Appl
	Cfg       *config.Config
	SharedCfg *sharedconfig.Shared
}
