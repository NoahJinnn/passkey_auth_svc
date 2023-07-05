package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInstitution(t *testing.T) {
	p := NewProviderSvc()
	p.NewConnect("user1")
	ctx := context.Background()
	instis, err := p.ListInstitution(ctx, "user1")
	assert.Nil(t, err)
	assert.Equal(t, len(instis), 0)
}

func TestListConnection(t *testing.T) {
	p := NewProviderSvc()
	p.NewConnect("user1")
	ctx := context.Background()
	conns, err := p.ListConnection(ctx, "user1")
	assert.Nil(t, err)
	assert.Equal(t, len(conns), 0)
}
