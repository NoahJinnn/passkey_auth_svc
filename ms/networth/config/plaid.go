package config

// // Ref: https://github.com/plaid/quickstart/blob/master/.env.example
type PlaidConfig struct {
	// 	// See https://dashboard.plaid.com/account/keys
	ClientId string
	Secret   string
	// See sandbox, development, product
	Env string
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-products
	Products string
	// See https://plaid.com/docs/api/tokens/#link-token-create-request-country-codes
	CountryCodes string
	// See https://dashboard.plaid.com/team/api
	RedirectUri string
}
