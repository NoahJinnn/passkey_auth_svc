// Package auth provides networth service.
package networth

import (
	"context"

	"github.com/hellohq/hqservice/internal/db"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/networth/app"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	server "github.com/hellohq/hqservice/ms/networth/srv/http"
	"github.com/hellohq/hqservice/pkg/concurrent"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Service implements main.embeddedService interface.
type Service struct {
	cfg            *config.Config
	sharedCfg      *sharedconfig.Shared
	sessionManager *session.Manager
	appl           *app.App
	repo           *dal.NwRepo
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return "networth" }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *sharedconfig.Shared, serveCmd *cobra.Command) error {
	s.sharedCfg = sharedCfg
	return config.Init(sharedCfg, config.FlagSets{
		Serve: serveCmd.Flags(),
	})
}

// RunServe implements main.embeddedService interface.
func (s *Service) RunServe(ctxStartup Ctx, ctxShutdown Ctx, shutdown func(), dbClient *db.Db, sessionManager *session.Manager) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)

	if s.cfg == nil {
		s.cfg, err = config.GetServe()
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}
	s.sessionManager = sessionManager
	s.repo = dal.New(dbClient)
	s.appl = app.New(s.cfg, s.repo)

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveEcho,
	)

	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *Service) serveEcho(ctx Ctx) error {
	e, err := server.NewServer(s.appl, s.sessionManager, s.sharedCfg, s.cfg)
	e.Logger.Fatal(e.StartTLS(s.cfg.Server.BindAddr.String(), "configs/http-pki/cacert.pem", "configs/http-pki/cakey.pem"))
	return err
}
