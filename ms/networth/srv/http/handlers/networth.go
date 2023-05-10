package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type NetworthHandler struct {
	*HttpDeps
}

func NewNetworthHandler(srv *HttpDeps) *NetworthHandler {
	return &NetworthHandler{srv}
}

func (h *NetworthHandler) Get(c echo.Context) error {
	h.GetSeSvc()
	return c.JSON(http.StatusOK, map[string]string{"nw": "success"})
}
