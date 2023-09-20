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
	GetFvSessionRepo() IFvSessionRepo
}

type NwRepo struct {
	Db            *db.Db
	fvSessionRepo *fvSessionRepo
}

func New(client *db.Db) *NwRepo {
	return &NwRepo{
		Db:            client,
		fvSessionRepo: NewFvSessionRepo(client.PgEnt),
	}
}

func (r NwRepo) WithTx(ctx Ctx, exec func(ctx Ctx, client *ent.Client) error) error {
	return pgsql.WithTx(ctx, r.Db.PgEnt, exec)
}

func (r NwRepo) GetFvSessionRepo() IFvSessionRepo {
	if r.fvSessionRepo == nil {
		r.fvSessionRepo = NewFvSessionRepo(r.Db.PgEnt)
	}
	return r.fvSessionRepo
}
