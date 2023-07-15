package sqlite

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteClient(dsn string) *ent.Client {
	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// client.

	return client
}
