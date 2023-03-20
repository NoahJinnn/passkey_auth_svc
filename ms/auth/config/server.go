package config

import "github.com/hellohq/hqservice/pkg/netx"

type Server struct {
	BindAddr        netx.Addr
	BindAddrInt     netx.Addr
	BindMetricsAddr netx.Addr
	Cors            Cors `yaml:"cors" json:"cors" koanf:"cors"`
}

type Cors struct {
	Enabled          bool     `yaml:"enabled" json:"enabled" koanf:"enabled"`
	AllowCredentials bool     `yaml:"allow_credentials" json:"allow_credentials" koanf:"allow_credentials" split_words:"true"`
	AllowOrigins     []string `yaml:"allow_origins" json:"allow_origins" koanf:"allow_origins" split_words:"true"`
	AllowMethods     []string `yaml:"allow_methods" json:"allow_methods" koanf:"allow_methods" split_words:"true"`
	AllowHeaders     []string `yaml:"allow_headers" json:"allow_headers" koanf:"allow_headers" split_words:"true"`
	ExposeHeaders    []string `yaml:"expose_headers" json:"expose_headers" koanf:"expose_headers" split_words:"true"`
	MaxAge           int      `yaml:"max_age" json:"max_age" koanf:"max_age" split_words:"true"`
}
