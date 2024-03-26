package db

import (
	"context"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/internal/db/pgsql"
	"github.com/NoahJinnn/passkey_auth_svc/internal/sharedconfig"
	"github.com/powerman/pqx"
)

type Db struct {
	PgEnt *ent.Client
	// We can declare multiple clients here, e.g: MySQLClient *ent.Client, SQLiteClient *ent.Client
}

func InitDbClient(ctxStartupCmdServe context.Context, cfg *sharedconfig.Shared) *Db {
	cfg.Postgres.SSLMode = pqx.SSLRequire
	dateSourceName := cfg.Postgres.FormatDSN()
	pgEnt := pgsql.NewPgEnt(dateSourceName)
	return &Db{
		PgEnt: pgEnt,
	}
}
