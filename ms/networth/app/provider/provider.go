package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
)

type ProviderSvc struct {
	userStorage map[string]*ent.Client
}

func NewProviderSvc() *ProviderSvc {
	return &ProviderSvc{
		userStorage: nil,
	}
}

func (p *ProviderSvc) NewSqliteConnect(userId string) *ent.Client {
	dns := sqliteDns(userId)
	if p.userStorage == nil {
		p.userStorage = make(map[string]*ent.Client)
	}
	if p.userStorage[userId] == nil {
		p.userStorage[userId] = sqlite.NewSqliteClient(dns)
	}
	return p.userStorage[userId]
}

func (p *ProviderSvc) GetSqliteConnect(userId string) *ent.Client {
	storage := p.userStorage[userId]
	if storage == nil {
		storage = p.NewSqliteConnect(userId)
	}
	return storage
}

func (p *ProviderSvc) ListConnection(ctx context.Context, userId string) ([]*ent.Connection, error) {
	storage := p.GetSqliteConnect(userId)
	conns, err := storage.Connection.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return conns, nil
}

func (p *ProviderSvc) ConnectionByProviderName(ctx context.Context, userId string, providerName string) (*ent.Connection, error) {
	storage := p.GetSqliteConnect(userId)
	conn, err := storage.Connection.Query().Where(connection.ProviderName(providerName)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return conn, nil
}

func (p *ProviderSvc) SaveConnection(ctx context.Context, providerName string, userId string, data interface{}) error {
	storage := p.GetSqliteConnect(userId)

	json := toJSON(data)
	_, err := storage.Connection.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) SaveAccount(ctx context.Context, providerName string, userId string, data interface{}) error {
	storage := p.GetSqliteConnect(userId)

	json := toJSON(data)
	_, err := storage.Account.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) AccountByProviderName(ctx context.Context, userId string, providerName string) (*ent.Account, error) {
	storage := p.GetSqliteConnect(userId)
	a, err := storage.Account.Query().Where(account.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func sqliteDns(userId string) string {
	if userId == "test_id" {
		return "file:" + userId + "file:ent?mode=memory&_fk=1"

	}
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
