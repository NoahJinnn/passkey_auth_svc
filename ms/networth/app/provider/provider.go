package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
)

type ProviderSvc struct {
	userStorage map[string]*ent.Client
}

func NewProviderSvc() *ProviderSvc {
	return &ProviderSvc{
		userStorage: nil,
	}
}

func (p *ProviderSvc) NewConnect(userId string) {
	dns := sqliteDns(userId)
	if p.userStorage == nil {
		p.userStorage = make(map[string]*ent.Client)
		p.userStorage[userId] = sqlite.NewSqliteClient(dns)
	} else {
		if p.userStorage[userId] != nil {
			return
		}
		p.userStorage[userId] = sqlite.NewSqliteClient(dns)
	}
}

func (p *ProviderSvc) ListInstitution(ctx context.Context, userId string) ([]*ent.Institution, error) {
	storage := p.userStorage[userId]
	instis, err := storage.Institution.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	defer storage.Close()
	return instis, nil
}

func (p *ProviderSvc) ListConnection(ctx context.Context, userId string) ([]*ent.Connection, error) {
	storage := p.userStorage[userId]
	conns, err := storage.Connection.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return conns, nil
}

func (p *ProviderSvc) SaveConnection(ctx context.Context, userId string, env string, data interface{}) error {
	storage := p.userStorage[userId]
	json := toJSON(data)
	_, err := storage.Connection.Create().
		SetData(json).
		SetEnv(env).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
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
