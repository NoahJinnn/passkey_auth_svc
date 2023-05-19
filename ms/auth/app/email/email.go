package email

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type Ctx = context.Context

type IEmailSvc interface {
	ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error)
	Delete(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error
	SetPrimaryEmail(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error
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

// TODO: This is for multiple emails per user feature
func (svc emailSvc) SetPrimaryEmail(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed to fetch user from db: %w", err)
	}

	emails := user.Edges.Emails
	for _, email := range emails {
		if email.ID == emailId {
			if email.ID == user.Edges.PrimaryEmail.EmailID {
				return fmt.Errorf("email is already primary: %w", err)
			}
		}
	}

	return svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		primaryEmail, err := svc.repo.GetEmailRepo().GetPrimary(ctx, emailId)
		if err != nil {
			return fmt.Errorf("failed to fetch primary email from db: %w", err)
		}

		if primaryEmail == nil {
			_, err = client.PrimaryEmail.Create().
				SetUserID(user.ID).
				SetEmailID(emailId).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to store primary email: %w", err)
			}
		} else {
			primaryEmail.EmailID = emailId
			err = svc.repo.GetEmailRepo().UpdatePrimary(ctx, *primaryEmail)
			if err != nil {
				return fmt.Errorf("failed to change primary email: %w", err)
			}
		}

		return nil
	})

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
