package connection

import (
	"github.com/raystack/salt/log"
	"google.golang.org/grpc"
)

type Insecure struct {
	l log.Logger
}

func NewInsecure(l log.Logger) *Insecure { _ = "STUB: not implemented"; return nil }

func (*Insecure) Create(host string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
