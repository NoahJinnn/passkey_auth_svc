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

type seSvc struct {
	client *SeClient
}

func NewSeAccountInfoSvc(cfg *config.Config) ISeAccountInfoSvc {
	client := NewSeClient(cfg.SaltEdgeConfig)
	return &seSvc{
		client: client,
	}
}

func (svc *seSvc) Customer(ctx context.Context, customerId string) (*Customer, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(http.MethodGet, path, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result Customer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateCustomer(ctx context.Context, ccr *CreateCustomer) (*Customer, error) {
	resp, err := svc.client.DoReq(http.MethodPost, "/customers", nil, ccr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result Customer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) RemoveCustomer(ctx context.Context, customerId string) (*RemoveCustomer, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(http.MethodDelete, path, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result RemoveCustomer
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) CreateConnectSession(ctx context.Context, ccsr *CreateConnectSession) (*ConnectSession, error) {
	resp, err := svc.client.DoReq(http.MethodPost, "/connect_sessions/create", nil, ccsr)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result ConnectSession
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetConnectionByCustomerId(ctx context.Context, customerId string) (interface{}, error) {
	path := fmt.Sprintf("/customers/%s", customerId)

	resp, err := svc.client.DoReq(http.MethodGet, path, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetAccountByConnectionId(ctx context.Context, connectionId string) (interface{}, error) {
	resp, err := svc.client.DoReq(http.MethodGet, "/accounts", map[string][]string{
		"connection_id": {connectionId},
	}, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}

func (svc *seSvc) GetTxByConnectionIdAndAccountId(ctx context.Context, connectionId string, accountId string) (interface{}, error) {
	resp, err := svc.client.DoReq(http.MethodGet, "/transactions", map[string][]string{
		"connection_id": {connectionId},
		"account_id":    {accountId},
	}, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(resp, &HttpBody{
		Data: &result,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &result, nil
}
