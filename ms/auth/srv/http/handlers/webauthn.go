package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type WebauthnHandler struct {
	*HttpDeps
	sessionManager session.IManager
}

func NewWebauthnHandler(srv *HttpDeps, sessionManager session.IManager) *WebauthnHandler {
	return &WebauthnHandler{
		srv,
		sessionManager,
	}
}

// InitRegistration returns credential creation options for the WebAuthnAPI. It expects a valid session JWT in the request.
func (h *WebauthnHandler) InitRegistration(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		httperr := errorhandler.NewHTTPError(http.StatusBadRequest, fmt.Errorf("failed to cast session object").Error())
		return c.JSON(httperr.Code, httperr)
	}

	uId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	options, err := h.GetWebauthnSvc().InitRegistration(c.Request().Context(), uId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
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
		httperr := errorhandler.NewHTTPError(http.StatusBadRequest, err.Error())
		return c.JSON(httperr.Code, httperr)
	}

	credentialId, userId, err := h.GetWebauthnSvc().FinishRegistration(
		c.Request().Context(),
		request,
		sessionToken.Subject(),
	)
	if err != nil {
		httperr := errorhandler.NewHTTPError(http.StatusBadRequest, err.Error())
		return c.JSON(httperr.Code, httperr)
	}
	return c.JSON(http.StatusOK, map[string]string{"credential_id": credentialId, "user_id": userId})
}

type BeginAuthenticationBody struct {
	UserID *string `json:"user_id" validate:"uuid4"`
}

// InitLogin returns credential assertion options for the WebAuthnAPI.
func (h *WebauthnHandler) InitLogin(c echo.Context) error {
	var request BeginAuthenticationBody

	if err := (&echo.DefaultBinder{}).BindBody(c, &request); err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	options, err := h.GetWebauthnSvc().InitLogin(c.Request().Context(), request.UserID)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
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
		httperr := errorhandler.NewHTTPError(http.StatusBadRequest, err.Error())
		return c.JSON(httperr.Code, httperr)
	}

	credentialId, userId, err := h.GetWebauthnSvc().FinishLogin(c.Request().Context(), request)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	token, err := h.sessionManager.GenerateJWT(userId)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	cookie, err := h.sessionManager.GenerateCookie(token)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	c.SetCookie(cookie)

	if h.SharedCfg.Session.EnableAuthTokenHeader {
		c.Response().Header().Set("X-Auth-Token", token)
		c.Response().Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
	}

	return c.JSON(http.StatusOK, map[string]string{"credential_id": credentialId, "user_id": userId})
}
