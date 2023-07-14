package dal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/itemtable"
)

type IItemTableRepo interface {
	ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.ItemTable, error)
	Create(ctx context.Context, userID uuid.UUID, asset *ent.ItemTable) (*ent.ItemTable, error)
	Update(ctx context.Context, userID uuid.UUID, uAsset *ent.ItemTable) error
	Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error
}

type itemTableRepo struct {
	pgsql *ent.Client
}

func NewItemTableRepo(pgsql *ent.Client) *itemTableRepo {
	return &itemTableRepo{pgsql: pgsql}
}

func (r *itemTableRepo) ListByUser(ctx context.Context, userID uuid.UUID) ([]*ent.ItemTable, error) {
	s, err := r.pgsql.ItemTable.
		Query().
		Where(itemtable.UserID(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *itemTableRepo) Create(ctx context.Context, userID uuid.UUID, itemtable *ent.ItemTable) (*ent.ItemTable, error) {
	newAsset, err := r.pgsql.ItemTable.
		Create().
		SetUserID(userID).
		SetSheet(itemtable.Sheet).
		SetSection(itemtable.Section).
		SetDescription(itemtable.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return newAsset, nil
}

func (r *itemTableRepo) Update(ctx context.Context, userID uuid.UUID, uAsset *ent.ItemTable) error {
	_, err := r.pgsql.ItemTable.
		Update().
		Where(
			itemtable.And(
				itemtable.ID(uAsset.ID),
				itemtable.UserID(userID),
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

func (r *itemTableRepo) Delete(ctx context.Context, userID uuid.UUID, assetID uuid.UUID) error {
	_, err := r.pgsql.ItemTable.
		Delete().
		Where(
			itemtable.And(
				itemtable.ID(assetID),
				itemtable.UserID(userID),
			),
		).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
