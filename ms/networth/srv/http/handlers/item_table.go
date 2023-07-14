package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type ItemTableHandler struct {
	*HttpDeps
}

func NewItemTableHandler(srv *HttpDeps) *ItemTableHandler {
	return &ItemTableHandler{srv}
}

func (h *ItemTableHandler) ListByUser(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	result, err := h.GetItemTableSvc().ListByUser(c.Request().Context(), userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *ItemTableHandler) Create(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body dto.ItemTableBody
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	result, err := h.GetItemTableSvc().Create(c.Request().Context(), userId, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *ItemTableHandler) Update(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body dto.ItemTableBody
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	err = h.GetItemTableSvc().Update(c.Request().Context(), userId, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}

func (h *ItemTableHandler) Delete(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	assetIdStr := c.Param("itemTableId")
	assetId, err := uuid.FromString(assetIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	err = h.GetItemTableSvc().Delete(c.Request().Context(), userId, assetId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
