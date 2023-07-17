package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/stretchr/testify/suite"
)

type SqliteTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *SqliteTestSuite) SetupTest() {
	db := NewSqliteDrive("file:" + "userId" + "file:ent?mode=memory&_fk=1")
	db.Exec(`CREATE TABLE todo ("id" primary key, "listId", "completed", "text")`)
	db.Exec(`SELECT crsql_as_crr('todo')`)

	suite.db = db
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

func run(db *sql.DB, q []interface{}) {
	query := q[0].(string)
	params := q[1].([]interface{})
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	fmt.Println("Query executed successfully!")
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

func (s *SqliteTestSuite) TestInsert() {
	query, params := createInsert("1", "2", "Sample text", true)
	_, err := s.db.Exec(query, params...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	query = `SELECT * FROM todo`
	rows, err := s.db.Query(query)
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

	fmt.Println("Insertion successful!")
}
