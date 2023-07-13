package provider

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListConnection(t *testing.T) {
	p := NewProviderSvc()
	uid := uuid.FromStringOrNil("test_id")
	p.NewSqliteConnect(uid.String())
	ctx := context.Background()
	conns, err := p.AllConnection(ctx, uid)
	assert.Nil(t, err)
	assert.Equal(t, len(conns), 0)
}
