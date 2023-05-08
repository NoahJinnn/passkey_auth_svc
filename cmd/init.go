package cmd

import (
	"context"
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/sharedConfig"
	"github.com/hellohq/hqservice/internal/sharedDal"
	"github.com/powerman/pqx"
)

func InitEntClient(ctxStartupCmdServe context.Context, cfg *sharedConfig.Shared) *ent.Client {
	cfg.Postgres.SSLMode = pqx.SSLRequire
	dateSourceName := cfg.Postgres.FormatDSN()
	entClient := sharedDal.CreateEntClient(ctxStartupCmdServe, dateSourceName)
	return entClient
}

func InitSessionManager(ctxStartupCmdServe context.Context, cfg *sharedConfig.Shared, repo sharedDal.IJwkRepo) session.Manager {
	jwkManager, err := session.NewDefaultManager(cfg.Secrets.Keys, repo)
	if err != nil {
		panic(fmt.Errorf("failed to create jwk manager: %w", err))
	}
	sessionManager, err := session.NewManager(jwkManager, cfg.Session)
	if err != nil {
		panic(fmt.Errorf("failed to create session generator: %w", err))
	}
	return sessionManager
}
