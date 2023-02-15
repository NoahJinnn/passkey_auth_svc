package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/model"
	"github.com/hellohq/hqservice/ms/hq/app"
)

type CustomResponder func(http.ResponseWriter, runtime.Producer)

func (c CustomResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}

func NewCustomResponder(r *http.Request, h http.Handler) middleware.Responder {
	return CustomResponder(func(w http.ResponseWriter, _ runtime.Producer) {
		h.ServeHTTP(w, r)
	})
}

func apiGetSandboxAccessToken(v *app.GetAccessTokenResp) *model.GetAccessTokenResp {
	return &model.GetAccessTokenResp{
		AccessToken: &v.AccessToken,
		ItemID:      &v.ItemId,
	}
}

func apiLinkTokenCreate(v *app.LinkTokenCreateResp) *model.LinkTokenCreateResp {
	return &model.LinkTokenCreateResp{
		LinkToken: &v.LinkToken,
	}
}

func apiGetAccessToken(v *app.GetAccessTokenResp) *model.GetAccessTokenResp {
	return &model.GetAccessTokenResp{
		AccessToken: &v.AccessToken,
		ItemID:      &v.ItemId,
	}
}
