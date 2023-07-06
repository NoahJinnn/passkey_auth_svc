package dal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/asset"
)

type IAssetRepo interface {
	All(ctx context.Context) ([]*ent.Asset, error)
	ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.Asset, error)
	Create(ctx context.Context, userID uuid.UUID, asset *ent.Asset) (*ent.Asset, error)
	Update(ctx context.Context, userID uuid.UUID, uAsset *ent.Asset) error
	Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error
}

type assetRepo struct {
	pgsql *ent.Client
}

func NewAssetRepo(pgsql *ent.Client) *assetRepo {
	return &assetRepo{pgsql: pgsql}
}

func (r assetRepo) All(ctx context.Context) ([]*ent.Asset, error) {
	s, err := r.pgsql.Asset.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r assetRepo) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.Asset, error) {
	s, err := r.pgsql.Asset.
		Query().
		Where(asset.UserID(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r assetRepo) Create(ctx context.Context, userID uuid.UUID, asset *ent.Asset) (*ent.Asset, error) {
	newAsset, err := r.pgsql.Asset.
		Create().
		SetUserID(userID).
		SetSheet(asset.Sheet).
		SetSection(asset.Section).
		SetType(asset.Type).
		SetProviderName(asset.ProviderName).
		SetDescription(*asset.Description).
		SetCurrency(asset.Currency).
		SetValue(asset.Value).
		SetType(asset.Type).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return newAsset, nil
}

func (r assetRepo) Update(ctx context.Context, userID uuid.UUID, uAsset *ent.Asset) error {
	_, err := r.pgsql.Asset.
		Update().
		Where(
			asset.And(
				asset.ID(uAsset.ID),
				asset.UserID(userID),
			),
		).
		SetSheet(uAsset.Sheet).
		SetSection(uAsset.Section).
		SetType(uAsset.Type).
		SetProviderName(uAsset.ProviderName).
		SetDescription(*uAsset.Description).
		SetCurrency(uAsset.Currency).
		SetValue(uAsset.Value).
		SetType(uAsset.Type).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r assetRepo) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	_, err := r.pgsql.Asset.
		Delete().
		Where(
			asset.And(
				asset.ID(assetID),
				asset.UserID(userID),
			),
		).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
