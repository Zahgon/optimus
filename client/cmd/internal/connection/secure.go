package connection

import (
	"time"

	"github.com/raystack/salt/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/raystack/optimus/config"
)

const authTimeout = time.Minute * 1

type Secure struct {
	l          log.Logger
	authConfig config.Auth
}

func NewSecure(l log.Logger, cfg *config.ClientConfig) *Secure {
	_ = "STUB: not implemented"
	return nil
}

func (s *Secure) Create(host string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Secure) getOptionsWithAuth() ([]grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setup https connection

// add the token for authentication

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}
