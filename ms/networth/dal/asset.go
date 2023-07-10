package dal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/assettable"
)

type IAssetRepo interface {
	ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.AssetTable, error)
	Create(ctx context.Context, userID uuid.UUID, asset *ent.AssetTable) (*ent.AssetTable, error)
	Update(ctx context.Context, userID uuid.UUID, uAsset *ent.AssetTable) error
	Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error
}

type assetRepo struct {
	pgsql *ent.Client
}

func NewAssetRepo(pgsql *ent.Client) *assetRepo {
	return &assetRepo{pgsql: pgsql}
}

func (r *assetRepo) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.AssetTable, error) {
	s, err := r.pgsql.AssetTable.
		Query().
		Where(assettable.UserID(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *assetRepo) Create(ctx context.Context, userID uuid.UUID, assettable *ent.AssetTable) (*ent.AssetTable, error) {
	newAsset, err := r.pgsql.AssetTable.
		Create().
		SetUserID(userID).
		SetSheet(assettable.Sheet).
		SetSection(assettable.Section).
		SetDescription(assettable.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return newAsset, nil
}

func (r *assetRepo) Update(ctx context.Context, userID uuid.UUID, uAsset *ent.AssetTable) error {
	_, err := r.pgsql.AssetTable.
		Update().
		Where(
			assettable.And(
				assettable.ID(uAsset.ID),
				assettable.UserID(userID),
			),
		).
		SetSheet(uAsset.Sheet).
		SetSection(uAsset.Section).
		SetDescription(uAsset.Description).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *assetRepo) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	_, err := r.pgsql.AssetTable.
		Delete().
		Where(
			assettable.And(
				assettable.ID(assetID),
				assettable.UserID(userID),
			),
		).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
