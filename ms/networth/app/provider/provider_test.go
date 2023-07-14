package provider

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSqliteConnection(t *testing.T) {
	p := NewProviderSvc()

	tests := []struct {
		name string
		uid  []string
	}{
		{
			name: "1 conn success",
			uid:  []string{"test_id"},
		},
		{
			name: "3 conn success",
			uid:  []string{"id1", "id2", "id3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, uid := range tt.uid {
				p.NewSqliteConnect(uid)
				ctx := context.Background()
				conns, err := p.AllConnection(ctx, uuid.FromStringOrNil(uid))
				assert.Nil(t, err)
				assert.Equal(t, len(conns), len(tt.uid))
			}
		})
	}

}
