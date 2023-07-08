package finverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
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
	config *config.Config
	req    *httpx.Req
	repo   dal.INwRepo
}

func NewFvDataSvc(cfg *config.Config, repo dal.INwRepo) *FvDataSvc {
	req := httpx.NewReq("https://api.sandbox.finverse.net/", map[string]string{
		"Content-Type": "application/json",
	}, nil)

	return &FvDataSvc{config: cfg, req: req, repo: repo}
}

func (svc *FvDataSvc) AllInstitution(ctx context.Context, userId uuid.UUID) ([]interface{}, error) {
	fvSession, err := svc.repo.GetFvSessionRepo().GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
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

	var result []interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv institutions: %w", err))
	}

	return result, nil
}

func (svc *FvDataSvc) AllAccount(ctx context.Context, userId uuid.UUID) (interface{}, error) {
	resp, err := svc.req.
		InitReq(ctx, "GET", "/accounts", nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv accounts: %w", err))
	}

	return result, nil
}

func (svc *FvDataSvc) AllTransactions(ctx context.Context, offset string, limit string, userId uuid.UUID) (interface{}, error) {
	var queryStr = ""
	if offset != "" && limit != "" {
		queryStr = fmt.Sprintf("?offset=%s&limit=%s", offset, limit)
	} else if limit != "" {
		queryStr = fmt.Sprintf("?limit=%s", limit)
	} else if offset != "" {
		queryStr = fmt.Sprintf("?offset=%s", offset)
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/transactions"+queryStr, nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv balance history: %w", err))
	}

	return result, nil
}

func (svc *FvDataSvc) GetBalanceHistoryByAccountId(ctx context.Context, accountId string, userId uuid.UUID) (interface{}, error) {
	resp, err := svc.req.
		InitReq(ctx, "GET", "/balance_history/"+accountId, nil).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken,
		}).
		Send()
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv balance history: %w", err))
	}

	return result, nil
}