package finverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/provider"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	"github.com/hellohq/hqservice/pkg/httpx"
)

type IFvDataSvc interface {
	AllInstitution(ctx context.Context, userId uuid.UUID) ([]interface{}, error)
	AllAccount(ctx context.Context, userId uuid.UUID) (interface{}, error)
	AllTransactions(ctx context.Context, userId uuid.UUID) (interface{}, error)
	GetBalanceHistoryByAccountId(ctx context.Context, accountId string, userId uuid.UUID) (interface{}, error)
}

type FvDataSvc struct {
	req      *httpx.Req
	repo     dal.INwRepo
	provider *provider.ProviderSvc
}

func NewFvDataSvc(cfg *config.Config, provider *provider.ProviderSvc, repo dal.INwRepo) *FvDataSvc {
	req := httpx.NewReq("https://api.sandbox.finverse.net/", map[string]string{
		"Content-Type": "application/json",
	}, nil)

	return &FvDataSvc{req: req, repo: repo, provider: provider}
}

func (svc *FvDataSvc) AllInstitution(ctx context.Context, userId uuid.UUID) ([]byte, error) {
	fvSession, err := svc.repo.GetFvSessionRepo().GetByUserId(ctx, userId)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/institutions", nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + fvSession.AccessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (svc *FvDataSvc) GetAccessToken(ctx context.Context, providerName string, userId uuid.UUID) (*AccessToken, error) {
	var accessToken *AccessToken
	accessTokenPayload, err := svc.provider.ConnectionByProviderName(ctx, userId.String(), providerName)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv connection: %w", err))
	}

	if accessTokenPayload == nil {
		return nil, errorhandler.NewHTTPError(http.StatusNotFound, "confidential connection not found")
	}

	err = json.Unmarshal([]byte(accessTokenPayload.Data), &accessToken)

	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return accessToken, nil
}

func (svc *FvDataSvc) AllAccount(ctx context.Context, userId uuid.UUID) ([]byte, error) {
	accessToken, err := svc.GetAccessToken(ctx, PROVIDER_NAME, userId)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/accounts", nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken.AccessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (svc *FvDataSvc) AllTransactions(ctx context.Context, offset string, limit string, userId uuid.UUID) ([]byte, error) {
	var queryStr = ""
	if offset != "" && limit != "" {
		queryStr = fmt.Sprintf("?offset=%s&limit=%s", offset, limit)
	} else if limit != "" {
		queryStr = fmt.Sprintf("?limit=%s", limit)
	} else if offset != "" {
		queryStr = fmt.Sprintf("?offset=%s", offset)
	}

	accessToken, err := svc.GetAccessToken(ctx, PROVIDER_NAME, userId)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/transactions"+queryStr, nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken.AccessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func (svc *FvDataSvc) GetBalanceHistoryByAccountId(ctx context.Context, accountId string, userId uuid.UUID) ([]byte, error) {
	accessToken, err := svc.GetAccessToken(ctx, PROVIDER_NAME, userId)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/balance_history/"+accountId, nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken.AccessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

/** ===== Aggregator ===== **/
func (svc *FvDataSvc) AggregateAccountBalances(ctx context.Context, userId uuid.UUID) ([]interface{}, error) {
	allAccount, err := svc.AllAccount(ctx, userId)
	if err != nil {
		return nil, err
	}

	var accounts *Accounts
	err = json.Unmarshal(allAccount, &accounts)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var accountBalances []interface{}
	for _, account := range accounts.Accounts {
		bh, err := svc.GetBalanceHistoryByAccountId(ctx, account.AccountID, userId)
		if err != nil {
			return nil, err
		}

		var balance interface{}
		err = json.Unmarshal(bh, &balance)
		if err != nil {
			return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		accountBalances = append(accountBalances, balance)
	}

	err = svc.provider.SaveAccount(ctx, "finverse", userId.String(), accountBalances)
	if err != nil {
		return nil, errorhandler.
			NewHTTPError(http.StatusInternalServerError).
			SetInternal(fmt.Errorf("failed to save fv exchange token to sqlite: %w", err))
	}

	return accountBalances, nil
}
