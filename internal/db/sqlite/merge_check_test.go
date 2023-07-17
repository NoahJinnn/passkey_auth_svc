package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type sqliteTestSuite struct {
	suite.Suite
	db []*sql.DB
}

func (suite *sqliteTestSuite) SetupTest() {
	for i := 0; i < 2; i++ {
		db := NewSqliteDrive("file:" + fmt.Sprintf("test %d", i) + "file:ent?mode=memory&_fk=1")
		db.Exec(`CREATE TABLE todo ("id" primary key, "listId", "completed", "text")`)
		db.Exec(`SELECT crsql_as_crr('todo')`)

		suite.db = append(suite.db, db)
	}
}

func TestSqliteTestSuite(t *testing.T) {
	// t.Parallel()
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

func all(db *sql.DB, q []interface{}) []map[string]interface{} {
	query := q[0].(string)
	params := q[1].([]interface{})
	stmt, err := db.Prepare(query)
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

func sync(left *sql.DB, right *sql.DB) {
	changesets, err := left.Query(`SELECT * FROM crsql_changes`)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := right.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO crsql_changes VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for changesets.Next() {
		var table string
		var pk string
		var cid string
		var val interface{}
		var col_version int
		var db_version int
		var site_id string
		err = changesets.Scan(&table, &pk, &cid, &val, &col_version, &db_version, &site_id)
		log.Fatal(err)
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

func assertAll(dbs []*sql.DB) {
	results := make([][]map[string]interface{}, len(dbs))

	for i, db := range dbs {
		query := fmt.Sprintf(`SELECT * FROM todo ORDER BY id DESC`)
		results[i] = all(db, []interface{}{query, []interface{}{}})
	}

	for i := 0; i < len(results)-1; i++ {
		if !isEqual(results[i], results[i+1]) {
			fmt.Println("Assertion failed!")
			fmt.Println("Results:")
			for _, rows := range results {
				fmt.Println(rows)
			}
			return
		}
		fmt.Println("Results:")
		for _, rows := range results {
			fmt.Println(rows)
		}
	}

	fmt.Println("All assertions passed!")
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

func (s *sqliteTestSuite) TestInsert() {
	query, params := createInsert("1", "2", "Sample text", true)
	_, err := s.db[0].Exec(query, params...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	query = `SELECT * FROM todo`
	rows, err := s.db[0].Query(query)
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

type Todo struct {
	ID        int
	Text      string
	Completed bool
}

func (s *sqliteTestSuite) TestMerge() {
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
		for _, db := range s.db {
			_, err := db.Exec(query, params...)
			if err != nil {
				fmt.Println("Error executing query:", err)
				return
			}
		}
	}
	sync(s.db[0], s.db[1])
	assertAll(s.db)
}
