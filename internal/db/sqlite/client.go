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
	"github.com/mattn/go-sqlite3"
)

func NewSqliteClient(dsn string) *ent.Client {
	db := NewSqliteDrive(dsn)
	drv := entsql.OpenDB(dialect.SQLite, db)

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

func NewSqliteDrive(dsn string) *sql.DB {
	db, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatalf("failed getting connection: %v", err)
	}
	defer conn.Close()

	err = conn.Raw(func(driverConn interface{}) error {
		sqliteConn := driverConn.(*sqlite3.SQLiteConn)
		err := sqliteConn.LoadExtension("crsqlite", "sqlite3_crsqlite_init")
		if err != nil {
			return err
		}
		fmt.Println("load extension success")
		return nil
	})
	if err != nil {
		log.Fatalf("failed loading extension: %v", err)
	}

	r := conn.QueryRowContext(context.Background(), "SELECT quote(crsql_siteid());")
	var siteid string
	if err = r.Scan(&siteid); err == sql.ErrNoRows {
		log.Printf("Id not found")
		log.Fatalf("failed to query crsql lite id: %v", err)
	}

	fmt.Printf("r: %v\n", siteid)

	return db
}
