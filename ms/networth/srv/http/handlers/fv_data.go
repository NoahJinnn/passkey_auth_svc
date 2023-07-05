package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type FvDataHandler struct {
	*HttpDeps
}

func NewFvDataHandler(srv *HttpDeps) *FvDataHandler {
	return &FvDataHandler{srv}
}

func (h *FvDataHandler) AllInstitution(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	instis, err := h.GetFvDataSvc().AllInstitution(c.Request().Context(), userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, instis)
}

func (h *FvDataHandler) AllAccount(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	instis, err := h.GetFvDataSvc().AllAccount(c.Request().Context(), userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, instis)
}
