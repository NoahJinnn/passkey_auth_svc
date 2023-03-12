//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE -imports=

// Package app provides business logic.
package app

import (
	"context"
	"errors"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/hellohq/hqservice/ms/auth/app/svcs"
	"github.com/hellohq/hqservice/ms/auth/config"
	"github.com/hellohq/hqservice/ms/auth/dal"
	plaid "github.com/plaid/plaid-go/v3/plaid"
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
	GetWebauthnSvc() svcs.IWebauthnSvc
}

// // Ref: https://github.com/plaid/quickstart/blob/master/.env.example
// // Config contains configuration for business-logic.
// type config struct {
// 	// See https://dashboard.plaid.com/account/keys
// 	ClientId appcfg.String `env:"PLAID_CLIENT_ID"`
// 	Secret   appcfg.String `env:"PLAID_SECRET"`
// 	// See sandbox, development, product
// 	Env appcfg.String `env:"PLAID_ENV"`
// 	// See https://plaid.com/docs/api/tokens/#link-token-create-request-products
// 	Products appcfg.String `env:"PLAID_PRODUCTS"`
// 	// See https://plaid.com/docs/api/tokens/#link-token-create-request-country-codes
// 	CountryCodes appcfg.String `env:"PLAID_COUNTRY_CODES"`
// 	// See https://dashboard.plaid.com/team/api
// 	RedirectUri appcfg.String `env:"PLAID_REDIRECT_URI"`
// }

// App implements interface Appl.
type App struct {
	cfg         *config.Config
	wAuthn      *webauthn.WebAuthn
	plaidClient *plaid.APIClient
	repo        dal.Repo
}

// New creates and returns new App.
func New(cfg *config.Config, repo dal.Repo) App {

	return App{
		cfg:  cfg,
		repo: repo,
	}
}

func (a App) GetWebauthnSvc() svcs.IWebauthnSvc {
	return svcs.NewWebAuthn(a.cfg, a.repo)
}
