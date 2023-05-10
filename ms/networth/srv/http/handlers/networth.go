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
	resp, err := h.GetSeSvc().CreateCountries()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"nw": resp})
}
