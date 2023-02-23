//go:generate -command genny sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" genny
//go:generate genny -in=$GOFILE -out=gen.$GOFILE gen "HealthCheck=LinkTokenCreate,GetAccessToken,GetAccounts,GetAuthAccount,GetBalance,GetIdentity,GetSandboxAccessToken,GetTransactions,GetUsers,GetUserByID,CreateUser,UpdateUser"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/hellohq/hqservice/api/openapi/model"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
	"github.com/hellohq/hqservice/pkg/def"
)

func errHealthCheck(log Log, err error, code errCode) middleware.Responder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError { // Do no expose details about internal errors.
		msg = "internal error" //nolint:goconst // Duplicated by go:generate.
	}

	return op.NewHealthCheckDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(msg),
	})
}
