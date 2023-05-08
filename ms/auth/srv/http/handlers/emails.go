package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type EmailHandler struct {
	*HttpDeps
	sessionManager session.Manager
}

func NewEmailHandler(srv *HttpDeps, sessionManager session.Manager) *EmailHandler {
	return &EmailHandler{
		srv,
		sessionManager,
	}
}

func (h *EmailHandler) ListByUser(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	emails, err := h.GetEmailSvc().ListByUser(c.Request().Context(), userId)
	if err != nil {
		return fmt.Errorf("failed to fetch emails from db: %w", err)
	}

	response := make([]*dto.EmailResponse, len(emails))

	for i := range emails {
		response[i] = dto.FromEmailModel(emails[i])
	}

	return c.JSON(http.StatusOK, response)
}
