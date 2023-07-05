package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type FvAuthHandler struct {
	*HttpDeps
}

func NewFvAuthHandler(srv *HttpDeps) *FvAuthHandler {
	return &FvAuthHandler{srv}
}

type CreateCustomerTokenResp struct {
	GrantType string `json:"grant_type"`
	IsSuccess bool   `json:"is_success"`
}

func (h *FvAuthHandler) CreateCustomerToken(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	grantType := "client_credentials"
	body := finverse.CreateCustomerToken{
		ClientID:     h.Cfg.Finverse.ClientID,
		ClientSecret: h.Cfg.Finverse.Secret,
		GrantType:    grantType,
	}
	isSuccess, err := h.GetFvAuthSvc().CreateCustomerToken(c.Request().Context(), &body, &userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, CreateCustomerTokenResp{
		IsSuccess: isSuccess,
		GrantType: grantType,
	})
}

func (h *FvAuthHandler) CreateLinkToken(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	body := finverse.CreateLinkToken{
		ClientID:    h.Cfg.Finverse.ClientID,
		RedirectURI: h.Cfg.Finverse.RedirectURI,
		State:       h.Cfg.Finverse.AppId + "-stateparameter",
		GrantType:   "client_credentials",
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	if err := c.Validate(body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	token, err := h.GetFvAuthSvc().CreateLinkToken(c.Request().Context(), &body, userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, token)
}

func (h *FvAuthHandler) ExchangeAccessToken(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body finverse.ExchangeAccessToken
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	if err := c.Validate(body); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	token, err := h.GetFvAuthSvc().ExchangeAccessToken(c.Request().Context(), body.Code, userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, token)
}
