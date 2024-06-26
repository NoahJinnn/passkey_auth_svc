package ws

import (
	"context"

	"github.com/NoahJinnn/passkey_auth_svc/ent"
	"github.com/NoahJinnn/passkey_auth_svc/ent/changeset"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/dal"
	"github.com/gofrs/uuid"
)

type ChangesetSvc struct {
	repo dal.IAuthRepo
}

func NewChangesetSvc(repo dal.IAuthRepo) *ChangesetSvc {
	return &ChangesetSvc{repo: repo}
}

func (s *ChangesetSvc) Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error) {
	return s.repo.GetChangesetRepo().Latest(ctx, userId)
}

func (s *ChangesetSvc) Upsert(ctx context.Context, userId uuid.UUID, newCs *ent.Changeset) error {
	if err := s.repo.WithTx(ctx, func(ctx context.Context, client *ent.Client) error {
		latest, err := client.Changeset.
			Query().
			Where(changeset.UserID(userId)).
			Only(ctx)

		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		if latest == nil {
			_, err = client.Changeset.
				Create().
				SetUserID(userId).
				SetDbVersion(newCs.DbVersion).
				SetSiteID(newCs.SiteID).
				SetFirstLaunch(newCs.FirstLaunch).
				Save(ctx)
		} else {
			_, err = client.Changeset.
				Update().
				Where(changeset.UserID(userId)).
				SetDbVersion(newCs.DbVersion).
				SetSiteID(newCs.SiteID).
				SetFirstLaunch(newCs.FirstLaunch).
				Save(ctx)
		}
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *ChangesetSvc) Delete(ctx context.Context, userId uuid.UUID) error {
	return s.repo.GetChangesetRepo().Delete(ctx, userId)
}

func (s *ChangesetSvc) FirstLaunch(ctx context.Context, userId uuid.UUID) (bool, error) {
	latest, err := s.repo.GetChangesetRepo().Latest(ctx, userId)
	if err != nil {
		return false, err
	}
	if latest == nil {
		return false, err
	}
	return latest.FirstLaunch, nil
}
