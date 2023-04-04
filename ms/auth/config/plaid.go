package config

import "github.com/powerman/appcfg"

// // Ref: https://github.com/plaid/quickstart/blob/master/.env.example
type PlaidConfig struct {
	// 	// See https://dashboard.plaid.com/account/keys
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
