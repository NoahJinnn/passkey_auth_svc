package email

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
)

type Ctx = context.Context

type IEmailSvc interface {
	ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error)
	Create(ctx Ctx, userId uuid.UUID, email string) (*ent.Email, error)
	Delete(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error
	SetPrimaryEmail(ctx Ctx, userId uuid.UUID, emailId uuid.UUID) error
}

type emailSvc struct {
	cfg  *config.Config
	repo dal.IAuthRepo
}

func NewEmailSvc(cfg *config.Config, repo dal.IAuthRepo) IEmailSvc {
	return emailSvc{cfg: cfg, repo: repo}
}

func (svc emailSvc) ListByUser(ctx Ctx, userId uuid.UUID) ([]*ent.Email, error) {
	return svc.repo.GetEmailRepo().ListByUser(ctx, userId)
}

func (svc emailSvc) Create(ctx Ctx, userId uuid.UUID, address string) (*ent.Email, error) {

	emailCount, err := svc.repo.GetEmailRepo().CountByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to count user emails: %w", err)
	}

	if emailCount >= svc.cfg.MaxEmailAddresses {
		return nil, errorhandler.NewHTTPError(http.StatusConflict).SetInternal(errors.New("max number of email addresses reached"))
	}

	var newMail *ent.Email
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		// Query email by email address
		email, err := client.Email.Query().Where(email.Address(address)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return fmt.Errorf("failed querying email by address: %w", err)
		}
		if email != nil {
			if !email.UserID.IsNil() {
				// The email already exists and is assigned already.
				return errorhandler.NewHTTPError(http.StatusConflict).SetInternal(fmt.Errorf("user with email %s already exists", address))
			}
		} else {
			email, err = client.Email.Create().
				SetAddress(address).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed creating email: %w", err)
			}
		}

		newMail = email
		return nil
	}); err != nil {
		return nil, err
	}
	return newMail, nil

}

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
