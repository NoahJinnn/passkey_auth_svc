// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/db"
	"github.com/hellohq/hqservice/internal/db/pgsql"
)

type Ctx = context.Context

// Error names.
const (
	PostgresUniqueViolation     = "unique_violation"
	PostgresForeignKeyViolation = "foreign_key_violation"
	schemaVersion               = 4
	dbMaxOpenConns              = 100 / 10 // Use up to 1/10 of server's max_connections.
	dbMaxIdleConns              = 5        // A bit more than default (2).
)

// Repo provides data storage.
type INwRepo interface {
	WithTx(ctx Ctx, exec func(ctx Ctx, client *ent.Client) error) error
}

type NwRepo struct {
	Db *db.DbClient
}

func New(client *db.DbClient) *NwRepo {
	return &NwRepo{
		Db: client,
	}
}

func (r NwRepo) WithTx(ctx Ctx, exec func(ctx Ctx, client *ent.Client) error) error {
	return pgsql.WithTx(ctx, r.Db.PgClient, exec)
}
