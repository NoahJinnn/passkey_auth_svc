package pgsql

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ent/migrate"
)

func NewPgEnt(dsn string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
