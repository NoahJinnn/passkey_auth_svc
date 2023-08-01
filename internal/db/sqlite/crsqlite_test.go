package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

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
		db := NewSqliteDrive("file:" + fmt.Sprintf("test_%d", i) + ".db?cache=shared&_fk=1")
		client := NewSqliteEnt(ctx, db)
		s.clients = append(s.clients, client)
		conn := NewSqliteConn(ctx, db)
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
