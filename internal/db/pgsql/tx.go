package pgsql

import (
	"context"
	"fmt"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
)

type Tx func(ctx context.Context, conn *ent.Client, exec func(ctx context.Context, client *ent.Client) error) error

func WithTx(ctx context.Context, conn *ent.Client, exec func(ctx context.Context, client *ent.Client) error) error {
	txForw := func(db *ent.Client) error {
		return exec(ctx, db)
	}

	tx, err := conn.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			if err != nil {
				panic(fmt.Errorf("%w: rolling back transaction: %v", err, v))
			}
			panic(v)
		}
	}()
	if err := txForw(tx.Client()); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
