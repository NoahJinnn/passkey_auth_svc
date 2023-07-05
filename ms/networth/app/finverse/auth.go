package finverse

import (
	"context"
	"encoding/json"
	"fmt"

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
	config *config.Config
	req    *httpx.Req
	repo   dal.INwRepo
}

func NewFvAuthSvc(cfg *config.Config) *FvAuthSvc {
	req := httpx.NewReq("https://api.sandbox.finverse.net/", map[string]string{
		"Content-Type": "application/json",
	}, nil)

	return &FvAuthSvc{config: cfg, req: req}
}

func (svc *FvAuthSvc) CreateCustomerToken(ctx context.Context, cct *CreateCustomerToken, userId string) (*CustomerToken, error) {
	b, err := json.Marshal(cct)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "POST", "auth/customer/token", b).
		WithDefaultOpts().
		Send()
	if err != nil {
		return nil, err
	}

	var result CustomerToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	accessToken = result.AccessToken
	// TODO: Save finverse session to db
	// svc.repo.GetFvSessionRepo()
	return &result, nil
}

func (svc *FvAuthSvc) CreateLinkToken(ctx context.Context, clt *CreateLinkToken) (*LinkToken, error) {
	b, err := json.Marshal(clt)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.
		InitReq(ctx, "POST", "/link/token", b).
		WithDefaultOpts().
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken,
		}).
		Send()

	if err != nil {
		return nil, err
	}

	var result LinkToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *FvAuthSvc) ExchangeAccessToken(ctx context.Context, exchangeCode string) (*AccessToken, error) {
	payload := fmt.Sprintf("client_id=%s&code=%s&redirect_uri=%s&grant_type=authorization_code", svc.config.Finverse.ClientID, exchangeCode, svc.config.Finverse.RedirectURI)
	resp, err := svc.req.
		InitReq(ctx, "POST", "/auth/token", []byte(payload)).
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + accessToken,
			"Content-Type":  "application/x-www-form-urlencoded",
		}).
		Send()
	if err != nil {
		return nil, err
	}

	var result AccessToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
