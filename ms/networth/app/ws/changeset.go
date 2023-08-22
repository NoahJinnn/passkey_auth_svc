package ws

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/changeset"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

type IChangesetSvc interface {
	Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error)
	Create(ctx context.Context, userId uuid.UUID, changeset *ent.Changeset) error
}

type ChangesetSvc struct {
	repo dal.INwRepo
}

func NewChangesetSvc(repo dal.INwRepo) *ChangesetSvc {
	return &ChangesetSvc{repo: repo}
}

func (s *ChangesetSvc) Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error) {
	return s.repo.GetChangesetRepo().Latest(ctx, userId)
}

func (s *ChangesetSvc) Create(ctx context.Context, userId uuid.UUID, newCs *ent.Changeset) error {
	if err := s.repo.WithTx(ctx, func(ctx context.Context, client *ent.Client) error {
		latest, err := client.Changeset.
			Query().
			Where(changeset.UserID(userId)).
			Only(ctx)

		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		if latest.DbVersion < newCs.DbVersion {
			_, err = client.Changeset.
				Create().
				SetUserID(userId).
				SetCsList(newCs.CsList).
				SetDbVersion(newCs.DbVersion).
				SetSiteID(newCs.SiteID).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}
