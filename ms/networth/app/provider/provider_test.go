package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInstitution(t *testing.T) {
	p := NewProviderSvc()
	p.NewConnect("test_id")
	ctx := context.Background()
	instis, err := p.ListInstitution(ctx, "test_id")
	assert.Nil(t, err)
	assert.Equal(t, len(instis), 0)
}

func TestListConnection(t *testing.T) {
	p := NewProviderSvc()
	p.NewConnect("test_id")
	ctx := context.Background()
	conns, err := p.ListConnection(ctx, "test_id")
	assert.Nil(t, err)
	assert.Equal(t, len(conns), 0)
}
