package sqlite

import (
	"context"
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/migrate"
	"github.com/mattn/go-sqlite3"
)

func NewSqliteClient(dsn string) *ent.Client {
	db, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer conn.Close()
	err = conn.Raw(func(driverConn interface{}) error {
		sqliteConn := driverConn.(*sqlite3.SQLiteConn)
		sqliteConn.LoadExtension("crsqlite-darwin-aarch64", "sqlite3_crsqlite_init")
		return nil
	})
	drv := entsql.OpenDB(dialect.SQLite, db)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
