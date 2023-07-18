package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type Todo struct {
	ID        int
	Text      string
	Completed bool
}
type sqliteTestSuite struct {
	suite.Suite
	conns []*sql.Conn
}

func (s *sqliteTestSuite) SetupTest() {
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		db := NewSqliteDrive("file:" + fmt.Sprintf("test %d", i) + "file:ent?mode=memory&_fk=1")
		conn := NewSqliteConn(db)
		conn.ExecContext(ctx, `CREATE TABLE todo ("id" primary key, "listId", "completed", "text")`)
		conn.ExecContext(ctx, `SELECT crsql_as_crr('todo')`)
		s.conns = append(s.conns, conn)
	}
}

func (s *sqliteTestSuite) TearDownTest() {
	for _, conn := range s.conns {
		conn.Close()
	}
}

func TestSqliteTestSuite(t *testing.T) {
	suite.Run(t, new(sqliteTestSuite))
}

func createInsert(id interface{}, listId interface{}, text string, completed bool) (string, []interface{}) {
	query := `INSERT INTO todo ("id", "listId", "text", "completed") VALUES (?, ?, ?, ?)`
	completedInt := 0
	if completed {
		completedInt = 1
	}
	params := []interface{}{id, listId, text, completedInt}
	return query, params
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

func (s *sqliteTestSuite) TestInsert() {
	ctx := context.Background()
	query, params := createInsert("1", "2", "Sample text", true)
	_, err := s.conns[0].ExecContext(ctx, query, params...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	query = `SELECT * FROM todo`
	rows, err := s.conns[0].QueryContext(ctx, query)
	s.NoError(err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var listId string
		var text string
		var completed int
		err = rows.Scan(&id, &listId, &completed, &text)
		s.NoError(err)

		s.Equal("1", id)
		s.Equal("2", listId)
		s.Equal("Sample text", text)
		s.Equal(1, completed)
	}
}

func (s *sqliteTestSuite) TestMerge() {
	ctx := context.Background()
	todos := make([]Todo, gofakeit.IntRange(4, 20))
	for i := range todos {
		todos[i] = Todo{
			ID:        gofakeit.IntRange(1, 1000),
			Text:      gofakeit.Sentence(3),
			Completed: gofakeit.Bool(),
		}
	}

	for i, todo := range todos {
		query, params := createInsert(i, todo.ID, todo.Text, todo.Completed)
		_, err := s.conns[0].ExecContext(ctx, query, params...)
		if err != nil {
			fmt.Println("Error executing query:", err)
			return
		}
	}
	sync(ctx, s.conns[0], s.conns[1])
	sync(ctx, s.conns[0], s.conns[2])
	s.True(assertAllRowsByConn(ctx, s.conns), "All rows should be equal, DBs sync failed")
	sync(ctx, s.conns[1], s.conns[0])
	sync(ctx, s.conns[2], s.conns[0])
	assertAllRowsByConn(ctx, s.conns)
	s.True(assertAllRowsByConn(ctx, s.conns), "All rows should be equal, DBs sync failed")
}

func assertAllRowsByConn(ctx context.Context, conns []*sql.Conn) bool {
	results := make([][]map[string]interface{}, len(conns))

	for i, db := range conns {
		query := `SELECT * FROM todo ORDER BY id DESC`
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
