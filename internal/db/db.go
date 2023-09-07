package db

import (
	"context"
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/db/pgsql"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/powerman/pqx"
)

type Db struct {
	PgEnt *ent.Client
	// We can declare multiple clients here, e.g: MySQLClient *ent.Client, SQLiteClient *ent.Client
}

func InitDbClient(ctxStartupCmdServe context.Context, cfg *sharedconfig.Shared) *Db {
	cfg.Postgres.SSLMode = pqx.SSLRequire
	// TODO: Set the following values to env
	cfg.Postgres.SSLCert = "../configs/pg-pki/client-cert.pem"
	cfg.Postgres.SSLKey = "../configs/pg-pki/client-key.pem"
	dateSourceName := cfg.Postgres.FormatDSN()
	fmt.Println(dateSourceName)
	pgEnt := pgsql.NewPgEnt(dateSourceName)
	return &Db{
		PgEnt: pgEnt,
	}
}
