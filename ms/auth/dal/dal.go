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
	GetUserRepo() IUserRepo
	GetWebauthnCredentialRepo() IWebauthnCredentialRepo
	GetWebauthnSessionRepo() IWebauthnSessionRepo
	GetEmailRepo() IEmailRepo
	GetPasscodeRepo() IPasscodeRepo
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

func (r AuthRepo) GetUserRepo() IUserRepo {
	return r.userRepo
}

func (r AuthRepo) GetWebauthnCredentialRepo() IWebauthnCredentialRepo {
	return r.waCredentialRepo
}

func (r AuthRepo) GetWebauthnSessionRepo() IWebauthnSessionRepo {
	return r.waSessionRepo
}

func (r AuthRepo) GetEmailRepo() IEmailRepo {
	return r.emailRepo
}

func (r AuthRepo) GetPasscodeRepo() IPasscodeRepo {
	return r.passcodeRepo
}
