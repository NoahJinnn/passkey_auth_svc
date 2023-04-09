package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/labstack/echo/v4"
)

type PasscodeHandler struct {
	*HttpDeps
	sessionManager session.Manager
}

func NewPasscodeHandler(srv *HttpDeps, sessionManager session.Manager) *PasscodeHandler {
	return &PasscodeHandler{
		srv,
		sessionManager,
	}
}

func (h *PasscodeHandler) Init(c echo.Context) error {
	var body dto.PasscodeInitRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return dto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return dto.ToHttpError(err)
	}

	userId, err := uuid.FromString(body.UserId)
	if err != nil {
		return dto.NewHTTPError(http.StatusBadRequest, "failed to parse userId as uuid").SetInternal(err)
	}

	var emailId uuid.UUID
	if body.EmailId != nil {
		emailId, err = uuid.FromString(*body.EmailId)
		if err != nil {
			return dto.NewHTTPError(http.StatusBadRequest, "failed to parse emailId as uuid").SetInternal(err)
		}
	}

	_, passcodeEnt, err := h.GetPasscodeSvc().InitPasscode(c.Request().Context(), userId, emailId)
	if err != nil {
		return dto.ToHttpError(err)
	}

	// durationTTL := time.Duration(h.Cfg.Passcode.TTL) * time.Second
	// data := map[string]interface{}{
	// 	"Code":        passcode,
	// 	"ServiceName": h.Cfg.ServiceName,
	// 	"TTL":         fmt.Sprintf("%.0f", durationTTL.Minutes()),
	// }

	// lang := c.Request().Header.Get("Accept-Language")
	// str, err := h.renderer.Render("loginTextMail", lang, data)
	// if err != nil {
	// 	return fmt.Errorf("failed to render email template: %w", err)
	// }

	// TODO: Impl Email sender
	// message := gomail.NewMessage()
	// message.SetAddressHeader("To", email.Address, "")
	// message.SetAddressHeader("From", h.emailConfig.FromAddress, h.emailConfig.FromName)

	// message.SetHeader("Subject", h.renderer.Translate(lang, "email_subject_login", data))

	// message.SetBody("text/plain", str)

	// err = h.mailer.Send(message)
	// if err != nil {
	// 	return fmt.Errorf("failed to send passcode: %w", err)
	// }

	return c.JSON(http.StatusOK, dto.PasscodeReturn{
		Id:        passcodeEnt.ID.String(),
		TTL:       h.Cfg.Passcode.TTL,
		CreatedAt: passcodeEnt.CreatedAt,
	})
}
