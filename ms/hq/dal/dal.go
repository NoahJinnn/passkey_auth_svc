// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"
	"os"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/powerman/pqx"
	"github.com/powerman/structlog"
)

// Error names.
const (
	PostgresUniqueViolation     = "unique_violation"
	PostgresForeignKeyViolation = "foreign_key_violation"
	schemaVersion               = 4
	dbMaxOpenConns              = 100 / 10 // Use up to 1/10 of server's max_connections.
	dbMaxIdleConns              = 5        // A bit more than default (2).
)

type Repo struct {
	db  *ent.Client
	log *structlog.Logger
}
type Ctx = context.Context

func New(ctx Ctx, cfg *config.PostgresConfig) (_ *Repo, err error) {
	log := structlog.FromContext(ctx, nil)
	cfg.SSLMode = pqx.SSLRequire
	dateSourceName := cfg.FormatDSN()

	client, err := ent.Open("postgres", dateSourceName)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	r := &Repo{
		db:  client,
		log: log,
	}
	return r, nil
}
