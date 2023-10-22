package user

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

type UserSvc struct {
	repo dal.IAuthRepo
	cfg  *config.Config
}

func NewUserSvc(cfg *config.Config, repo dal.IAuthRepo) *UserSvc {
	return &UserSvc{
		repo: repo,
		cfg:  cfg,
	}
}

func (svc *UserSvc) Create(ctx Ctx, address string) (newU *ent.User, emailID uuid.UUID, err error) {
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		newU, err = client.User.Create().Save(ctx)
		if err != nil {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Query email by email address
		email, err := client.Email.Query().Where(email.Address(address)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if email != nil {
			if !email.UserID.IsNil() {
				// The email already exists and is assigned already.
				return errorhandler.NewHTTPError(http.StatusConflict).SetInternal(fmt.Errorf("user with email %s already exists", address))
			}
		} else {
			if !svc.cfg.RequireEmailVerification {
				email, err = client.Email.Create().
					SetUserID(newU.ID).
					SetAddress(address).
					Save(ctx)
				if err != nil {
					return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				_, err = client.PrimaryEmail.Create().
					SetUserID(newU.ID).
					SetEmailID(email.ID).
					Save(ctx)
				if err != nil {
					return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			} else {
				email, err = client.Email.Create().
					SetAddress(address).
					Save(ctx)
				if err != nil {
					return errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			}
		}

		emailID = email.ID
		return nil
	}); err != nil {
		return newU, emailID, err
	}
	return newU, emailID, nil
}

func (svc *UserSvc) GetById(ctx Ctx, userID uuid.UUID) (*ent.User, *string, error) {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userID)
	if err != nil {
		return nil, nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return nil, nil, errorhandler.NewHTTPError(http.StatusNotFound).SetInternal(errors.New("user not found"))
	}

	var emailAddress *string
	if e := user.Edges.PrimaryEmail; e != nil {
		model, err := e.QueryEmail().Only(ctx)
		if err != nil {
			return nil, nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		emailAddress = &model.Address
	}
	return user, emailAddress, nil
}

func (svc *UserSvc) GetUserIdByEmail(ctx Ctx, addresss string) (*ent.Email, bool, error) {
	email, err := svc.repo.GetEmailRepo().GetByAddress(ctx, addresss)
	if err != nil {
		return nil, false, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if email == nil || email.UserID == nil {
		return nil, false, errorhandler.NewHTTPError(http.StatusNotFound).SetInternal(errors.New("user not found"))
	}

	credentials, err := svc.repo.GetWebauthnCredentialRepo().ListByUser(ctx, *email.UserID)
	if err != nil {
		return nil, false, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	hasCredentials := len(credentials) > 0
	return email, hasCredentials, nil
}
