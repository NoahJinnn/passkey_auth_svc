package sqlite

import (
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteClient(dateSourceName string) *ent.Client {
	client, err := ent.Open(dialect.SQLite, dateSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return client
}
