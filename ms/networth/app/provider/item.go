package provider

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/manualitem"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
)

func (p *ProviderSvc) AllManualItem(ctx context.Context, userId uuid.UUID) ([]*ent.ManualItem, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	ma, err := entClient.ManualItem.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return ma, nil
}

func (p *ProviderSvc) CreateManualItem(ctx context.Context, userId uuid.UUID, mi *dto.ManualItemBody) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	_, err := entClient.ManualItem.Create().
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

func (p *ProviderSvc) UpdateManualItem(ctx context.Context, userId uuid.UUID, mi *dto.ManualItemBody) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	_, err := entClient.ManualItem.Update().
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

func (p *ProviderSvc) DeleteManualItem(ctx context.Context, userId uuid.UUID, itemId uuid.UUID) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	_, err := entClient.ManualItem.Delete().Where(manualitem.ID(itemId)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
