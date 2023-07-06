package asset

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

type IAssetSvc interface {
	All(ctx context.Context) ([]*ent.Asset, error)
	ListByUser(ctx context.Context, userID string) ([]*ent.Asset, error)
	Create(ctx context.Context, userID string, asset *ent.Asset) (*ent.Asset, error)
	Update(ctx context.Context, userID string, asset *ent.Asset) error
	Delete(ctx context.Context, userID string, assetID string) error
}

type AssetSvc struct {
	config *config.Config
	repo   dal.INwRepo
}

func NewAssetSvc(cfg *config.Config, repo dal.INwRepo) *AssetSvc {
	return &AssetSvc{config: cfg, repo: repo}
}

func (svc AssetSvc) All(ctx context.Context) ([]*ent.Asset, error) {
	return svc.repo.GetAssetRepo().All(ctx)
}

func (svc AssetSvc) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.Asset, error) {
	return svc.repo.GetAssetRepo().ListByUser(ctx, userID)
}

func (svc AssetSvc) Create(ctx context.Context, userID uuid.UUID, asset *ent.Asset) (*ent.Asset, error) {
	return svc.repo.GetAssetRepo().Create(ctx, userID, asset)
}

func (svc AssetSvc) Update(ctx context.Context, userID uuid.UUID, asset *ent.Asset) error {
	return svc.repo.GetAssetRepo().Update(ctx, userID, asset)
}
