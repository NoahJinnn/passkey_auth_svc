// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateUserParams, opts ...ClientOption) (*CreateUserOK, error)

	GetAccessToken(params *GetAccessTokenParams, opts ...ClientOption) (*GetAccessTokenOK, error)

	GetAccounts(params *GetAccountsParams, opts ...ClientOption) (*GetAccountsOK, error)

	GetAuthAccount(params *GetAuthAccountParams, opts ...ClientOption) (*GetAuthAccountOK, error)

	GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error)

	GetIdentity(params *GetIdentityParams, opts ...ClientOption) (*GetIdentityOK, error)

	GetInfo(params *GetInfoParams, opts ...ClientOption) (*GetInfoOK, error)

	GetSandboxAccessToken(params *GetSandboxAccessTokenParams, opts ...ClientOption) (*GetSandboxAccessTokenOK, error)

	GetTransactions(params *GetTransactionsParams, opts ...ClientOption) (*GetTransactionsOK, error)

	GetUserByID(params *GetUserByIDParams, opts ...ClientOption) (*GetUserByIDOK, error)

	GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error)

	LinkTokenCreate(params *LinkTokenCreateParams, opts ...ClientOption) (*LinkTokenCreateOK, error)

	UpdateUser(params *UpdateUserParams, opts ...ClientOption) (*UpdateUserOK, error)

	HealthCheck(params *HealthCheckParams, opts ...ClientOption) (*HealthCheckOK, error)

	WebauthnLoginInit(params *WebauthnLoginInitParams, opts ...ClientOption) (*WebauthnLoginInitOK, error)

	WebauthnRegFinal(params *WebauthnRegFinalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebauthnRegFinalOK, error)

	WebauthnRegInit(params *WebauthnRegInitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebauthnRegInitOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Create create user API
*/
func (a *Client) Create(params *CreateUserParams, opts ...ClientOption) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAccessToken get access token API
*/
func (a *Client) GetAccessToken(params *GetAccessTokenParams, opts ...ClientOption) (*GetAccessTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccessToken",
		Method:             "POST",
		PathPattern:        "/set_access_token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccessTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAccessTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAccounts get accounts API
*/
func (a *Client) GetAccounts(params *GetAccountsParams, opts ...ClientOption) (*GetAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAccounts",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAuthAccount get auth account API
*/
func (a *Client) GetAuthAccount(params *GetAuthAccountParams, opts ...ClientOption) (*GetAuthAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuthAccount",
		Method:             "GET",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuthAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAuthAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBalance get balance API
*/
func (a *Client) GetBalance(params *GetBalanceParams, opts ...ClientOption) (*GetBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBalanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBalance",
		Method:             "GET",
		PathPattern:        "/balance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBalanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBalanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIdentity get identity API
*/
func (a *Client) GetIdentity(params *GetIdentityParams, opts ...ClientOption) (*GetIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIdentity",
		Method:             "GET",
		PathPattern:        "/identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIdentityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIdentityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetInfo get info API
*/
func (a *Client) GetInfo(params *GetInfoParams, opts ...ClientOption) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInfo",
		Method:             "POST",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSandboxAccessToken get sandbox access token API
*/
func (a *Client) GetSandboxAccessToken(params *GetSandboxAccessTokenParams, opts ...ClientOption) (*GetSandboxAccessTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSandboxAccessTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSandboxAccessToken",
		Method:             "GET",
		PathPattern:        "/sandbox_access_token/{institution_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSandboxAccessTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSandboxAccessTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSandboxAccessTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTransactions get transactions API
*/
func (a *Client) GetTransactions(params *GetTransactionsParams, opts ...ClientOption) (*GetTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTransactions",
		Method:             "GET",
		PathPattern:        "/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTransactionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUserByID get user by ID API
*/
func (a *Client) GetUserByID(params *GetUserByIDParams, opts ...ClientOption) (*GetUserByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserByID",
		Method:             "GET",
		PathPattern:        "/user/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUserByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUsers get users API
*/
func (a *Client) GetUsers(params *GetUsersParams, opts ...ClientOption) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
LinkTokenCreate link token create API
*/
func (a *Client) LinkTokenCreate(params *LinkTokenCreateParams, opts ...ClientOption) (*LinkTokenCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLinkTokenCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LinkTokenCreate",
		Method:             "POST",
		PathPattern:        "/create_link_token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LinkTokenCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LinkTokenCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LinkTokenCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateUser update user API
*/
func (a *Client) UpdateUser(params *UpdateUserParams, opts ...ClientOption) (*UpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateUser",
		Method:             "PATCH",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HealthCheck Returns 200 if service works okay.
*/
func (a *Client) HealthCheck(params *HealthCheckParams, opts ...ClientOption) (*HealthCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "healthCheck",
		Method:             "GET",
		PathPattern:        "/health-check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HealthCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HealthCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	WebauthnLoginInit initializes web authn login

	Initialize a login with Webauthn. Returns a JSON representation of CredentialRequestOptions for use

with the Webauthn API's `navigator.credentials.get()`.

Omitting the optional request body or using an empty JSON object results in generation of request options for a
login using a [discoverable credential](https://www.w3.org/TR/webauthn-2/#client-side-discoverable-public-key-credential-source)
(i.e. they will not contain
[allowCredentials](https://www.w3.org/TR/webauthn-2/#dom-publickeycredentialrequestoptions-allowcredentials)).
*/
func (a *Client) WebauthnLoginInit(params *WebauthnLoginInitParams, opts ...ClientOption) (*WebauthnLoginInitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebauthnLoginInitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "webauthnLoginInit",
		Method:             "POST",
		PathPattern:        "/webauthn/login/initialize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebauthnLoginInitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WebauthnLoginInitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for webauthnLoginInit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WebauthnRegFinal finalizes web authn registration

Finalize a registration with Webauthn using the WebAuthn API response to a `navigator.credentials.create()` call.
*/
func (a *Client) WebauthnRegFinal(params *WebauthnRegFinalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebauthnRegFinalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebauthnRegFinalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "webauthnRegFinal",
		Method:             "POST",
		PathPattern:        "/webauthn/registration/finalize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebauthnRegFinalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WebauthnRegFinalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for webauthnRegFinal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WebauthnRegInit initializes web authn registration

	Initialize a registration with Webauthn. Returns a JSON representation of CredentialCreationOptions for use

with the Webauthn API's `navigator.credentials.create()`.
*/
func (a *Client) WebauthnRegInit(params *WebauthnRegInitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WebauthnRegInitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebauthnRegInitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "webauthnRegInit",
		Method:             "POST",
		PathPattern:        "/webauthn/registration/initialize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WebauthnRegInitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WebauthnRegInitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for webauthnRegInit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
