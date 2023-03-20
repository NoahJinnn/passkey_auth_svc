package config

import "github.com/hellohq/hqservice/pkg/netx"

type Server struct {
	BindAddr        netx.Addr
	BindAddrInt     netx.Addr
	BindMetricsAddr netx.Addr
}
