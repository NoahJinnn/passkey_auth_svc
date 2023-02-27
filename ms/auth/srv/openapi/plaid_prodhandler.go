package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
)

func (srv *httpServer) GetAuthAccount(params op.GetAuthAccountParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	aa, err := srv.app.GetAuthAccount(ctx)

	switch {
	default:
		return errGetAuthAccount(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, aa); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}

func (srv *httpServer) GetTransactions(params op.GetTransactionsParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	txs, err := srv.app.GetTransactions(ctx)

	switch {
	default:
		return errGetTransactions(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, txs); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}

func (srv *httpServer) GetIdentity(params op.GetIdentityParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	id, err := srv.app.GetIdentity(ctx)

	switch {
	default:
		return errGetIdentity(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, id); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}

func (srv *httpServer) GetBalance(params op.GetBalanceParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	bl, err := srv.app.GetBalance(ctx)

	switch {
	default:
		return errGetBalance(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, bl); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}

func (srv *httpServer) GetAccounts(params op.GetAccountsParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	accs, err := srv.app.GetAccounts(ctx)

	switch {
	default:
		return errGetAccounts(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, accs); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}
