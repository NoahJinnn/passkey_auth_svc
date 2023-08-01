package provider

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type ProviderSvc struct {
	userStorage map[string]*SqliteInstance
}

type SqliteInstance struct {
	entClient *ent.Client
	conn      *sql.Conn
}

func NewProviderSvc() *ProviderSvc {
	return &ProviderSvc{
		userStorage: nil,
	}
}

func (p *ProviderSvc) NewSqliteConnect(ctx context.Context, userId string) *SqliteInstance {
	if p.userStorage == nil {
		p.userStorage = make(map[string]*SqliteInstance)
	}
	if p.userStorage[userId] == nil {
		dns := sqliteDns(userId)
		db := sqlite.NewSqliteDrive(dns)
		entClient := sqlite.NewSqliteEnt(ctx, db)
		p.userStorage[userId] = &SqliteInstance{
			entClient: entClient,
			conn:      sqlite.NewSqliteConn(ctx, db),
		}
	}
	return p.userStorage[userId]
}

func (p *ProviderSvc) getSqliteConnect(ctx context.Context, userId string) *SqliteInstance {
	storage := p.userStorage[userId]
	if storage == nil {
		storage = p.NewSqliteConnect(ctx, userId)
	}
	return storage
}

func (p *ProviderSvc) ClearSqliteDB(userId string) {
	sqlite := p.userStorage[userId]
	sqlite.entClient.Close()
	if p.userStorage != nil {
		delete(p.userStorage, userId)
	}
}

func sqliteDns(userId string) string {
	return "file:" + userId + ".db?cache=shared&_fk=1"
}

func toJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return string(jsonData)
}
