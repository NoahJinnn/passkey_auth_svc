package dal

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/changeset"
)

type IChangesetRepo interface {
	Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error)
	Create(ctx context.Context, userId uuid.UUID, changeset *ent.Changeset) (*ent.Changeset, error)
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

func (r *changesetRepo) Create(ctx Ctx, userId uuid.UUID, s *ent.Changeset) (*ent.Changeset, error) {
	news, err := r.pgsql.Changeset.
		Create().
		SetUserID(userId).
		SetCsList(s.CsList).
		SetDbVersion(s.DbVersion).
		SetSiteID(s.SiteID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return news, nil
}
