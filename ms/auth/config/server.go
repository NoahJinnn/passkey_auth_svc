package config

import "github.com/NoahJinnn/passkey_auth_svc/pkg/netx"

type Server struct {
	BindAddr        netx.Addr
	BindAddrInt     netx.Addr
	BindMetricsAddr netx.Addr
	Cors            Cors
}

type Cors struct {
	Enabled          bool
	AllowCredentials bool     `split_words:"true"`
	AllowOrigins     []string `split_words:"true"`
	AllowMethods     []string `split_words:"true"`
	AllowHeaders     []string `split_words:"true"`
	ExposeHeaders    []string `split_words:"true"`
	MaxAge           int      `split_words:"true"`
}
