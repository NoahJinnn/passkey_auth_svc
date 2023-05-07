// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/sharedDal"
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
	WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error
	GetJwkRepo() IJwkRepo
	GetEmailRepo() IEmailRepo
}

type Repo struct {
	Db  *ent.Client
	log *structlog.Logger
}
type Ctx = context.Context

func New(client *ent.Client) *Repo {
	return &Repo{
		Db: client,
	}
}

func (r Repo) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	return sharedDal.WithTx(ctx, r.Db, exec)
}

func (r Repo) GetJwkRepo() IJwkRepo {
	return NewJwkRepo(r.Db)
}

func (r Repo) GetEmailRepo() IEmailRepo {
	return NewEmailRepo(r.Db)
}
