// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/migrate"
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
	GetUserRepo() IUserRepo
	GetWebauthnCredentialRepo() IWebauthnCredentialRepo
	GetWebauthnSessionRepo() IWebauthnSessionRepo
	GetEmailRepo() IEmailRepo
	GetPasscodeRepo() IPasscodeRepo
}

type Repo struct {
	Db  *ent.Client
	log *structlog.Logger
}
type Ctx = context.Context

func New(ctx Ctx, dateSourceName string) (_ *Repo, err error) {
	// TODO: Move ent client initialization to main.go
	log := structlog.FromContext(ctx, nil)
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

func (r Repo) WithTx(ctx context.Context, exec func(ctx Ctx, client *ent.Client) error) error {
	return sharedDal.WithTx(ctx, r.Db, exec)
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

func (r Repo) GetEmailRepo() IEmailRepo {
	return NewEmailRepo(r.Db)
}

func (r Repo) GetPasscodeRepo() IPasscodeRepo {
	return NewPasscodeRepo(r.Db)
}
