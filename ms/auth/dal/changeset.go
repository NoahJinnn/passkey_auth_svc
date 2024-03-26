package dal

import (
	"context"
	"fmt"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ent/changeset"
	"github.com/gofrs/uuid"
)

type IChangesetRepo interface {
	Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error)
	Delete(ctx context.Context, userId uuid.UUID) error
}

type changesetRepo struct {
	pgsql *ent.Client
}

func NewChangesetRepo(pgsql *ent.Client) *changesetRepo {
	return &changesetRepo{pgsql: pgsql}
}

func (r *changesetRepo) Latest(ctx Ctx, userId uuid.UUID) (*ent.Changeset, error) {
	cs, err := r.pgsql.Changeset.
		Query().
		Where(changeset.UserID(userId)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return cs, nil
}

func (r *changesetRepo) Delete(ctx Ctx, userId uuid.UUID) error {
	rows, err := r.pgsql.Changeset.
		Delete().
		Where(changeset.UserID(userId)).
		Exec(ctx)
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no changeset deleted")
	}
	return nil
}
