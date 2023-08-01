package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDrive(dsn string) *sql.DB {
	db, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return db
}

func NewSqliteEnt(ctx context.Context, db *sql.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.SQLite, db)
	client := ent.NewClient(ent.Driver(drv))
	// Run the auto migration tool.
	if err := client.Debug().Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewSqliteConn(ctx context.Context, db *sql.DB) *sql.Conn {
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatalf("failed getting connection: %v", err)
	}

	err = conn.Raw(func(driverConn interface{}) error {
		fmt.Println("load extension success")
		return nil
	})
	if err != nil {
		log.Fatalf("failed loading extension: %v", err)
	}
	return conn
}
