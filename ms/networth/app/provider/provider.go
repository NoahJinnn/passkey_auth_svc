package provider

import (
	"context"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type ProviderSvc struct {
	userConn map[string]*ent.Client
}

func NewProviderSvc() *ProviderSvc {
	return &ProviderSvc{
		userConn: nil,
	}
}

func (p *ProviderSvc) NewConnect(userId string) {
	dns := sqliteDns(userId)
	if p.userConn == nil {
		p.userConn = make(map[string]*ent.Client)
		p.userConn[dns] = sqlite.NewSqliteClient(dns)
	} else {
		p.userConn[dns] = sqlite.NewSqliteClient(dns)
	}
}

func (p *ProviderSvc) ListInstitution(ctx context.Context, userId string) []*ent.Institution {
	dns := sqliteDns(userId)
	conn := p.userConn[dns]
	return conn.Institution.Query().AllX(ctx)
}

func sqliteDns(userId string) string {
	return "file:" + userId + ".db?cache=shared&_fk=1"
}
