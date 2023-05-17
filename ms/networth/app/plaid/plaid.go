package plaid

import (
	// 	"context"
	// 	"fmt"
	// 	"sort"
	// 	"strings"
	// 	"time"

	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hellohq/hqservice/ms/networth/config"
	plaid "github.com/plaid/plaid-go/v3/plaid"
)

type Ctx = context.Context

//nolint:gochecknoglobals // Config, flags and metrics are global anyway.
var (
	// We store the access_token in memory - in production, store it in a secure persistent data store.
	accessToken string
	itemID      string

	// The transfer_id is only relevant for the Transfer ACH product.
	// We store the transfer_id in memory - in production, store it in a secure
	// persistent data store
	transferID string

	environments = map[string]plaid.Environment{
		"sandbox":     plaid.Sandbox,
		"development": plaid.Development,
		"production":  plaid.Production,
	}
)

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

type plaidSvc struct {
	plaidClient *plaid.APIClient
	cfg         *config.Config
}

func NewPlaidClient(cfg *config.Config) *plaidSvc {
	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", cfg.Plaid.ClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", cfg.Plaid.Secret)
	configuration.UseEnvironment(environments[cfg.Plaid.Env])
	return &plaidSvc{
		plaidClient: plaid.NewAPIClient(configuration),
		cfg:         cfg,
	}
}

func (svc *plaidSvc) Info() *GetInfoResp {
	return &GetInfoResp{
		AccessToken: accessToken,
		ItemId:      itemID,
		Products:    svc.cfg.Plaid.Products,
	}
}

func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	countryCodes := []plaid.CountryCode{}

	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}

	return countryCodes
}

func convertProducts(productStrs []string) []plaid.Products {
	products := []plaid.Products{}

	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}

	return products
}

// linkTokenCreate creates a link token using the specified parameters
func (svc *plaidSvc) CreateLinkToken(
	ctx Ctx, paymentInitiation *plaid.LinkTokenCreateRequestPaymentInitiation,
) (*LinkTokenCreateResp, error) {
	countryCodes := convertCountryCodes(strings.Split(svc.cfg.Plaid.CountryCodes, ","))
	products := convertProducts(strings.Split(svc.cfg.Plaid.Products, ","))
	redirectURI := svc.cfg.Plaid.RedirectUri

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: time.Now().String(),
	}

	request := plaid.NewLinkTokenCreateRequest(
		"hellohq connector",
		"en",
		countryCodes,
		user,
	)

	request.SetProducts(products)

	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}

	if paymentInitiation != nil {
		request.SetPaymentInitiation(*paymentInitiation)
	}

	linkTokenCreateResp, _, err := svc.plaidClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		return nil, err
	}

	return &LinkTokenCreateResp{
		LinkToken: linkTokenCreateResp.GetLinkToken(),
	}, nil
}

// For sandbox testing
func (svc *plaidSvc) ExchangeSandboxAccessToken(ctx Ctx, institutionID string) (*GetAccessTokenResp, error) {
	products := convertProducts(strings.Split(svc.cfg.Plaid.Products, ","))
	// Create a one-time use public_token for the Item.
	// This public_token can be used to initialize Link in update mode for a user
	options := plaid.NewSandboxPublicTokenCreateRequestOptions()
	options.SetOverrideUsername("custom_noah")
	options.SetOverridePassword("123456")
	sbReq := *plaid.NewSandboxPublicTokenCreateRequest(
		institutionID,
		products,
	)
	sbReq.SetOptions(*options)

	sandboxPublicTokenResp, _, err := svc.plaidClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		sbReq,
	).Execute()

	if err != nil {
		return nil, err
	}

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := svc.plaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken()),
	).Execute()

	if err != nil {
		return nil, err
	}

	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()

	return &GetAccessTokenResp{
		AccessToken: accessToken,
		ItemId:      itemID,
	}, nil
}

func (svc *plaidSvc) ExchangeAccessToken(ctx Ctx, publicToken string) (*GetAccessTokenResp, error) {

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := svc.plaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()

	if err != nil {
		return nil, err
	}

	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()
	if itemExists(strings.Split(svc.cfg.Plaid.Products, ","), "transfer") {
		transferID, err = authorizeAndCreateTransfer(ctx, svc.plaidClient, accessToken)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("transfer id: " + transferID)
	fmt.Println("public token: " + publicToken)
	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)

	return &GetAccessTokenResp{
		AccessToken: accessToken,
		ItemId:      itemID,
	}, nil
}

func (svc *plaidSvc) GetAuthAccount(ctx Ctx) (*GetAuthAccountResp, error) {
	authGetResp, _, err := svc.plaidClient.PlaidApi.AuthGet(ctx).AuthGetRequest(
		*plaid.NewAuthGetRequest(accessToken),
	).Execute()

	if err != nil {
		return nil, err
	}

	return &GetAuthAccountResp{
		Accounts: authGetResp.GetAccounts(),
		Numbers:  authGetResp.GetNumbers(),
	}, nil
}

func (svc *plaidSvc) GetTransactions(Ctx) (*GetTransactionsResp, error) {
	ctx := context.Background()

	// Set cursor to empty to receive all historical updates
	var cursor *string

	// New transaction updates since "cursor"
	var added []plaid.Transaction
	var modified []plaid.Transaction
	var removed []plaid.RemovedTransaction // Removed transaction ids
	hasMore := true
	// Iterate through each page of new transaction updates for item
	for hasMore {
		request := plaid.NewTransactionsSyncRequest(accessToken)
		if cursor != nil {
			request.SetCursor(*cursor)
		}
		resp, _, err := svc.plaidClient.PlaidApi.TransactionsSync(
			ctx,
		).TransactionsSyncRequest(*request).Execute()
		if err != nil {
			return nil, err
		}

		// Add this page of results
		added = append(added, resp.GetAdded()...)
		_ = append(modified, resp.GetModified()...)
		_ = append(removed, resp.GetRemoved()...)
		hasMore = resp.GetHasMore()
		// Update cursor to the next cursor
		nextCursor := resp.GetNextCursor()
		cursor = &nextCursor
	}

	//TODO: Fix logx
	// logx.Infof("Added amounts: %v - Modified amounts: %v - Removed amounts: %v\n", len(added), len(modified), len(removed))

	sort.Slice(added, func(i, j int) bool {
		return added[i].GetDate() < added[j].GetDate()
	})
	latestTransactions := added[len(added)-9:] // TODO: This cause out of bound bug, need to be considered

	return &GetTransactionsResp{
		LatestTransactions: latestTransactions,
	}, nil
}

func (svc *plaidSvc) GetIdentity(Ctx) (*GetIdentityResp, error) {
	ctx := context.Background()

	identityGetResp, _, err := svc.plaidClient.PlaidApi.IdentityGet(ctx).IdentityGetRequest(
		*plaid.NewIdentityGetRequest(accessToken),
	).Execute()
	if err != nil {
		return nil, err
	}

	return &GetIdentityResp{
		Identity: identityGetResp.GetAccounts(),
	}, nil
}

func (svc *plaidSvc) GetBalance(Ctx) (*GetAccountsResp, error) {
	ctx := context.Background()

	balancesGetResp, _, err := svc.plaidClient.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
		*plaid.NewAccountsBalanceGetRequest(accessToken),
	).Execute()

	if err != nil {
		return nil, err
	}

	return &GetAccountsResp{
		Accounts: balancesGetResp.GetAccounts(),
	}, nil
}

func (svc *plaidSvc) GetAccounts(Ctx) (*GetAccountsResp, error) {
	ctx := context.Background()

	accountsGetResp, _, err := svc.plaidClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	if err != nil {
		return nil, err
	}

	return &GetAccountsResp{
		Accounts: accountsGetResp.GetAccounts(),
	}, nil
}

// Helper function to determine if Transfer is in Plaid product array
func itemExists(array []string, product string) bool {
	for _, item := range array {
		if item == product {
			return true
		}
	}

	return false
}

// This is a helper function to authorize and create a Transfer after successful
// exchange of a public_token for an access_token. The transfer_id is then used
// to obtain the data about that particular Transfer.
func authorizeAndCreateTransfer(ctx context.Context, client *plaid.APIClient, accessToken string) (string, error) {
	// We call /accounts/get to obtain first account_id - in production,
	// account_id's should be persisted in a data store and retrieved
	// from there.
	accountsGetResp, _, _ := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	accountID := accountsGetResp.GetAccounts()[0].AccountId

	transferAuthorizationCreateUser := plaid.NewTransferUserInRequest("FirstName LastName")
	transferAuthorizationCreateRequest := plaid.NewTransferAuthorizationCreateRequest(
		accessToken,
		accountID,
		"credit",
		"ach",
		"1.34",
		"ppd",
		*transferAuthorizationCreateUser,
	)
	transferAuthorizationCreateResp, _, err := client.PlaidApi.TransferAuthorizationCreate(ctx).TransferAuthorizationCreateRequest(*transferAuthorizationCreateRequest).Execute()
	if err != nil {
		return "", err
	}
	authorizationID := transferAuthorizationCreateResp.GetAuthorization().Id

	transferCreateRequest := plaid.NewTransferCreateRequest(
		accessToken,
		accountID,
		authorizationID,
		"credit",
		"ach",
		"1.34",
		"Payment",
		"ppd",
		*transferAuthorizationCreateUser,
	)
	transferCreateResp, _, err := client.PlaidApi.TransferCreate(ctx).TransferCreateRequest(*transferCreateRequest).Execute()
	if err != nil {
		return "", err
	}

	return transferCreateResp.GetTransfer().Id, nil
}
