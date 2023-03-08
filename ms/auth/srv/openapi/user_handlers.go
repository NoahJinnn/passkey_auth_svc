package openapi

// func (srv *httpServer) GetUsers(params op.GetUsersParams) middleware.Responder {
// 	ctx, log := fromRequest(params.HTTPRequest)
// 	us, err := srv.app.GetAllUsers(ctx)

// 	resp := make([]*model.User, len(us))
// 	for _, u := range us {
// 		resp = append(resp, u.ToOAIResp())
// 	}

// 	switch {
// 	default:
// 		return errGetUsers(log, err, codeInternal)
// 	case err == nil:
// 		return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
// 			if err := producer.Produce(w, resp); err != nil {
// 				panic(err) // let the recovery middleware deal with this
// 			}
// 		})
// 	}
// }

// func (srv *httpServer) GetUserById(params op.GetUserByIDParams) middleware.Responder {
// 	ctx, log := fromRequest(params.HTTPRequest)
// 	u, err := srv.app.GetUserById(ctx, uint(params.UserID))

// 	switch {
// 	default:
// 		return errGetUserByID(log, err, codeInternal)
// 	case err == nil:
// 		return op.NewGetUserByIDOK().WithPayload(u.ToOAIResp())
// 	}
// }

// func (srv *httpServer) CreateUser(params op.CreateUserParams) middleware.Responder {
// 	ctx, log := fromRequest(params.HTTPRequest)
// 	domU := &dom.User{}
// 	domU = domU.FromOAIReq(params.User)
// 	u, err := srv.app.CreateUser(ctx, domU)

// 	switch {
// 	default:
// 		return errCreateUser(log, err, codeInternal)
// 	case err == nil:
// 		return op.NewCreateUserOK().WithPayload(u.ToOAIResp())
// 	}
// }
