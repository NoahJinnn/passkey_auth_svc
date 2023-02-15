package apix

import (
	"context"
	"net"
	"path"

	"github.com/hellohq/hqservice/internal/dom"
	"github.com/hellohq/hqservice/pkg/def"
	"github.com/powerman/structlog"

	"github.com/powerman/go-monolith-example/pkg/grpcx"
)

const grpcGatewayIP = "127.0.0.1"

func isGRPCGateway(peerIP string) bool { return peerIP == grpcGatewayIP }

// GRPCNewContext returns a new context.Context that carries values describing
// this request without any deadline, plus result of authn.Authenticate.
func GRPCNewContext(ctx Ctx, fullMethod string, authn Authn) (_ Ctx, auth dom.Auth, err error) {
	remoteIP := net.ParseIP(grpcx.RemoteIP(ctx, isGRPCGateway)).String()
	if remoteIP != "<nil>" {
		ctx = context.WithValue(ctx, contextKeyRemoteIP, remoteIP)
		ctx = grpcx.AppendXFF(ctx, remoteIP)
		structlog.FromContext(ctx, nil).SetDefaultKeyvals(def.LogRemoteIP, remoteIP)
	}

	ctx = context.WithValue(ctx, contextKeyMethodName, path.Base(fullMethod))

	accessToken := AccessToken(grpcx.AccessToken(ctx))
	ctx = context.WithValue(ctx, contextKeyAccessToken, accessToken)

	if accessToken != "" {
		auth, err = authn.Authenticate(ctx, accessToken)
		if err == nil {
			ctx = context.WithValue(ctx, contextKeyAuth, auth)
			structlog.FromContext(ctx, nil).SetDefaultKeyvals(def.LogUserName, auth.UserName)
		}
	}

	return ctx, auth, err
}
