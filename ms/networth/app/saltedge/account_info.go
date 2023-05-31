package saltedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellohq/hqservice/ms/networth/config"
)

type Ctx = context.Context

type ISeAccountInfoSvc interface {
	Customer(ctx Ctx, customerId string) (*Customer, error)
	CreateCustomer(ctx Ctx, ccr *CreateCustomer) (*Customer, error)
	RemoveCustomer(ctx context.Context, customerId string) (*RemoveCustomer, error)
	CreateConnectSession(ctx Ctx, ccsr *CreateConnectSession) (*ConnectSession, error)
	GetConnectionByCustomerId(ctx Ctx, customerId string) (interface{}, error)
	GetAccountByConnectionId(ctx context.Context, connectionId string) (interface{}, error)
	GetTxByConnectionIdAndAccountId(ctx context.Context, connectionId string, accountId string) (interface{}, error)
}

type SeAccountInfoSvc struct {
	client *SeClient
}

func NewSeAccountInfoSvc(cfg *config.Config) *SeAccountInfoSvc {
	client := NewSeClient(cfg.SaltEdge)
	return &SeAccountInfoSvc{
		client: client,
	}
}

func (svc *SeAccountInfoSvc) Customer(ctx context.Context, customerId string) (*Customer, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result Customer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) CreateCustomer(ctx context.Context, ccr *CreateCustomer) (*Customer, error) {
	resp, err := svc.client.DoReq(ctx, http.MethodPost, "/customers", nil, ccr)
	if err != nil {
		return nil, err
	}

	var result Customer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) RemoveCustomer(ctx context.Context, customerId string) (*RemoveCustomer, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(ctx, http.MethodDelete, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result RemoveCustomer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) CreateConnectSession(ctx context.Context, ccsr *CreateConnectSession) (*ConnectSession, error) {
	resp, err := svc.client.DoReq(ctx, http.MethodPost, "/connect_sessions/create", nil, ccsr)
	if err != nil {
		return nil, err
	}

	var result ConnectSession
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) GetConnectionByCustomerId(ctx context.Context, customerId string) (interface{}, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) GetAccountByConnectionId(ctx context.Context, connectionId string) (interface{}, error) {
	resp, err := svc.client.DoReq(ctx, http.MethodGet, "/accounts", map[string][]string{
		"connection_id": {connectionId},
	}, nil)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc *SeAccountInfoSvc) GetTxByConnectionIdAndAccountId(ctx context.Context, connectionId string, accountId string) (interface{}, error) {
	resp, err := svc.client.DoReq(ctx, http.MethodGet, "/transactions", map[string][]string{
		"connection_id": {connectionId},
		"account_id":    {accountId},
	}, nil)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}
