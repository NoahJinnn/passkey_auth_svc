// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/db"
	"github.com/hellohq/hqservice/internal/db/pgsql"
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
type IAuthRepo interface {
	WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error
	GetUserRepo() *userRepo
	GetWebauthnCredentialRepo() *waCredentialRepo
	GetWebauthnSessionRepo() *waSessionRepo
	GetEmailRepo() *emailRepo
	GetPasscodeRepo() *passcodeRepo
}

type AuthRepo struct {
	Db               *db.Db
	userRepo         *userRepo
	waCredentialRepo *waCredentialRepo
	waSessionRepo    *waSessionRepo
	emailRepo        *emailRepo
	passcodeRepo     *passcodeRepo
}
type Ctx = context.Context

func New(client *db.Db) *AuthRepo {
	userRepo := NewUserRepo(client.PgClient)
	waCredentialRepo := NewWebauthnCredentialRepo(client.PgClient)
	waSessionRepo := NewWebauthnSessionRepo(client.PgClient)
	emailRepo := NewEmailRepo(client.PgClient)
	passcodeRepo := NewPasscodeRepo(client.PgClient)
	return &AuthRepo{
		Db:               client,
		userRepo:         userRepo,
		waCredentialRepo: waCredentialRepo,
		waSessionRepo:    waSessionRepo,
		emailRepo:        emailRepo,
		passcodeRepo:     passcodeRepo,
	}
}

func (r AuthRepo) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	return pgsql.WithTx(ctx, r.Db.PgClient, exec)
}

func (r AuthRepo) GetUserRepo() *userRepo {
	return r.userRepo
}

func (r AuthRepo) GetWebauthnCredentialRepo() *waCredentialRepo {
	return r.waCredentialRepo
}

func (r AuthRepo) GetWebauthnSessionRepo() *waSessionRepo {
	return r.waSessionRepo
}

func (r AuthRepo) GetEmailRepo() *emailRepo {
	return r.emailRepo
}

func (r AuthRepo) GetPasscodeRepo() *passcodeRepo {
	return r.passcodeRepo
}
