package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListConnection(t *testing.T) {
	p := NewProviderSvc()
	p.NewSqliteConnect("test_id")
	ctx := context.Background()
	conns, err := p.ListConnection(ctx, "test_id")
	assert.Nil(t, err)
	assert.Equal(t, len(conns), 0)
}
