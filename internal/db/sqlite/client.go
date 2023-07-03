package sqlite

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/ent"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteClient(ctx context.Context, dateSourceName string) *ent.Client {
	client, err := ent.Open(dialect.SQLite, dateSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return client
}
