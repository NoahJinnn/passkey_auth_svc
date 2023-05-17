package handlers

import (
	"net/http"

	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/labstack/echo/v4"
)

type FvAuthHandler struct {
	*HttpDeps
}

func NewFvAuthHandler(srv *HttpDeps) *FvAuthHandler {
	return &FvAuthHandler{srv}
}

func (h *FvAuthHandler) CreateCustomerToken(c echo.Context) error {
	body := finverse.CreateCustomerToken{
		ClientID:     h.Cfg.Finverse.ClientID,
		ClientSecret: h.Cfg.Finverse.Secret,
		GrantType:    "client_credentials",
	}
	token, err := h.GetFvAuthSvc().CreateCustomerToken(c.Request().Context(), &body)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}
	return c.JSON(http.StatusOK, token)
}

func (h *FvAuthHandler) CreateLinkToken(c echo.Context) error {
	body := finverse.CreateLinkToken{
		ClientID:    h.Cfg.Finverse.ClientID,
		RedirectURI: h.Cfg.Finverse.RedirectURI,
		State:       h.Cfg.Finverse.AppId + "-stateparameter",
		GrantType:   "client_credentials",
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	token, err := h.GetFvAuthSvc().CreateLinkToken(c.Request().Context(), &body)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}
	return c.JSON(http.StatusOK, token)
}

func (h *FvAuthHandler) ExchangeAccessToken(c echo.Context) error {
	var body finverse.ExchangeAccessToken
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}
	token, err := h.GetFvAuthSvc().ExchangeAccessToken(c.Request().Context(), body.Code)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}
	return c.JSON(http.StatusOK, token)
}
