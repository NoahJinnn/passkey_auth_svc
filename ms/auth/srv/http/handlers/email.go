package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/internal/http/session"
	"github.com/hellohq/hqservice/ms/auth/srv/http/dto"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type EmailHandler struct {
	*HttpDeps
	sessionManager session.IManager
}

func NewEmailHandler(srv *HttpDeps, sessionManager session.IManager) *EmailHandler {
	return &EmailHandler{
		srv,
		sessionManager,
	}
}

func (h *EmailHandler) ListByUser(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	emails, err := h.GetEmailSvc().ListByUser(c.Request().Context(), userId)
	if err != nil {
		return err
	}

	resp := make([]*dto.EmailResponse, len(emails))

	for i := range emails {
		resp[i] = dto.FromEmailModel(emails[i])
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *EmailHandler) Create(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	var body dto.EmailCreateRequest
	err = (&echo.DefaultBinder{}).BindBody(c, &body)
	if err != nil {
		httperr := errorhandler.ToHttpError(err)
		return c.JSON(httperr.Code, httperr)
	}

	newEmailAddress := strings.ToLower(body.Address)

	email, err := h.GetEmailSvc().Create(c.Request().Context(), userId, newEmailAddress)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, email)
}

func (h *EmailHandler) SetPrimaryEmail(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	emailId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}

	err = h.GetEmailSvc().SetPrimaryEmail(c.Request().Context(), userId, emailId)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *EmailHandler) Delete(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	emailId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}

	err = h.GetEmailSvc().Delete(c.Request().Context(), userId, emailId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
