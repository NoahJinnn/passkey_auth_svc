package config

import "github.com/powerman/appcfg"

type SaltEdgeConfig struct {
	AppId  appcfg.String `env:"SALTEDGE_APP_ID"`
	Secret appcfg.String `env:"SALTEDGE_SECRET"`
}
