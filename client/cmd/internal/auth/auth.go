package auth

import (
	"context"

	"github.com/raystack/salt/log"
	"golang.org/x/oauth2"

	"github.com/raystack/optimus/config"
)

type Auth struct {
	logger log.Logger
	cfg    *oauth2.Config
}

func NewAuth(logger log.Logger, authConfig config.Auth) *Auth {
	_ = "STUB: not implemented"
	return nil
}

func (a Auth) GetToken(ctx context.Context) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (Auth) getTokenFromServer(ctx context.Context, cfg *oauth2.Config) (*oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toAuthConfig(authConfig config.Auth) *oauth2.Config { _ = "STUB: not implemented"; return nil }
