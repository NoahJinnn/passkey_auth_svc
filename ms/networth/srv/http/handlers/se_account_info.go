package handlers

import (
	"net/http"

	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/hellohq/hqservice/ms/networth/app/dom"
	"github.com/labstack/echo/v4"
)

type NetworthHandler struct {
	*HttpDeps
}

func NewSeHandler(srv *HttpDeps) *NetworthHandler {
	return &NetworthHandler{srv}
}

func (h *NetworthHandler) CreateCustomer(c echo.Context) error {
	var body dom.CreateCustomerReq
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return sharedDto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return sharedDto.ToHttpError(err)
	}
	resp, err := h.GetSeAccountInfoSvc().CreateCustomer(c.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) CreateConnectSession(c echo.Context) error {
	var body dom.CreateConnectSessionReq
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return sharedDto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return sharedDto.ToHttpError(err)
	}
	resp, err := h.GetSeAccountInfoSvc().CreateConnectSession(c.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
