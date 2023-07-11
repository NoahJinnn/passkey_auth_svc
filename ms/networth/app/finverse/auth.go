package finverse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/internal/http/errorhandler"
	"github.com/hellohq/hqservice/ms/networth/app/provider"
	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/ms/networth/dal"
	"github.com/hellohq/hqservice/pkg/httpx"
)

// We store the access_token in memory - in production, store it in a secure persistent data store.
var accessToken string

type IFvAuthSvc interface {
	CreateCustomerToken(ctx context.Context, cct *CreateCustomerToken) (*CustomerToken, error)
	CreateLinkToken(ctx context.Context, clt *CreateLinkToken) (*LinkToken, error)
	ExchangeAccessToken(ctx context.Context, exchangeCode string) (*AccessToken, error)
}

type FvAuthSvc struct {
	config   *config.Config
	req      *httpx.Req
	repo     dal.INwRepo
	provider *provider.ProviderSvc
}

func NewFvAuthSvc(cfg *config.Config, provider *provider.ProviderSvc, repo dal.INwRepo) *FvAuthSvc {
	req := httpx.NewReq("https://api.sandbox.finverse.net/", map[string]string{
		"Content-Type": "application/json",
	}, nil)

	return &FvAuthSvc{config: cfg, req: req, repo: repo}
}

func (svc *FvAuthSvc) CreateCustomerToken(ctx context.Context, cct *CreateCustomerToken, userId uuid.UUID) (bool, error) {
	b, err := json.Marshal(cct)
	if err != nil {
		return false, err
	}

	resp, err := svc.req.
		InitReq(ctx, "POST", "auth/customer/token", b).
		WithDefaultOpts().
		Send()
	if err != nil {
		return false, err
	}

	var result CustomerToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return false, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv session token: %w", err))
	}

	svc.repo.GetFvSessionRepo().Create(ctx, &ent.FvSession{
		UserID:      userId,
		AccessToken: result.AccessToken,
		ExpiresIn:   result.ExpiresIn,
		IssuedAt:    result.IssuedAt,
		TokenType:   result.TokenType,
	})
	return true, nil
}

func (svc *FvAuthSvc) CreateLinkToken(ctx context.Context, clt *CreateLinkToken, userId uuid.UUID) (*LinkToken, error) {
	b, err := json.Marshal(clt)
	if err != nil {
		return nil, err
	}

	fvSession, err := svc.repo.GetFvSessionRepo().GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "POST", "/link/token", b).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + fvSession.AccessToken,
		}).
		Send()

	if err != nil {
		return nil, err
	}

	var result LinkToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv link token: %w", err))
	}

	return &result, nil
}

func (svc *FvAuthSvc) ExchangeAccessToken(ctx context.Context, exchangeCode string, userId uuid.UUID) (*AccessToken, error) {
	fvSession, err := svc.repo.GetFvSessionRepo().GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	payload := fmt.Sprintf("client_id=%s&code=%s&redirect_uri=%s&grant_type=authorization_code", svc.config.Finverse.ClientID, exchangeCode, svc.config.Finverse.RedirectURI)
	resp, err := svc.req.
		InitReq(ctx, "POST", "/auth/token", []byte(payload)).
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + fvSession.AccessToken,
			"Content-Type":  "application/x-www-form-urlencoded",
		}).
		Send()
	if err != nil {
		return nil, err
	}

	var result AccessToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errorhandler.NewHTTPError(http.StatusInternalServerError).SetInternal(fmt.Errorf("failed to get fv exchange token: %w", err))
	}
	accessToken = result.AccessToken
	return &result, nil
}
