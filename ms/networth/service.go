// Package hq provides embedded microservice.
package hq

import (
	"context"
	"net/http"

	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	server "github.com/hellohq/hqservice/ms/networth/srv/http"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/labstack/echo/v4"
	"github.com/powerman/pqx"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Service implements main.embeddedService interface.
type Service struct {
	cfg  *config.Config
	srv  *echo.Echo
	appl app.App
	repo *dal.Repo
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return config.ServiceName }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *sharedConfig.Shared, _, serveCmd *cobra.Command) error {
	return config.Init(sharedCfg, config.FlagSets{
		Serve: serveCmd.Flags(),
	})
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

	err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo,
	})
	if err != nil {
		return log.Err("failed to connect", "err", err)
	}
	s.appl = app.New(s.cfg, s.repo)

	s.srv, err = server.NewServer(s.appl, *s.repo, s.cfg)
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveEcho,
	)

	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *Service) serveEcho(ctx Ctx) error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e.Start(":1323")
}

func (s *Service) connectRepo(ctx Ctx) (interface{}, error) {
	s.cfg.Postgres.SSLMode = pqx.SSLRequire
	dateSourceName := s.cfg.Postgres.FormatDSN()
	return dal.New(ctx, dateSourceName)
}
