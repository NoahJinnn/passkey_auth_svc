package provider

import (
	"context"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type ProviderSvc struct {
	sqliteClient *ent.Client
}

func NewProviderSvc(dsn string) *ProviderSvc {
	sqliteClient := sqlite.NewSqliteClient(dsn)
	return &ProviderSvc{
		sqliteClient: sqliteClient,
	}
}

func (p *ProviderSvc) List(ctx context.Context) []*ent.Provider {
	return p.sqliteClient.Provider.Query().AllX(ctx)
}
