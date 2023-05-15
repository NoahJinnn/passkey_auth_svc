// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/pgsql"
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
	Db *ent.Client
}
type Ctx = context.Context

func New(client *ent.Client) *AuthRepo {
	return &AuthRepo{
		Db: client,
	}
}

func (r AuthRepo) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	return pgsql.WithTx(ctx, r.Db, exec)
}

func (r AuthRepo) GetUserRepo() IUserRepo {
	return NewUserRepo(r.Db)
}

func (r AuthRepo) GetWebauthnCredentialRepo() IWebauthnCredentialRepo {
	return NewWebauthnCredentialRepo(r.Db)
}

func (r AuthRepo) GetWebauthnSessionRepo() IWebauthnSessionRepo {
	return NewWebauthnSessionRepo(r.Db)
}

func (r AuthRepo) GetEmailRepo() IEmailRepo {
	return NewEmailRepo(r.Db)
}

func (r AuthRepo) GetPasscodeRepo() IPasscodeRepo {
	return NewPasscodeRepo(r.Db)
}
