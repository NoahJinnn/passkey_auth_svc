package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProviders(t *testing.T) {
	p := NewProviderSvc()
	p.NewConnect("user1")
	ctx := context.Background()
	instis := p.ListInstitution(ctx, "user1")
	assert.Equal(t, len(instis), 0)
}
