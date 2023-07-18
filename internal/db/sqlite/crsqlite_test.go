package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/stretchr/testify/suite"
)

type Todo struct {
	ID        string
	ListId    int
	Text      string
	Completed bool
}
type sqliteTestSuite struct {
	suite.Suite
	ctx     context.Context
	conns   []*sql.Conn
	clients []*ent.Client
}

func (s *sqliteTestSuite) SetupTest() {
	ctx := context.Background()
	s.ctx = ctx
	for i := 0; i < 3; i++ {
		db := NewSqliteDrive("file:" + fmt.Sprintf("test %d", i) + ".db?cache=shared&_fk=1")
		client := NewSqliteEnt(db)
		s.clients = append(s.clients, client)
		conn := NewSqliteConn(db)
		s.conns = append(s.conns, conn)
	}
}

func (s *sqliteTestSuite) TearDownTest() {
	for _, conn := range s.conns {
		conn.ExecContext(s.ctx, `SELECT crsql_finalize();`)
		conn.Close()
	}
}

func TestSqliteTestSuite(t *testing.T) {
	suite.Run(t, new(sqliteTestSuite))
}

func (s *sqliteTestSuite) TestMerge() {
	todos := make([]Todo, gofakeit.IntRange(4, 20))
	for i := range todos {
		todos[i] = Todo{
			ID:        gofakeit.UUID(),
			ListId:    gofakeit.IntRange(1, 1000),
			Text:      gofakeit.Sentence(3),
			Completed: gofakeit.Bool(),
		}
	}

	for _, todo := range todos {
		err := s.clients[0].Todo.Create().
			SetID(uuid.FromStringOrNil(todo.ID)).
			SetListId(todo.ListId).
			SetText(todo.Text).
			SetCompleted(todo.Completed).
			Exec(s.ctx)
		if err != nil {
			fmt.Println("Error insert todo:", err)
			return
		}
	}

	for _, conn := range s.conns {
		_, err := conn.ExecContext(s.ctx, `SELECT crsql_as_crr('todos')`)
		s.NoError(err)
	}

	sync(s.ctx, s.conns[0], s.conns[1])
	sync(s.ctx, s.conns[0], s.conns[2])
	s.True(assertAllRowsByConn(s.ctx, s.conns), "All rows should be equal, DBs sync failed")
	sync(s.ctx, s.conns[1], s.conns[0])
	sync(s.ctx, s.conns[2], s.conns[0])
	assertAllRowsByConn(s.ctx, s.conns)
	s.True(assertAllRowsByConn(s.ctx, s.conns), "All rows should be equal, DBs sync failed")
}

func queryAllRows(ctx context.Context, db *sql.Conn, q []interface{}) []map[string]interface{} {
	query := q[0].(string)
	params := q[1].([]interface{})
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting column names:", err)
		return nil
	}

	result := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range values {
			values[i] = new(interface{})
		}
		err = rows.Scan(values...)
		if err != nil {
			fmt.Println("Error scanning row values:", err)
			return nil
		}
		rowData := make(map[string]interface{})
		for i, column := range columns {
			rowData[column] = *values[i].(*interface{})
		}
		result = append(result, rowData)
	}

	return result
}

func sync(ctx context.Context, left *sql.Conn, right *sql.Conn) {
	changesets, err := left.QueryContext(ctx, `SELECT * FROM crsql_changes WHERE db_version > 0 AND site_id IS NULL`)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := right.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`INSERT INTO crsql_changes VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		_, err := stmt.Exec(table, pk, cid, val, col_version, db_version, site_id)
		if err != nil {
			fmt.Println("Error executing statement:", err)
			fmt.Println(table, pk, cid, val, col_version, db_version, site_id)
			return
		}
	}

	defer stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func assertAllRowsByConn(ctx context.Context, conns []*sql.Conn) bool {
	results := make([][]map[string]interface{}, len(conns))

	for i, db := range conns {
		query := `SELECT * FROM todos ORDER BY id DESC`
		results[i] = queryAllRows(ctx, db, []interface{}{query, []interface{}{}})
	}

	for i := 0; i < len(results)-1; i++ {
		if !isEqual(results[i], results[i+1]) {
			fmt.Println("Assertion failed!")
			fmt.Println("Results of", i)
			for _, rows := range results[i] {
				fmt.Println(rows)
			}
			fmt.Println("Results of", i+1)
			for _, rows := range results[i+1] {
				fmt.Println(rows)
			}

			return false
		}
	}

	return true
}

func isEqual(a, b []map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !isRowEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func isRowEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valA := range a {
		valB, ok := b[key]
		if !ok {
			return false
		}
		if valA != valB {
			return false
		}
	}

	return true
}
