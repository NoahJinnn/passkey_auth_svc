package finverse

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/pkg/httpx"
)

var (
	// We store the access_token in memory - in production, store it in a secure persistent data store.
	accessToken string
)

type IFvAuthSvc interface {
	CreateCustomerToken(ctx context.Context, cct *CreateCustomerToken) (*CustomerToken, error)
	CreateLinkToken(ctx context.Context, clt *CreateLinkToken) (*LinkToken, error)
	ExchangeAccessToken(ctx context.Context, exchangeCode string) (*AccessToken, error)
}

type authSvc struct {
	config *config.Config
	req    *httpx.Req
}

func NewFvAuthSvc(cfg *config.Config) IFvAuthSvc {
	req := httpx.NewReq("https://api.sandbox.finverse.net/")
	req.SetHeader("Content-Type", "application/json")

	return &authSvc{config: cfg, req: req}
}

func (svc *authSvc) CreateCustomerToken(ctx context.Context, cct *CreateCustomerToken) (*CustomerToken, error) {
	b, err := json.Marshal(cct)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.Post("auth/customer/token", b)
	fmt.Printf("resp: %+v\n", svc.req)
	if err != nil {
		return nil, err
	}

	var result CustomerToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	accessToken = result.AccessToken
	return &result, nil
}

func (svc *authSvc) CreateLinkToken(ctx context.Context, clt *CreateLinkToken) (*LinkToken, error) {
	b, err := json.Marshal(clt)
	if err != nil {
		return nil, err
	}

	svc.req.SetHeader("Authorization", "Bearer "+accessToken)
	resp, err := svc.req.Post("/link/token", b)
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

func (svc *authSvc) ExchangeAccessToken(ctx context.Context, exchangeCode string) (*AccessToken, error) {
	svc.req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	payload := fmt.Sprintf("client_id=%s&code=%s&redirect_uri=%s&grant_type=authorization_code", svc.config.Finverse.ClientID, exchangeCode, svc.config.Finverse.RedirectURI)

	resp, err := svc.req.Post("/auth/token", []byte(payload))
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
