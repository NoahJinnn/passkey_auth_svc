package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type WebauthnCredentialHandler struct {
	*HttpDeps
}

func (h *WebauthnHandler) ListCredentials(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	credentials, err := h.GetWebauthnSvc().ListCredentials(c.Request().Context(), userId)
	if err != nil {
		return fmt.Errorf("failed to get webauthn credentials: %w", err)
	}

	resp := make([]*dto.WebauthnCredentialResponse, len(credentials))

	for i := range credentials {
		resp[i] = dto.FromWebauthnCredentialModel(credentials[i])
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *WebauthnHandler) UpdateCredential(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	credentialID := c.Param("id")

	var body dto.WebauthnCredentialUpdateRequest

	err = (&echo.DefaultBinder{}).BindBody(c, &body)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}

	return h.GetWebauthnSvc().UpdateCredential(c.Request().Context(), userId, credentialID, body.Name)
}

func (h *WebauthnHandler) DeleteCredential(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	credentialId := c.Param("id")

	err = h.GetWebauthnSvc().DeleteCredential(c.Request().Context(), userId, credentialId)
	if err != nil {
		return fmt.Errorf("to delete credential from db: %w", err)
	}

	return c.NoContent(http.StatusNoContent)
}
