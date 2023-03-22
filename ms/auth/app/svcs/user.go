package svcs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
)

type IUserSvc interface {
	Create(ctx Ctx, email string) (newU *ent.User, emailID uuid.UUID, err error)
}

type userSvc struct {
	repo *dal.Repo
}

func NewUserSvc(repo *dal.Repo) IUserSvc {
	return &userSvc{
		repo: repo,
	}
}

func (svc *userSvc) Create(ctx Ctx, address string) (newU *ent.User, emailID uuid.UUID, err error) {

	if err := svc.repo.WithTx(ctx, func(tx *ent.Tx) error {
		exec := (func(ctx Ctx, client *ent.Client) error {
			newU, err = client.User.Create().Save(ctx)
			if err != nil {
				return fmt.Errorf("failed creating user: %w", err)
			}

			// Query email by email address
			email, err := client.Email.Query().Where(email.Address(address)).Only(ctx)
			if err != nil {
				return fmt.Errorf("failed querying email by address: %w", err)
			}

			if email != nil {
				if !email.UserID.IsNil() {
					// The email already exists and is assigned already.
					return dto.NewHTTPError(http.StatusConflict).SetInternal(fmt.Errorf("user with email %s already exists", address))
				}
			} else {
				// Create new email
				email, err = client.Email.Create().SetAddress(address).SetUserID(newU.ID).Save(ctx)
				if err != nil {
					return fmt.Errorf("failed creating email: %w", err)
				}
			}
			emailID = email.ID
			return nil
		})
		return exec(ctx, tx.Client())
	}); err != nil {
		log.Fatal(err)
	}
	return newU, emailID, nil
}
