// Package auth provides auth service.
package auth

import (
	"context"
	"fmt"

	"github.com/NoahJinnn/passkey_auth_svc/internal/db"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/session"
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/app"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	server "github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/mail"
	"github.com/NoahJinnn/passkey_auth_svc/pkg/concurrent"
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
	repo           *dal.AuthRepo
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return "auth" }

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

	mailer := mail.NewMailer(&s.cfg.Passcode)
	renderer, err := mail.NewRenderer()
	if err != nil {
		panic(fmt.Errorf("failed to create new renderer: %w", err))
	}

	s.sessionManager = sessionManager
	s.repo = dal.New(dbClient)
	s.appl = app.New(mailer, renderer, s.cfg, s.repo)

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
	e.Logger.Fatal(e.Start(s.cfg.Server.BindAddr.String()))
	return err
}
