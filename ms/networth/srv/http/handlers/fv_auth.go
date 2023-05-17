package handlers

import (
	"net/http"

	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/labstack/echo/v4"
)

type FvAuthHandler struct {
	*HttpDeps
}

func NewFvAuthHandler(srv *HttpDeps) *FvAuthHandler {
	return &FvAuthHandler{srv}
}

func (h *FvAuthHandler) CreateCustomerToken(c echo.Context) error {
	token, err := h.GetFvAuthSvc().CreateCustomerToken(c.Request().Context())
	if err != nil {
		return errorhandler.ToHttpError(err)
	}
	return c.JSON(http.StatusOK, token)
}
