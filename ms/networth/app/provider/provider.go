package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/income"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/manualitem"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/transaction"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
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

func (p *ProviderSvc) getSqliteConnect(userId string) *ent.Client {
	storage := p.userStorage[userId]
	if storage == nil {
		storage = p.NewSqliteConnect(userId)
	}
	return storage
}

func (p *ProviderSvc) AllManualItem(ctx context.Context, userId uuid.UUID) ([]*ent.ManualItem, error) {
	storage := p.getSqliteConnect(userId.String())
	ma, err := storage.ManualItem.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return ma, nil
}

func (p *ProviderSvc) AllConnection(ctx context.Context, userId uuid.UUID) ([]*ent.Connection, error) {
	storage := p.getSqliteConnect(userId.String())
	conns, err := storage.Connection.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return conns, nil
}

func (p *ProviderSvc) CreateManualItem(ctx context.Context, userId uuid.UUID, mi *dto.ManualItemBody) error {
	valid := ValidateProvider(mi.ProviderName)
	if !valid {
		return errorhandler.NewHTTPError(http.StatusBadRequest, "invalid provider name")
	}

	storage := p.getSqliteConnect(userId.String())
	_, err := storage.ManualItem.Create().
		SetProviderName(mi.ProviderName).
		SetCategory(mi.Category).
		SetItemTableID(mi.ItemTableID).
		SetType(mi.Type).
		SetDescription(mi.Description).
		SetValue(mi.Value).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProviderSvc) UpdateManualItem(ctx context.Context, userId uuid.UUID, ma *dto.ManualItemBody) error {
	valid := ValidateProvider(ma.ProviderName)
	if !valid {
		return errorhandler.NewHTTPError(http.StatusBadRequest, "invalid provider name")
	}

	storage := p.getSqliteConnect(userId.String())
	_, err := storage.ManualItem.Update().
		SetProviderName(ma.ProviderName).
		SetCategory(ma.Category).
		SetItemTableID(ma.ItemTableID).
		SetType(ma.Type).
		SetDescription(ma.Description).
		SetValue(ma.Value).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProviderSvc) DeleteManualItem(ctx context.Context, userId uuid.UUID, itemId uuid.UUID) error {
	storage := p.getSqliteConnect(userId.String())
	_, err := storage.ManualItem.Delete().Where(manualitem.ID(itemId)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProviderSvc) ConnectionByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Connection, error) {
	storage := p.getSqliteConnect(userId.String())
	conn, err := storage.Connection.Query().Where(connection.ProviderName(providerName)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return conn, nil
}

func (p *ProviderSvc) AccountByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Account, error) {
	storage := p.getSqliteConnect(userId.String())
	a, err := storage.Account.Query().Where(account.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *ProviderSvc) TransactionByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Transaction, error) {
	storage := p.getSqliteConnect(userId.String())
	a, err := storage.Transaction.Query().Where(transaction.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *ProviderSvc) IncomeByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Income, error) {
	storage := p.getSqliteConnect(userId.String())
	i, err := storage.Income.Query().Where(income.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (p *ProviderSvc) SaveConnection(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	storage := p.getSqliteConnect(userId.String())

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

func (p *ProviderSvc) SaveAccount(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	storage := p.getSqliteConnect(userId.String())

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

func (p *ProviderSvc) SaveTransaction(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	storage := p.getSqliteConnect(userId.String())

	json := toJSON(data)
	_, err := storage.Transaction.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) SaveIncome(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	storage := p.getSqliteConnect(userId.String())

	json := toJSON(data)
	_, err := storage.Income.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) CheckAccountExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	storage := p.getSqliteConnect(userId.String())
	exist, err := storage.Account.Query().Where(account.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}
	return exist, nil
}

func (p *ProviderSvc) CheckTransactionExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	storage := p.getSqliteConnect(userId.String())

	exist, err := storage.Transaction.Query().Where(transaction.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	return exist, nil
}

func (p *ProviderSvc) CheckIncomeExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	storage := p.getSqliteConnect(userId.String())

	exist, err := storage.Income.Query().Where(income.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	return exist, nil
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
