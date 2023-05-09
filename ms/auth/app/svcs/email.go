package svcs

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type IEmailSvc interface {
	ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error)
	Delete(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error
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

func (svc emailSvc) Delete(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed to fetch user from db: %w", err)
	}

	var emailToBeDeleted *ent.Email
	for _, email := range user.Edges.Emails {
		if email.ID == emailId {
			emailToBeDeleted = email
			break
		}
	}
	if emailToBeDeleted == nil {
		return errors.New("email with given emailId not available")
	}

	return svc.repo.GetEmailRepo().Delete(ctx, emailToBeDeleted)
}
