package config

type SaltEdgeConfig struct {
	AppId  string `env:"SALTEDGE_APP_ID"`
	Secret string `env:"SALTEDGE_SECRET"`
	PK     string `env:"SALTEDGE_PK"`
}
