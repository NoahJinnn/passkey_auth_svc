// Package mono provides embedded microservice.
package srv

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"regexp"

	"github.com/hellohq/hqservice/api/openapi/restapi"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/hq/app"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/ms/hq/dal"
	"github.com/hellohq/hqservice/ms/hq/srv/openapi"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/hellohq/hqservice/pkg/netx"
	"github.com/hellohq/hqservice/pkg/serve"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

var reg = prometheus.NewPedanticRegistry()

// Service implements main.embeddedService interface.
type Service struct {
	cfg     *config.Config
	srv     *restapi.Server
	appl    *app.App
	repo    *dal.Repo
	ca      *x509.CertPool
	cert    tls.Certificate
	certInt tls.Certificate
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return config.ServiceName }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *sharedconfig.Shared, _, serveCmd *cobra.Command) error {
	namespace := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(def.ProgName, "_")
	initMetrics(reg, namespace)
	openapi.InitMetrics(reg, namespace)

	return config.Init(s.Name(), serveCmd.Flags(), sharedCfg)
}

// RunServe implements main.embeddedService interface.
func (s *Service) RunServe(ctxStartup Ctx, ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)

	if s.cfg == nil {
		s.cfg, err = config.GetServe()
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}

	if err == nil {
		s.ca, err = netx.LoadCACert(s.cfg.TLSCACert)
	}
	if err == nil {
		s.cert, err = tls.LoadX509KeyPair(s.cfg.TLSCert, s.cfg.TLSKey)
	}
	if err == nil {
		s.certInt, err = tls.LoadX509KeyPair(s.cfg.TLSCertInt, s.cfg.TLSKeyInt)
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}

	err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo,
	})
	if err != nil {
		return log.Err("failed to connect", "err", err)
	}

	if s.appl == nil {
		s.appl, err = app.New()
		if err != nil {
			return log.Err("failed to create Appl", "err", err)
		}
	}

	s.srv, err = openapi.NewServer(s.appl, openapi.Config{
		Addr: s.cfg.BindAddr,
	})
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveMetrics,
		s.serveOpenAPI,
	)

	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *Service) serveMetrics(ctx Ctx) error {
	return serve.Metrics(ctx, s.cfg.BindMetricsAddr, reg)
}

func (s *Service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}

func (s *Service) connectRepo(ctx Ctx) (interface{}, error) {
	return dal.New(ctx, s.cfg.Postgres)
}
