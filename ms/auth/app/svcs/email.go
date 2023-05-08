package svcs

import (
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type IEmailSvc interface {
	ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error)
}

type emailSvc struct {
	repo dal.IAuthRepo
}

func NewEmailSvc(repo dal.IAuthRepo) IEmailSvc {
	return emailSvc{repo: repo}
}

func (svc emailSvc) ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error) {
	return svc.repo.GetEmailRepo().ListByUser(ctx, userId)
}
