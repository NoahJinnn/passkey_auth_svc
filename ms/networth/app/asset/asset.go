package asset

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
)

type IAssetSvc interface {
	ListByUser(ctx context.Context, userID string) ([]*ent.Asset, error)
	Create(ctx context.Context, userID string, body dto.AssetBodyRequest) (*ent.Asset, error)
	Update(ctx context.Context, userID string, body dto.AssetBodyRequest) error
	Delete(ctx context.Context, userID string, assetID string) error
}

type AssetSvc struct {
	config *config.Config
	repo   dal.INwRepo
}

func NewAssetSvc(cfg *config.Config, repo dal.INwRepo) *AssetSvc {
	return &AssetSvc{config: cfg, repo: repo}
}

func (svc *AssetSvc) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.Asset, error) {
	return svc.repo.GetAssetRepo().ListByUser(ctx, userID)
}

func (svc *AssetSvc) Create(ctx context.Context, userID uuid.UUID, body dto.AssetBodyRequest) (*ent.Asset, error) {
	asset := &ent.Asset{
		Sheet:        body.Sheet,
		Section:      body.Section,
		Type:         body.Type,
		ProviderName: body.ProviderName,
		Currency:     body.Currency,
		Value:        body.Value,
		Description:  body.Description,
	}

	return svc.repo.GetAssetRepo().Create(ctx, userID, asset)
}

func (svc *AssetSvc) Update(ctx context.Context, userID uuid.UUID, body dto.AssetBodyRequest) error {
	asset := &ent.Asset{
		Sheet:        body.Sheet,
		Section:      body.Section,
		Type:         body.Type,
		ProviderName: body.ProviderName,
		Currency:     body.Currency,
		Value:        body.Value,
		Description:  body.Description,
	}

	return svc.repo.GetAssetRepo().Update(ctx, userID, asset)
}

func (svc *AssetSvc) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	return svc.repo.GetAssetRepo().Delete(ctx, userID, assetID)
}
