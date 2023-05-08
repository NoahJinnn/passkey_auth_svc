package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"

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
		return sharedDto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return sharedDto.ToHttpError(err)
	}

	userId, err := uuid.FromString(body.UserId)
	if err != nil {
		return sharedDto.NewHTTPError(http.StatusBadRequest, "failed to parse userId as uuid").SetInternal(err)
	}

	var emailId uuid.UUID
	if body.EmailId != nil {
		emailId, err = uuid.FromString(*body.EmailId)
		if err != nil {
			return sharedDto.NewHTTPError(http.StatusBadRequest, "failed to parse emailId as uuid").SetInternal(err)
		}
	}
	lang := c.Request().Header.Get("Accept-Language")
	passcodeEnt, err := h.GetPasscodeSvc().InitPasscode(c.Request().Context(), userId, emailId, lang)
	if err != nil {
		return sharedDto.ToHttpError(err)
	}

	return c.JSON(http.StatusOK, dto.PasscodeReturn{
		Id:        passcodeEnt.ID.String(),
		TTL:       h.Cfg.Passcode.TTL,
		CreatedAt: passcodeEnt.CreatedAt,
	})
}

func (h *PasscodeHandler) Finish(c echo.Context) error {
	var body dto.PasscodeFinishRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return sharedDto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return sharedDto.ToHttpError(err)
	}

	passcodeId, err := uuid.FromString(body.Id)
	if err != nil {
		return sharedDto.NewHTTPError(http.StatusBadRequest, "failed to parse passcodeId as uuid").SetInternal(err)
	}

	passcode, err := h.GetPasscodeSvc().FinishPasscode(c.Request().Context(), passcodeId, body.Code)
	if err != nil {
		return err
	}
	token, err := h.sessionManager.GenerateJWT(passcode.UserID.String())
	if err != nil {
		return fmt.Errorf("failed to generate jwt: %w", err)
	}

	cookie, err := h.sessionManager.GenerateCookie(token)
	if err != nil {
		return fmt.Errorf("failed to create session token: %w", err)
	}

	c.SetCookie(cookie)

	if h.SharedCfg.Session.EnableAuthTokenHeader {
		c.Response().Header().Set("X-Auth-Token", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
	}

	// TODO: audit logger

	return c.JSON(http.StatusOK, dto.PasscodeReturn{
		Id:        passcode.ID.String(),
		TTL:       passcode.TTL,
		CreatedAt: passcode.CreatedAt,
	})
}
