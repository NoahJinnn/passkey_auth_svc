package sqlite

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteClient(ctx context.Context, dateSourceName string) *ent.Client {
	client, err := ent.Open(dialect.SQLite, dateSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
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
