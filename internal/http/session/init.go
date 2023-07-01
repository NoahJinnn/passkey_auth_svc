package session

import (
	"context"
	"fmt"

	"github.com/hellohq/hqservice/internal/sharedconfig"
)

func InitSessionManager(ctxStartupCmdServe context.Context, cfg *sharedconfig.Shared, repo IJwkRepo) *Manager {
	jwkManager, err := NewDefaultManager(cfg.Secrets.Keys, repo)
	if err != nil {
		panic(fmt.Errorf("failed to create jwk manager: %w", err))
	}
	sessionManager, err := NewManager(jwkManager, cfg.Session)
	if err != nil {
		panic(fmt.Errorf("failed to create session generator: %w", err))
	}
	return sessionManager
}
