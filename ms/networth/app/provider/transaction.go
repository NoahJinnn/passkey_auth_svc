package provider

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/transaction"
)

func (p *ProviderSvc) TransactionByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Transaction, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	a, err := entClient.Transaction.Query().Where(transaction.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *ProviderSvc) SaveTransaction(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	json := toJSON(data)
	_, err := entClient.Transaction.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProviderSvc) CheckTransactionExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	exist, err := entClient.Transaction.Query().Where(transaction.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	return exist, nil
}
