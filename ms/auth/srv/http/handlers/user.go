package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/hellohq/hqservice/ms/auth/srv/http/session"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	*HttpDeps
	sessionManager session.Manager
}

func NewUserHandler(srv *HttpDeps, sessionManager session.Manager) *UserHandler {
	return &UserHandler{
		srv,
		sessionManager,
	}
}

type UserCreateBody struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *UserHandler) Create(c echo.Context) error {
	var body UserCreateBody
	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return dto.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return dto.ToHttpError(err)
	}

	body.Email = strings.ToLower(body.Email)

	newUser, emailId, err := h.GetUserSvc().Create(c.Request().Context(), body.Email)
	if err != nil {
		return dto.ToHttpError(err)
	}

	token, err := h.sessionManager.GenerateJWT(newUser.ID)
	if err != nil {
		return fmt.Errorf("failed to generate jwt: %w", err)
	}

	cookie, err := h.sessionManager.GenerateCookie(token)
	if err != nil {
		return fmt.Errorf("failed to create session token: %w", err)
	}

	c.SetCookie(cookie)

	if h.Cfg.Session.EnableAuthTokenHeader && !h.Cfg.Emails.RequireVerification {
		c.Response().Header().Set("X-Auth-Token", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
	}

	return c.JSON(http.StatusOK, dto.CreateUserResponse{
		ID:      newUser.ID,
		UserID:  newUser.ID,
		EmailID: emailId,
	})
}
