//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/internal/sharedconfig"
	"github.com/hellohq/hqservice/ms/auth/app/dom"
	plaid "github.com/plaid/plaid-go/v3/plaid"
	"github.com/powerman/appcfg"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrAccessDenied  = errors.New("access denied")
	ErrAlreadyExist  = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrValidate      = errors.New("validate")
	ErrWrongPassword = errors.New("wrong password")
)

// Appl provides application features (use cases) service.
type Appl interface {
	// HealthCheck returns error if service is unhealthy or current
	// status otherwise.
	// Errors: none.
	HealthCheck(Ctx) (interface{}, error)
	IPlaidSvc
	IUserSvc
	IWebauthnSvc
}

type IPlaidSvc interface {
	Info() *GetInfoResp
	GetSandboxAccessToken(ctx Ctx, institutionID string) (*GetAccessTokenResp, error)
	LinkTokenCreate(
		ctx Ctx, paymentInitiation *plaid.LinkTokenCreateRequestPaymentInitiation,
	) (*LinkTokenCreateResp, error)
	GetAccessToken(ctx Ctx, publicToken string) (*GetAccessTokenResp, error)
	GetAuthAccount(ctx Ctx) (*GetAuthAccountResp, error)
	GetTransactions(ctx Ctx) (*GetTransactionsResp, error)
	GetIdentity(ctx Ctx) (*GetIdentityResp, error)
	GetBalance(ctx Ctx) (*GetAccountsResp, error)
	GetAccounts(ctx Ctx) (*GetAccountsResp, error)
}

type IUserSvc interface {
	GetAllUsers(ctx Ctx) ([]*dom.User, error)
	GetUserById(ctx Ctx, id uint) (*dom.User, error)
	CreateUser(ctx Ctx, u *dom.User) (*dom.User, error)
	UpdateUser(ctx Ctx, u *dom.User) (*dom.User, error)
}

type IWebauthnSvc interface {
	WebauthnBeginRegistration(ctx Ctx) (*protocol.CredentialCreation, *webauthn.SessionData, error)
}

// Repo provides data storage.
type Repo interface {
	GetAllUsers(ctx Ctx) ([]*dom.User, error)
	GetUserById(ctx Ctx, id uint) (*dom.User, error)
	CreateUser(ctx Ctx, u *dom.User) (*dom.User, error)
	UpdateUser(ctx Ctx, u *dom.User) (*dom.User, error)
}

// Ref: https://github.com/plaid/quickstart/blob/master/.env.example
// Config contains configuration for business-logic.
type config struct {
	// See https://dashboard.plaid.com/account/keys
	ClientId appcfg.String `env:"PLAID_CLIENT_ID"`
	Secret   appcfg.String `env:"PLAID_SECRET"`
	// See sandbox, development, product
	Env appcfg.String `env:"PLAID_ENV"`
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-products
	Products appcfg.String `env:"PLAID_PRODUCTS"`
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-country-codes
	CountryCodes appcfg.String `env:"PLAID_COUNTRY_CODES"`
	// See https://dashboard.plaid.com/team/api
	RedirectUri appcfg.String `env:"PLAID_REDIRECT_URI"`
}

// App implements interface Appl.
type App struct {
	cfg         *config
	wAuthn      *webauthn.WebAuthn
	plaidClient *plaid.APIClient
	repo        Repo
}

// New creates and returns new App.
func New(repo Repo) (*App, error) {
	var cfg = &config{}
	fromEnv := appcfg.NewFromEnv(sharedconfig.EnvPrefix)
	err := appcfg.ProvideStruct(cfg, fromEnv)

	if err != nil {
		return nil, fmt.Errorf("load app config failed: %w", err)
	}

	plaidClient := NewPlaidClient(*cfg)

	a := &App{
		cfg:         cfg,
		plaidClient: plaidClient,
		repo:        repo,
	}
	return a, nil
}

func (a *App) HealthCheck(_ Ctx) (interface{}, error) {
	return "OK", nil
}
