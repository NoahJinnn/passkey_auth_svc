package finverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	Income(ctx context.Context, userId uuid.UUID) (interface{}, error)
	PagingTransaction(ctx context.Context, userId uuid.UUID) (interface{}, error)
	GetBalanceHistoryByAccountId(ctx context.Context, accountId string, userId uuid.UUID) (interface{}, error)
	AggregateAccountBalances(ctx context.Context, userId uuid.UUID) ([]interface{}, error)
	AggregateTransactions(ctx context.Context, userId uuid.UUID) (interface{}, error)
	getAccessToken(ctx context.Context, providerName string, userId uuid.UUID) (*AccessToken, error)
	concatTransactions(ctx context.Context, userId uuid.UUID, offset int, limit int, aggregation *Transactions) (*Transactions, error)
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

func (svc *FvDataSvc) AllAccount(ctx context.Context, userId uuid.UUID) ([]byte, error) {
	accessToken, err := svc.getAccessToken(ctx, Finverse, userId)
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

func (svc *FvDataSvc) Income(ctx context.Context, userId uuid.UUID) ([]byte, error) {
	accessToken, err := svc.getAccessToken(ctx, Finverse, userId)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "GET", "/income", nil).
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

func (svc *FvDataSvc) PagingTransaction(ctx context.Context, offset string, limit string, userId uuid.UUID) ([]byte, error) {
	var queryStr = ""
	if offset != "" && limit != "" {
		queryStr = fmt.Sprintf("?offset=%s&limit=%s", offset, limit)
	} else if limit != "" {
		queryStr = fmt.Sprintf("?limit=%s", limit)
	} else if offset != "" {
		queryStr = fmt.Sprintf("?offset=%s", offset)
	}

	accessToken, err := svc.getAccessToken(ctx, Finverse, userId)
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
	accessToken, err := svc.getAccessToken(ctx, Finverse, userId)
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

	var aggregation []interface{}
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

		aggregation = append(aggregation, balance)
	}

	err = svc.provider.SaveAccount(ctx, userId, Finverse, aggregation)
	if err != nil {
		return nil, errorhandler.
			NewHTTPError(http.StatusInternalServerError).
			SetInternal(fmt.Errorf("failed to save fv accounts to sqlite: %w", err))
	}

	return aggregation, nil
}

func (svc *FvDataSvc) AggregateTransactions(ctx context.Context, userId uuid.UUID) (interface{}, error) {
	aggregation := &Transactions{}
	offset := 0
	limit := 10

	curTxs, err := svc.concatTransactions(ctx, userId, offset, limit, aggregation)
	if err != nil {
		return nil, err
	}

	totalTx := curTxs.TotalTransactions
	aggregation.TotalTransactions = totalTx
	aggregation.Other = curTxs.Other

	for {
		offset = offset + limit
		if totalTx-(offset+limit) < 0 {
			if totalTx > offset {
				svc.concatTransactions(ctx, userId, offset, limit, aggregation)
			}
			break
		}
		svc.concatTransactions(ctx, userId, offset, limit, aggregation)
	}

	err = svc.provider.SaveTransaction(ctx, userId, Finverse, aggregation)
	if err != nil {
		return nil, errorhandler.
			NewHTTPError(http.StatusInternalServerError).
			SetInternal(fmt.Errorf("failed to save fv transactions to sqlite: %w", err))
	}

	return aggregation, nil
}

func (svc *FvDataSvc) AggregateIncome(ctx context.Context, userId uuid.UUID) (interface{}, error) {
	var aggregation interface{}
	i, err := svc.Income(ctx, userId)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(i, &aggregation)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = svc.provider.SaveIncome(ctx, userId, Finverse, aggregation)
	if err != nil {
		return nil, errorhandler.
			NewHTTPError(http.StatusInternalServerError).
			SetInternal(fmt.Errorf("failed to save fv income to sqlite: %w", err))
	}

	return aggregation, nil
}

func (svc *FvDataSvc) concatTransactions(ctx context.Context, userId uuid.UUID, offset int, limit int, aggregation *Transactions) (*Transactions, error) {
	offsetStr := strconv.Itoa(offset)
	limitStr := strconv.Itoa(limit)
	txs, err := svc.PagingTransaction(ctx, offsetStr, limitStr, userId)
	if err != nil {
		return nil, err
	}
	var curTxs Transactions
	err = json.Unmarshal(txs, &curTxs)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	aggregation.Transactions = append(curTxs.Transactions, aggregation.Transactions...)
	return &curTxs, nil
}

func (svc *FvDataSvc) getAccessToken(ctx context.Context, providerName string, userId uuid.UUID) (*AccessToken, error) {
	var accessToken *AccessToken
	accessTokenPayload, err := svc.provider.ConnectionByProviderName(ctx, userId, providerName)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv connection: %w", err))
	}

	if accessTokenPayload == nil {
		return nil, errorhandler.NewHTTPError(http.StatusUnauthorized, "confidential connection not found")
	}

	err = json.Unmarshal([]byte(accessTokenPayload.Data), &accessToken)

	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return accessToken, nil
}
