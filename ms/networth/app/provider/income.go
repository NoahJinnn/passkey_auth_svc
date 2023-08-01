package provider

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/income"
)

func (p *ProviderSvc) CheckIncomeExist(ctx context.Context, userId uuid.UUID, providerName string) (bool, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	exist, err := entClient.Income.Query().Where(income.ProviderName(providerName)).Exist(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}

	return exist, nil
}
func (p *ProviderSvc) IncomeByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Income, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	i, err := entClient.Income.Query().Where(income.ProviderName(providerName)).First(ctx)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (p *ProviderSvc) SaveIncome(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	json := toJSON(data)
	_, err := entClient.Income.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
