package svcs

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/mail"
	"golang.org/x/crypto/bcrypt"
)

type IPasscodeSvc interface {
	InitPasscode(ctx Ctx, userId uuid.UUID, emailId uuid.UUID, acceptLang string) (*ent.Passcode, error)
	FinishPasscode(ctx Ctx, passcodeId uuid.UUID, reqCode string) (*ent.Passcode, error)
}

type passcodeSvc struct {
	repo              dal.IAuthRepo
	cfg               *config.Config
	passcodeGenerator PasscodeGenerator
	mailer            *mail.Mailer
	renderer          *mail.Renderer
}

var maxPasscodeTries = 3

func NewPasscodeSvc(cfg *config.Config, repo dal.IAuthRepo) IPasscodeSvc {
	mailer := mail.NewMailer(&cfg.Passcode)
	renderer, err := mail.NewRenderer()
	if err != nil {
		panic(fmt.Errorf("failed to create new renderer: %w", err))
	}
	return &passcodeSvc{
		repo:              repo,
		cfg:               cfg,
		passcodeGenerator: NewPasscodeGenerator(),
		mailer:            mailer,
		renderer:          renderer,
	}
}

func (svc *passcodeSvc) InitPasscode(ctx Ctx, userId uuid.UUID, emailId uuid.UUID, acceptLang string) (*ent.Passcode, error) {
	user, err := svc.repo.GetUserRepo().GetById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		// TODO: audit logger
		return nil, sharedDto.NewHTTPError(http.StatusBadRequest).SetInternal(errors.New("user not found"))
	}

	// if h.rateLimiter != nil {
	// 	err := rate_limiter.Limit(h.rateLimiter, userId, c)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// Determine where to send the passcode
	var email *ent.Email
	if !emailId.IsNil() {
		// Send the passcode to the specified email address
		email, err = svc.repo.GetEmailRepo().GetById(ctx, emailId)
		if err != nil {
			return nil, fmt.Errorf("failed to get email: %w", err)
		}
		if email == nil {
			return nil, sharedDto.NewHTTPError(http.StatusBadRequest, "the specified emailId is not available")
		}
	} else if e := user.Edges.PrimaryEmail; e == nil {
		// Can't determine email address to which the passcode should be sent to
		return nil, sharedDto.NewHTTPError(http.StatusBadRequest, "an emailId needs to be specified")
	} else {
		// Send the passcode to the primary email address
		email = e.Edges.Email
	}

	if email.Edges.User != nil && email.Edges.User.ID.String() != user.ID.String() {
		return nil, sharedDto.NewHTTPError(http.StatusForbidden).SetInternal(errors.New("email address is assigned to another user"))
	}

	passcode, err := svc.passcodeGenerator.Generate()
	if err != nil {
		return nil, fmt.Errorf("failed to generate passcode: %w", err)
	}

	hashedPasscode, err := bcrypt.GenerateFromPassword([]byte(passcode), 12)
	if err != nil {
		return nil, fmt.Errorf("failed to hash passcode: %w", err)
	}
	passcodeModel := &ent.Passcode{
		UserID:  userId,
		EmailID: email.ID,
		TTL:     svc.cfg.Passcode.TTL,
		Code:    string(hashedPasscode),
	}

	newPc, err := svc.repo.GetPasscodeRepo().Create(ctx, passcodeModel)
	if err != nil {
		return nil, fmt.Errorf("failed to store passcode: %w", err)
	}

	durationTTL := time.Duration(svc.cfg.Passcode.TTL) * time.Second
	data := map[string]interface{}{
		"Code":        passcode,
		"ServiceName": svc.cfg.ServiceName,
		"TTL":         fmt.Sprintf("%.0f", durationTTL.Minutes()),
	}

	str, err := svc.renderer.Render("loginTextMail", acceptLang, data)
	if err != nil {
		return nil, fmt.Errorf("failed to render email template: %w", err)
	}
	mailSubject := svc.renderer.Translate(acceptLang, "email_subject_login", data)
	err = svc.mailer.Send([]string{email.Address}, mailSubject, str)
	if err != nil {
		return nil, fmt.Errorf("failed to send passcode: %w", err)
	}

	return newPc, nil
}

func (svc *passcodeSvc) FinishPasscode(ctx Ctx, passcodeId uuid.UUID, reqCode string) (*ent.Passcode, error) {
	startTime := time.Now().UTC()
	var entPc *ent.Passcode
	// only if an internal server error occurs the transaction should be rolled back
	var businessError error
	if err := svc.repo.WithTx(ctx, func(ctx Ctx, client *ent.Client) error {
		passcodeRepo := svc.repo.GetPasscodeRepo()
		userRepo := svc.repo.GetUserRepo()

		passcode, err := passcodeRepo.GetById(ctx, passcodeId)
		if err != nil {
			return fmt.Errorf("failed to get passcode: %w", err)
		}

		if passcode == nil {
			// TODO: audit logger
			if err != nil {
				return fmt.Errorf("failed to create audit log: %w", err)
			}
			businessError = sharedDto.NewHTTPError(http.StatusUnauthorized, "passcode not found")
			return nil
		}

		user, err := userRepo.GetById(ctx, passcode.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
		lastVerificationTime := passcode.CreatedAt.Add(time.Duration(passcode.TTL) * time.Second)
		if lastVerificationTime.Before(startTime) {
			// TODO: audit logger
			if err != nil {
				return fmt.Errorf("failed to create audit log: %w", err)
			}
			businessError = sharedDto.NewHTTPError(http.StatusRequestTimeout, "passcode request timed out").SetInternal(fmt.Errorf("createdAt: %s -> lastVerificationTime: %s - current: %s", passcode.CreatedAt, lastVerificationTime, startTime)) // TODO: maybe we should use BadRequest, because RequestTimeout might be to technical and can refer to different error
			return nil
		}

		err = bcrypt.CompareHashAndPassword([]byte(passcode.Code), []byte(reqCode))
		if err != nil {
			passcode.TryCount = passcode.TryCount + 1

			if passcode.TryCount >= int32(maxPasscodeTries) {
				err = passcodeRepo.Delete(ctx, passcode)
				if err != nil {
					return fmt.Errorf("failed to delete passcode: %w", err)
				}
				// TODO: audit logger
				if err != nil {
					return fmt.Errorf("failed to create audit log: %w", err)
				}
				businessError = sharedDto.NewHTTPError(http.StatusGone, "max attempts reached")
				return nil
			}

			err = passcodeRepo.Update(ctx, passcode)
			if err != nil {
				return fmt.Errorf("failed to update passcode: %w", err)
			}

			// TODO: audit logger
			if err != nil {
				return fmt.Errorf("failed to create audit log: %w", err)
			}
			businessError = sharedDto.NewHTTPError(http.StatusUnauthorized).SetInternal(errors.New("passcode invalid"))
			return nil
		}

		err = passcodeRepo.Delete(ctx, passcode)
		if err != nil {
			return fmt.Errorf("failed to delete passcode: %w", err)
		}

		if passcode.Edges.User != nil && passcode.Edges.Email.UserID.String() != user.ID.String() {
			return sharedDto.NewHTTPError(http.StatusForbidden, "email address has been claimed by another user")
		}
		entPc = passcode
		return nil
	}); err != nil {
		return nil, err
	}
	if businessError != nil {
		return nil, businessError
	}
	return entPc, nil
}
