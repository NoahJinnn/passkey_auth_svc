package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProviders(t *testing.T) {
	p := NewProviderSvc("file:user1.db?cache=shared&_fk=1")
	ctx := context.Background()
	providers := p.List(ctx)
	assert.Equal(t, len(providers), 0)
}
