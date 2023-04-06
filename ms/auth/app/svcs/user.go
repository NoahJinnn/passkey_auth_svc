package svcs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
)

type IUserSvc interface {
	GetById(ctx Ctx, userID uuid.UUID) (*ent.User, *string, error)
	Create(ctx Ctx, email string) (newU *ent.User, emailID uuid.UUID, err error)
}

type userSvc struct {
	repo dal.IRepo
	cfg  *config.Config
}

func NewUserSvc(cfg *config.Config, repo dal.IRepo) IUserSvc {
	return &userSvc{
		repo: repo,
		cfg:  cfg,
	}
}

func (svc *userSvc) Create(ctx Ctx, address string) (newU *ent.User, emailID uuid.UUID, err error) {
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		newU, err = client.User.Create().Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating user: %w", err)
		}

		// Query email by email address
		email, err := client.Email.Query().Where(email.Address(address)).Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("failed querying email by address: %w", err)
			}
		}

		if email != nil {
			if !email.UserID.IsNil() {
				// The email already exists and is assigned already.
				return dto.NewHTTPError(http.StatusConflict).SetInternal(fmt.Errorf("user with email %s already exists", address))
			}

			// TODO: Implement email verification flow
			// if !svc.cfg.Emails.RequireVerification {
			// 	Assign the email address to the user because it's currently unassigned and email verification is turned off.
			// 	email.UserID = newU.ID
			// 	err = svc.repo.GetEmailRepo.Update(*email)
			// 	if err != nil {
			// 		return fmt.Errorf("failed to update email address: %w", err)
			// 	}
			// }
		} else {
			if svc.cfg.Emails.RequireVerification {
				// The email can only be assigned to the user via passcode verification.
				email, err = client.Email.Create().
					SetAddress(address).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed creating email: %w", err)
				}
			} else {
				email, err = client.Email.Create().
					SetAddress(address).
					SetUserID(newU.ID).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed creating email: %w", err)
				}
			}
		}

		if !svc.cfg.Emails.RequireVerification {
			_, err = client.PrimaryEmail.Create().
				SetUserID(newU.ID).
				SetEmailID(email.ID).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to store primary email: %w", err)
			}
		}
		emailID = email.ID
		return nil
	}); err != nil {
		return newU, emailID, err
	}
	return newU, emailID, nil
}

func (svc *userSvc) GetById(ctx Ctx, userID uuid.UUID) (*ent.User, *string, error) {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, nil, dto.NewHTTPError(http.StatusNotFound).SetInternal(errors.New("user not found"))
	}

	var emailAddress *string
	if e := user.Edges.PrimaryEmail; e != nil {
		model, err := e.QueryEmail().Only(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get email address of primary email: %w", err)
		}
		emailAddress = &model.Address
	}
	return user, emailAddress, nil
}