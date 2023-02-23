package openapi

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/hellohq/hqservice/api/openapi/restapi/op"
)

func (srv *httpServer) GetUsers(params op.GetUsersParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	us, err := srv.app.GetAllUsers(ctx)

	switch {
	default:
		return errGetUsers(log, err, codeInternal)
	case err == nil:
		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
			if err := producer.Produce(w, us); err != nil {
				panic(err) // let the recovery middleware deal with this
			}
		})
	}
}

func (srv *httpServer) GetUserById(params op.GetUserByIDParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	u, err := srv.app.GetUserById(ctx, uint(params.UserID))

	switch {
	default:
		return errGetUserByID(log, err, codeInternal)
	case err == nil:
		return op.NewGetUserByIDOK().WithPayload(u.ToOAIResp())
	}
}

func (srv *httpServer) CreateUser(params op.CreateUserParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	u, err := srv.app.CreateUser(ctx, params.User)

	switch {
	default:
		return errCreateUser(log, err, codeInternal)
	case err == nil:
		return op.NewCreateUserOK().WithPayload(u.ToOAIResp())
	}
}

func (srv *httpServer) UpdateUser(params op.UpdateUserParams) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest)
	u, err := srv.app.UpdateUser(ctx, params.User)

	switch {
	default:
		return errUpdateUser(log, err, codeInternal)
	case err == nil:
		return op.NewUpdateUserOK().WithPayload(u.ToOAIResp())
	}
}
