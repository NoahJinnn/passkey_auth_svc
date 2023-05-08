package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/internal/http/sharedDto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type WebauthnHandler struct {
	*HttpDeps
	sessionManager session.Manager
}

func NewWebauthnHandler(srv *HttpDeps, sessionManager session.Manager) *WebauthnHandler {
	return &WebauthnHandler{
		srv,
		sessionManager,
	}
}

// BeginRegistration returns credential creation options for the WebAuthnAPI. It expects a valid session JWT in the request.
func (h *WebauthnHandler) BeginRegistration(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	uId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse userId from JWT subject:%w", err)
	}

	options, err := h.GetWebauthnSvc().BeginRegistration(c.Request().Context(), uId)
	if err != nil {
		return fmt.Errorf("failed to create webauthn creation options: %w", err)
	}

	return c.JSON(http.StatusOK, options)
}

// FinishRegistration validates the WebAuthnAPI response and associates the credential with the user. It expects a valid session JWT in the request.
// The session JWT must be associated to the same user who requested the credential creation options.
func (h *WebauthnHandler) FinishRegistration(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	request, err := protocol.ParseCredentialCreationResponse(c.Request())
	if err != nil {
		errT, ok := err.(*protocol.Error)
		if ok {
			fmt.Printf("ParseCredentialCreationResponse err: %+v\n", errT.DevInfo)
		}
		return sharedDto.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	credentialId, userId, err := h.GetWebauthnSvc().FinishRegistration(
		c.Request().Context(),
		request,
		sessionToken.Subject(),
	)
	if err != nil {
		return sharedDto.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"credential_id": credentialId, "user_id": userId})
}

type BeginAuthenticationBody struct {
	UserID *string `json:"user_id" validate:"uuid4"`
}

// BeginLogin returns credential assertion options for the WebAuthnAPI.
func (h *WebauthnHandler) BeginLogin(c echo.Context) error {
	var request BeginAuthenticationBody

	if err := (&echo.DefaultBinder{}).BindBody(c, &request); err != nil {
		return sharedDto.ToHttpError(err)
	}

	options, err := h.GetWebauthnSvc().BeginLogin(c.Request().Context(), request.UserID)
	if err != nil {
		return sharedDto.ToHttpError(err)
	}

	return c.JSON(http.StatusOK, options)
}

// FinishLogin validates the WebAuthnAPI response and on success it returns a new session JWT.
func (h *WebauthnHandler) FinishLogin(c echo.Context) error {
	request, err := protocol.ParseCredentialRequestResponse(c.Request())
	if err != nil {
		errT, ok := err.(*protocol.Error)
		if ok {
			fmt.Printf("ParseCredentialRequestResponse err: %+v\n", errT.DevInfo)
		}
		return sharedDto.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	credentialId, userId, err := h.GetWebauthnSvc().FinishLogin(c.Request().Context(), request)
	if err != nil {
		return sharedDto.ToHttpError(err)
	}

	token, err := h.sessionManager.GenerateJWT(userId)
	if err != nil {
		return fmt.Errorf("failed to generate jwt: %w", err)
	}

	cookie, err := h.sessionManager.GenerateCookie(token)
	if err != nil {
		return fmt.Errorf("failed to create session cookie: %w", err)
	}

	c.SetCookie(cookie)

	if h.SharedCfg.Session.EnableAuthTokenHeader {
		c.Response().Header().Set("X-Auth-Token", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
	}

	return c.JSON(http.StatusOK, map[string]string{"credential_id": credentialId, "user_id": userId})
}
