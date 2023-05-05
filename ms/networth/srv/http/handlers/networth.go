package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type NetworthHandler struct{}

func NewNetworthHandler() *NetworthHandler {
	return &NetworthHandler{}
}

func (handler *NetworthHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"nw": "success"})
}
