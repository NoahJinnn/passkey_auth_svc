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
	lang := c.Request().Header.Get("Accept-Language")
	passcodeEnt, err := h.GetPasscodeSvc().InitPasscode(c.Request().Context(), userId, emailId, lang)
	if err != nil {
		return dto.ToHttpError(err)
	}

	return c.JSON(http.StatusOK, dto.PasscodeReturn{
		Id:        passcodeEnt.ID.String(),
		TTL:       h.Cfg.Passcode.TTL,
		CreatedAt: passcodeEnt.CreatedAt,
	})
}
