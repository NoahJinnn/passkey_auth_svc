//go:build integration
// +build integration

package hq

import (
	"context"
	"testing"

	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/powerman/check"
)

func TestSmoke(tt *testing.T) {
	t := check.T(tt)

	s := &Service{}
	cfg := &config.Config{}
	s.cfg = cfg
	const host = "localhost"
	s.cfg.BindAddr = netx.NewAddr(host, netx.UnusedTCPPort(host))

	ctxStartup, cancel := context.WithTimeout(ctx, def.TestTimeout)
	defer cancel()
	ctxShutdown, shutdown := context.WithCancel(ctx)
	errc := make(chan error)
	go func() { errc <- s.RunServe(ctxStartup, ctxShutdown, shutdown) }()
	defer func() {
		shutdown()
		t.Nil(<-errc, "RunServe")
	}()
	t.Must(t.Nil(netx.WaitTCPPort(ctxStartup, s.cfg.BindAddr), "connect to HTTP service"))
}
