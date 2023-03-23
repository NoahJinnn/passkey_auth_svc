package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
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
