package test

import (
	"context"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

func NewRepo(
	db *ent.Client,
) dal.INwRepo {
	return &repoT{
		Db: db,
	}
}

type repoT struct {
	Db *ent.Client
}

func (r repoT) WithTx(ctx context.Context, exec func(ctx context.Context, client *ent.Client) error) error {
	txForw := func(db *ent.Client) error {
		return exec(ctx, db)
	}
	return txForw(nil)
}
