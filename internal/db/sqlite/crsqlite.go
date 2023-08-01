package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

// TODO: Get latest "db_version", "site_id" from remote to optimize get changesets with condition WHERE "db_version" > ? AND "site_id" IS NOT ?
func GetCrrChanges(ctx context.Context, conn *sql.Conn) (*sql.Rows, error) {
	changesets, err := conn.QueryContext(ctx, `SELECT * FROM crsql_changes WHERE db_version > 0 AND site_id IS NULL`)
	if err != nil {
		return nil, err
	}
	return changesets, nil
}

func ApplyCrrChanges(ctx context.Context, conn *sql.Conn, changesets *sql.Rows) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`INSERT INTO crsql_changes VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	// [table] TEXT NOT NULL,
	// [pk] TEXT NOT NULL,
	// [cid] TEXT NOT NULL,
	// [val] ANY,
	// [col_version] INTEGER NOT NULL,
	// [db_version] INTEGER NOT NULL,
	// [site_id] BLOB --
	for changesets.Next() {
		var table string
		var pk string
		var cid string
		var val interface{}
		var col_version int
		var db_version int
		var site_id []byte
		err = changesets.Scan(&table, &pk, &cid, &val, &col_version, &db_version, &site_id)
		if err != nil {
			return err
		}
		_, err := stmt.Exec(table, pk, cid, val, col_version, db_version, site_id)
		if err != nil {
			fmt.Println("Error executing statement:", err)
			return err
		}
	}

	defer stmt.Close()
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
