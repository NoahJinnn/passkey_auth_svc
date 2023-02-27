package app

import (
	"github.com/hellohq/hqservice/api/openapi/model"
	plaid "github.com/plaid/plaid-go/v3/plaid"
)

type GetInfoResp struct {
	AccessToken string `json:"access_token"`
	ItemId      string `json:"item_id"`
	Products    string `json:"products"`
}

type GetAccessTokenReq struct {
	PublicToken string `form:"public_token"`
}
type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ItemId      string `json:"item_id"`
}

func (v *GetAccessTokenResp) ToOAIResp() *model.GetAccessTokenResp {
	return &model.GetAccessTokenResp{
		AccessToken: &v.AccessToken,
		ItemID:      &v.ItemId,
	}
}

type LinkTokenCreateResp struct {
	LinkToken string `json:"link_token"`
}

func (v *LinkTokenCreateResp) ToOAIResp() *model.LinkTokenCreateResp {
	return &model.LinkTokenCreateResp{
		LinkToken: &v.LinkToken,
	}
}

type GetAuthAccountResp struct {
	Accounts []plaid.AccountBase  `json:"accounts"`
	Numbers  plaid.AuthGetNumbers `json:"numbers"`
}

type GetTransactionsResp struct {
	LatestTransactions []plaid.Transaction `json:"latest_transactions"`
}

type GetIdentityResp struct {
	Identity []plaid.AccountIdentity `json:"identity"`
}

type GetAccountsResp struct {
	Accounts []plaid.AccountBase `json:"accounts"`
}

type GetSandboxAccessTokenReq struct {
	InstitutionID string `path:"institution_id"`
}
