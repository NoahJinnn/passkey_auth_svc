package ws

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/networth/dal"
)

type IChangesetSvc interface {
	Latest(ctx context.Context, userId uuid.UUID) (*ent.Changeset, error)
	Create(ctx context.Context, userId uuid.UUID, changeset *ent.Changeset) (*ent.Changeset, error)
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

func (s *ChangesetSvc) Create(ctx context.Context, userId uuid.UUID, changeset *ent.Changeset) (*ent.Changeset, error) {
	return s.repo.GetChangesetRepo().Create(ctx, userId, changeset)
}
