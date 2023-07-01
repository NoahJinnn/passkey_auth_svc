package db

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/db/pgsql"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/powerman/pqx"
)

type DbClient struct {
	PgClient *ent.Client
}

func InitDbClient(ctxStartupCmdServe context.Context, cfg *sharedconfig.Shared) *DbClient {
	cfg.Postgres.SSLMode = pqx.SSLRequire
	dateSourceName := cfg.Postgres.FormatDSN()
	PgClient := pgsql.NewPgClient(ctxStartupCmdServe, dateSourceName)

	return &DbClient{
		PgClient: PgClient,
	}
}
