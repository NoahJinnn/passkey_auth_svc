package pgsql

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/migrate"
)

func NewPgClient(ctx context.Context, dateSourceName string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, dateSourceName)
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

	return client
}
