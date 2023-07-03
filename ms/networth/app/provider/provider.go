package provider

import (
	"context"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type Provider struct {
	sqliteClient *ent.Client
}

func NewProvider(ctx context.Context) *Provider {
	sqliteClient := sqlite.NewSqliteClient(ctx, "file:ent_pgsql?mode=memory&cache=shared&_fk=1")
	return &Provider{
		sqliteClient: sqliteClient,
	}
}
