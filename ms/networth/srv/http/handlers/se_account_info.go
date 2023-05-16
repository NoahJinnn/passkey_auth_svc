package handlers

import (
	"net/http"

	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/labstack/echo/v4"
)

type NetworthHandler struct {
	*HttpDeps
}

func NewSeHandler(srv *HttpDeps) *NetworthHandler {
	return &NetworthHandler{srv}
}

func (h *NetworthHandler) Customer(c echo.Context) error {
	cId := c.Param("customer_id")

	resp, err := h.GetSeAccountInfoSvc().GetConnectionByCustomerId(c.Request().Context(), cId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) CreateCustomer(c echo.Context) error {
	var body saltedge.CreateCustomer
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}
	resp, err := h.GetSeAccountInfoSvc().CreateCustomer(c.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) DeleteCustomer(c echo.Context) error {
	cId := c.Param("customer_id")

	resp, err := h.GetSeAccountInfoSvc().RemoveCustomer(c.Request().Context(), cId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) CreateConnectSession(c echo.Context) error {
	var body saltedge.CreateConnectSession
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}
	resp, err := h.GetSeAccountInfoSvc().CreateConnectSession(c.Request().Context(), &body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) GetConnectionByCustomerId(c echo.Context) error {
	cId := c.QueryParam("customer_id")

	resp, err := h.GetSeAccountInfoSvc().GetConnectionByCustomerId(c.Request().Context(), cId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) GetAccountByConnectionId(c echo.Context) error {
	connId := c.QueryParam("connection_id")

	resp, err := h.GetSeAccountInfoSvc().GetAccountByConnectionId(c.Request().Context(), connId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *NetworthHandler) GetTxByConnectionIdAndAccountId(c echo.Context) error {
	connId := c.QueryParam("connection_id")
	aId := c.QueryParam("account_id")

	resp, err := h.GetSeAccountInfoSvc().GetTxByConnectionIdAndAccountId(c.Request().Context(), connId, aId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
