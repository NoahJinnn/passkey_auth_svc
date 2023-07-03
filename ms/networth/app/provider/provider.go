package provider

import (
	"context"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type ProviderSvc struct {
	sqliteClient *ent.Client
}

func NewProviderSvc() *ProviderSvc {
	sqliteClient := sqlite.NewSqliteClient("file:ent?mode=memory&cache=shared&_fk=1")
	return &ProviderSvc{
		sqliteClient: sqliteClient,
	}
}

func (p *ProviderSvc) Get(ctx context.Context) {
	p.sqliteClient.Provider.Query().AllX(ctx)
}
