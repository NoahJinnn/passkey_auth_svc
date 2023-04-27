package svcs

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/mail"
	"golang.org/x/crypto/bcrypt"
)

type IPasscodeSvc interface {
	InitPasscode(ctx Ctx, userId uuid.UUID, emailId uuid.UUID, acceptLang string) (*ent.Passcode, error)
}

type passcodeSvc struct {
	repo              dal.IRepo
	cfg               *config.Config
	passcodeGenerator PasscodeGenerator
	mailer            *mail.Mailer
	renderer          *mail.Renderer
}

func NewPasscodeSvc(cfg *config.Config, repo dal.IRepo) IPasscodeSvc {
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
		return nil, dto.NewHTTPError(http.StatusBadRequest).SetInternal(errors.New("user not found"))
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
		if email == nil {
			return nil, dto.NewHTTPError(http.StatusBadRequest, "the specified emailId is not available")
		}
	} else if e := user.Edges.PrimaryEmail; e == nil {
		// Can't determine email address to which the passcode should be sent to
		return nil, dto.NewHTTPError(http.StatusBadRequest, "an emailId needs to be specified")
	} else {
		// Send the passcode to the primary email address
		email = e.Edges.Email
	}

	if email.Edges.User != nil && email.Edges.User.ID.String() != user.ID.String() {
		return nil, dto.NewHTTPError(http.StatusForbidden).SetInternal(errors.New("email address is assigned to another user"))
	}

	passcode, err := svc.passcodeGenerator.Generate()
	if err != nil {
		return nil, fmt.Errorf("failed to generate passcode: %w", err)
	}

	passcodeId, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to create passcodeId: %w", err)
	}
	now := time.Now().UTC()
	hashedPasscode, err := bcrypt.GenerateFromPassword([]byte(passcode), 12)
	if err != nil {
		return nil, fmt.Errorf("failed to hash passcode: %w", err)
	}
	passcodeModel := &ent.Passcode{
		ID:        passcodeId,
		UserID:    userId,
		EmailID:   email.ID,
		TTL:       svc.cfg.Passcode.TTL,
		Code:      string(hashedPasscode),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = svc.repo.GetPasscodeRepo().Create(ctx, passcodeModel)
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

	return passcodeModel, nil
}
