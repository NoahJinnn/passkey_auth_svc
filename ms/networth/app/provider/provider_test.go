package provider

import (
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
			uids: []string{"test_id1", "test_id2", "test_id3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, uid := range tt.uids {
				p.NewSqliteConnect(uid)
			}
			assert.Equal(t, len(tt.uids), len(p.userStorage))

			for _, uid := range tt.uids {
				p.ClearSqliteConnect(uid)
			}
			assert.Equal(t, 0, len(p.userStorage))
		})
	}
}
