package provider

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/account"
)

func (p *ProviderSvc) AccountByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Account, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	a, err := entClient.Account.Query().Where(account.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *ProviderSvc) SaveAccount(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	json := toJSON(data)
	_, err := entClient.Account.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) CheckAccountExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	exist, err := entClient.Account.Query().Where(account.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}
	return exist, nil
}
