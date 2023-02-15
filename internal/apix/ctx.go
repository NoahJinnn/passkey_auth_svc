package apix

import (
	"context"

	"github.com/hellohq/hqservice/internal/dom"
)

type Ctx = context.Context

type contextKey int

const (
	_ contextKey = iota
	contextKeyRemoteIP
	contextKeyMethodName
	contextKeyAccessToken
	contextKeyAuth
)

// NewContextWithRemoteIP creates and returns new context containing given
// remoteIP.
func NewContextWithRemoteIP(ctx Ctx, remoteIP string) Ctx {
	return context.WithValue(ctx, contextKeyRemoteIP, remoteIP)
}

// FromContext returns values describing request stored in ctx, if any.
func FromContext(ctx Ctx) (remoteIP, methodName string, auth dom.Auth) {
	remoteIP, _ = ctx.Value(contextKeyRemoteIP).(string)
	methodName, _ = ctx.Value(contextKeyMethodName).(string)
	auth, _ = ctx.Value(contextKeyAuth).(dom.Auth)
	return remoteIP, methodName, auth
}

// AccessTokenFromContext returns AccessToken stored in ctx, if any.
func AccessTokenFromContext(ctx Ctx) (accessToken AccessToken) {
	accessToken, _ = ctx.Value(contextKeyAccessToken).(AccessToken)
	return
}
