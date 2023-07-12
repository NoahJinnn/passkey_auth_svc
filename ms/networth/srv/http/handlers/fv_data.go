package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/finverse"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type FvDataHandler struct {
	*HttpDeps
}

func NewFvDataHandler(srv *HttpDeps) *FvDataHandler {
	return &FvDataHandler{srv}
}

func (h *FvDataHandler) AllInstitution(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	instis, err := h.GetFvDataSvc().AllInstitution(c.Request().Context(), userId)
	if err != nil {
		return err
	}

	var result []interface{}
	err = json.Unmarshal(instis, &result)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *FvDataHandler) AllAccount(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	exist, err := h.GetProviderSvc().CheckAccountExist(c.Request().Context(), userId.String(), finverse.PROVIDER_NAME)
	if err != nil {
		return err
	}

	if !exist {
		_, err = h.GetFvDataSvc().AggregateAccountBalances(c.Request().Context(), userId)
		if err != nil {
			return err
		}
	}

	// Get account data from sqlite db
	a, err := h.GetProviderSvc().AccountByProviderName(c.Request().Context(), userId.String(), finverse.PROVIDER_NAME)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a.Data)
}

func (h *FvDataHandler) AllTransaction(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")

	txs, err := h.GetFvDataSvc().AllTransactions(c.Request().Context(), offset, limit, userId)
	if err != nil {
		return err
	}

	var result interface{}
	err = json.Unmarshal(txs, &result)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	return c.JSON(http.StatusOK, txs)
}

func (h *FvDataHandler) GetBalanceHistoryByAccountId(c echo.Context) error {
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}

	userId, err := uuid.FromString(sessionToken.Subject())
	if err != nil {
		return fmt.Errorf("failed to parse subject as uuid: %w", err)
	}

	accountId := c.Param("accountId")

	balance, err := h.GetFvDataSvc().GetBalanceHistoryByAccountId(c.Request().Context(), accountId, userId)
	if err != nil {
		return err
	}

	var result interface{}
	err = json.Unmarshal(balance, &result)
	if err != nil {
		return errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.JSON(http.StatusOK, result)
}
