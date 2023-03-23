// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"
	"fmt"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/migrate"
	"github.com/hellohq/hqservice/ms/auth/config"
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

// Repo provides data storage.
type IRepo interface {
	WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error
	GetJwkRepo() IJwkRepo
	GetUserRepo() IUserRepo
	GetWebauthnCredentialRepo() IWebauthnCredentialRepo
	GetWebauthnSessionRepo() IWebauthnSessionRepo
}

type Repo struct {
	Db  *ent.Client
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

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Repo{
		Db:  client,
		log: log,
	}, nil
}

func (r Repo) Close() {
	r.Db.Close()
}

func (r Repo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.Db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (r Repo) GetJwkRepo() IJwkRepo {
	return NewJwkRepo(r.Db)
}

func (r Repo) GetUserRepo() IUserRepo {
	return NewUserRepo(r.Db)
}

func (r Repo) GetWebauthnCredentialRepo() IWebauthnCredentialRepo {
	return NewWebauthnCredentialRepo(r.Db)
}

func (r Repo) GetWebauthnSessionRepo() IWebauthnSessionRepo {
	return NewWebauthnSessionRepo(r.Db)
}
