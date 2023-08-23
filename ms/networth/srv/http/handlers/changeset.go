package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type ChangesetHandler struct {
	*HttpDeps
}

func NewChangesetHandler(srv *HttpDeps) *ChangesetHandler {
	return &ChangesetHandler{srv}
}

func (h *ChangesetHandler) Delete(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	err = h.GetChangesetSvc().Delete(c.Request().Context(), userId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
