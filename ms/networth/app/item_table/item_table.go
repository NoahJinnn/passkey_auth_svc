package item_table

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
)

type IAssetTableSvc interface {
	ListByUser(ctx context.Context, userID string) ([]*ent.ItemTable, error)
	Create(ctx context.Context, userID string, body dto.ItemTableBody) (*ent.ItemTable, error)
	Update(ctx context.Context, userID string, body dto.ItemTableBody) error
	Delete(ctx context.Context, userID string, assetID string) error
}

type ItemTableSvc struct {
	config *config.Config
	repo   dal.INwRepo
}

func NewItemTableSvc(cfg *config.Config, repo dal.INwRepo) *ItemTableSvc {
	return &ItemTableSvc{config: cfg, repo: repo}
}

func (svc *ItemTableSvc) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.ItemTable, error) {
	return svc.repo.GetItemTableRepo().ListByUser(ctx, userID)
}

func (svc *ItemTableSvc) Create(ctx context.Context, userID uuid.UUID, body dto.ItemTableBody) (*ent.ItemTable, error) {
	asset := &ent.ItemTable{
		Category:    body.Category,
		Sheet:       body.Sheet,
		Section:     body.Section,
		Description: body.Description,
	}

	return svc.repo.GetItemTableRepo().Create(ctx, userID, asset)
}

func (svc *ItemTableSvc) Update(ctx context.Context, userID uuid.UUID, body dto.ItemTableBody) error {
	asset := &ent.ItemTable{
		Sheet:       body.Sheet,
		Section:     body.Section,
		Description: body.Description,
	}

	return svc.repo.GetItemTableRepo().Update(ctx, userID, asset)
}

func (svc *ItemTableSvc) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	return svc.repo.GetItemTableRepo().Delete(ctx, userID, assetID)
}
