package asset_table

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
)

type IAssetTableSvc interface {
	ListByUser(ctx context.Context, userID string) ([]*ent.AssetTable, error)
	Create(ctx context.Context, userID string, body dto.AssetTableRequest) (*ent.AssetTable, error)
	Update(ctx context.Context, userID string, body dto.AssetTableRequest) error
	Delete(ctx context.Context, userID string, assetID string) error
}

type AssetTableSvc struct {
	config *config.Config
	repo   dal.INwRepo
}

func NewAssetTableSvc(cfg *config.Config, repo dal.INwRepo) *AssetTableSvc {
	return &AssetTableSvc{config: cfg, repo: repo}
}

func (svc *AssetTableSvc) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.AssetTable, error) {
	return svc.repo.GetAssetTableRepo().ListByUser(ctx, userID)
}

func (svc *AssetTableSvc) Create(ctx context.Context, userID uuid.UUID, body dto.AssetTableRequest) (*ent.AssetTable, error) {
	asset := &ent.AssetTable{
		Sheet:       body.Sheet,
		Section:     body.Section,
		Description: body.Description,
	}

	return svc.repo.GetAssetTableRepo().Create(ctx, userID, asset)
}

func (svc *AssetTableSvc) Update(ctx context.Context, userID uuid.UUID, body dto.AssetTableRequest) error {
	asset := &ent.AssetTable{
		Sheet:       body.Sheet,
		Section:     body.Section,
		Description: body.Description,
	}

	return svc.repo.GetAssetTableRepo().Update(ctx, userID, asset)
}

func (svc *AssetTableSvc) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	return svc.repo.GetAssetTableRepo().Delete(ctx, userID, assetID)
}
