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

type AssetHandler struct {
	*HttpDeps
}

func NewAssetHandler(srv *HttpDeps) *AssetHandler {
	return &AssetHandler{srv}
}

func (h *AssetHandler) ListByUser(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	result, err := h.GetAssetSvc().ListByUser(c.Request().Context(), userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *AssetHandler) Create(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body dto.AssetBodyRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	if err := c.Validate(body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	result, err := h.GetAssetSvc().Create(c.Request().Context(), userId, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *AssetHandler) Update(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body dto.AssetBodyRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	if err := c.Validate(body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	err = h.GetAssetSvc().Update(c.Request().Context(), userId, body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}

func (h *AssetHandler) Delete(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	assetIdStr := c.Param("assetId")
	assetId, err := uuid.FromString(assetIdStr)
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	err = h.GetAssetSvc().Delete(c.Request().Context(), userId, assetId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
