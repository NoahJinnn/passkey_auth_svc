package dal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/changeset"
)

type IChangesetRepo interface {
	Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error)
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
