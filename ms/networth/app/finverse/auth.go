package finverse

import (
	"context"
	"encoding/json"

	"github.com/hellohq/hqservice/ms/networth/config"
	"github.com/hellohq/hqservice/pkg/httpx"
)

var (
	// We store the access_token in memory - in production, store it in a secure persistent data store.
	accessToken string
)

type IFvAuthSvc interface {
	CreateCustomerToken(ctx context.Context) (*CustomerToken, error)
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

func (svc *authSvc) CreateCustomerToken(ctx context.Context) (*CustomerToken, error) {
	payload := CreateCustomerToken{
		ClientId:     svc.config.Finverse.ClientId,
		ClientSecret: svc.config.Finverse.Secret,
		GrantType:    "client_credentials",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.Post("auth/customer/token", b)
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

func (svc *authSvc) CreateLinkToken(ctx context.Context) (*CustomerToken, error) {
	payload := CreateCustomerToken{
		ClientId:     svc.config.Finverse.ClientId,
		ClientSecret: svc.config.Finverse.Secret,
		GrantType:    "client_credentials",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := svc.req.Post("/customer/token", b)
	if err != nil {
		return nil, err
	}

	var result CustomerToken
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
