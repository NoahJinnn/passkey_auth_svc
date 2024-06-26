package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/NoahJinnn/passkey_auth_svc/internal/http/errorhandler"
	"github.com/NoahJinnn/passkey_auth_svc/internal/http/session"
	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/srv/http/dto"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type UserHandler struct {
	*HttpDeps
	sessionManager session.IManager
}

func NewUserHandler(srv *HttpDeps, sessionManager session.IManager) *UserHandler {
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
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(body); err != nil {
		return errorhandler.ToHttpError(err)
	}

	body.Email = strings.ToLower(body.Email)

	newUser, emailId, err := h.GetUserSvc().Create(c.Request().Context(), body.Email)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}

	if !h.Cfg.RequireEmailVerification {
		token, err := h.sessionManager.GenerateJWT(newUser.ID.String())
		if err != nil {
			return fmt.Errorf("failed to generate jwt: %w", err)
		}

		cookie, err := h.sessionManager.GenerateCookie(token)
		if err != nil {
			return fmt.Errorf("failed to create session token: %w", err)
		}

		c.SetCookie(cookie)

		if h.SharedCfg.Session.EnableAuthTokenHeader {
			c.Response().Header().Set("X-Auth-Token", token)
			c.Response().Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
		}
	}

	return c.JSON(http.StatusOK, dto.CreateUserResponse{
		ID:      newUser.ID,
		UserID:  newUser.ID,
		EmailID: emailId,
	})
}

func (h *UserHandler) Get(c echo.Context) error {
	userId := c.Param("id")

	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("missing or malformed jwt")
	}

	if sessionToken.Subject() != userId {
		return errorhandler.NewHTTPError(http.StatusForbidden).SetInternal(fmt.Errorf("user %s tried to get user %s", sessionToken.Subject(), userId))
	}

	user, emailAddress, err := h.GetUserSvc().GetById(c.Request().Context(), uuid.FromStringOrNil(userId))
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	return c.JSON(http.StatusOK, dto.GetUserResponse{
		ID:                  user.ID,
		WebauthnCredentials: user.Edges.WebauthnCredentials,
		Email:               emailAddress,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	})
}

type UserGetByEmailBody struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *UserHandler) GetUserIdByEmail(c echo.Context) error {
	var request UserGetByEmailBody
	if err := (&echo.DefaultBinder{}).BindBody(c, &request); err != nil {
		return errorhandler.ToHttpError(err)
	}

	if err := c.Validate(request); err != nil {
		return errorhandler.ToHttpError(err)
	}

	emailAddress := strings.ToLower(request.Email)

	email, hasCredentials, err := h.GetUserSvc().GetUserIdByEmail(c.Request().Context(), emailAddress)
	if err != nil {
		return errorhandler.ToHttpError(err)
	}

	return c.JSON(http.StatusOK, dto.UserInfoResponse{
		ID:                    *email.UserID,
		Verified:              email.Verified,
		EmailID:               email.ID,
		HasWebauthnCredential: hasCredentials,
	})
}

func (h *UserHandler) Logout(c echo.Context) error {
	_, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("missing or malformed jwt")
	}

	// TODO: audit logger
	// userId := uuid.FromStringOrNil(sessionToken.Subject())

	// user, emailAddress, err := h.GetUserSvc().GetById(c.Request().Context(), uuid.FromStringOrNil(userId))
	// if err != nil {
	// 	return fmt.Errorf("failed to get user: %w", err)
	// }

	// err = h.auditLogger.Create(c, models.AuditLogUserLoggedOut, user, nil)
	// if err != nil {
	// 	return fmt.Errorf("failed to write audit log: %w", err)
	// }

	cookie, err := h.sessionManager.DeleteCookie()
	if err != nil {
		return fmt.Errorf("failed to delete session token: %w", err)
	}

	c.SetCookie(cookie)

	return c.NoContent(http.StatusNoContent)
}
