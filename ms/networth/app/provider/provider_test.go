package provider

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteConnection(t *testing.T) {
	p := NewProviderSvc()

	tests := []struct {
		name string
		uids []string
	}{
		{
			name: "3 conn success",
			uids: []string{"id1", "id2", "id3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, uid := range tt.uids {
				conn := p.NewSqliteConnect(context.Background(), uid).conn
				r := conn.QueryRowContext(context.Background(), "SELECT quote(crsql_siteid());")
				var siteid string
				if err := r.Scan(&siteid); err == sql.ErrNoRows {
					log.Fatalf("failed to query crsql lite id: %v", err)
				}
			}
			assert.Equal(t, len(tt.uids), len(p.userStorage))

			for _, uid := range tt.uids {
				p.ClearSqliteDB(uid)
			}
			assert.Equal(t, 0, len(p.userStorage))
		})
	}
}
