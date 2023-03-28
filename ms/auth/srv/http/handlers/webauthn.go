package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type WebauthnHandler struct {
	*HttpDeps
}

func NewWebauthnHandler(srv *HttpDeps) *WebauthnHandler {
	return &WebauthnHandler{
		srv,
	}
}

func (h *WebauthnHandler) BeginRegistration(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	uId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse userId from JWT subject:%w", err)
	}

	options, err := h.GetWebauthnSvc().WebauthnBeginRegistration(c.Request().Context(), uId)
	if err != nil {
		return fmt.Errorf("failed to create webauthn creation options: %w", err)
	}

	return c.JSON(http.StatusOK, options)
}

func (h *WebauthnHandler) FinishRegistration(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	request, err := protocol.ParseCredentialCreationResponse(c.Request())
	if err != nil {
		return dto.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	credentialId, userId, err := h.Appl.GetWebauthnSvc().WebauthnFinishRegistration(c.Request().Context(), request, sessionToken.Subject())

	return c.JSON(http.StatusOK, map[string]string{"credential_id": credentialId, "user_id": userId})
}
