package connection

import (
	"errors"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/raystack/salt/log"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/config"
)

const (
	grpcMaxClientSendSize      = 128 << 20 // 128MB
	grpcMaxClientRecvSize      = 128 << 20 // 128MB
	grpcMaxRetry          uint = 3

	optimusDialTimeout = time.Second * 2
	backoffDuration    = 100 * time.Millisecond
)

var errServerNotReachable = func(host string) error {
	return errors.New(heredoc.Docf(`Unable to reach optimus server at %s, this can happen due to following reasons:
		1. Check if you are connected to internet
		2. Is the host correctly configured in optimus config
		3. Is Optimus server currently unreachable`, host))
}

type Connection interface {
	Create(host string) (*grpc.ClientConn, error)
}

func New(l log.Logger, cfg *config.ClientConfig) Connection {
	_ = "STUB: not implemented"
	return *new(Connection)
}

func useInsecure() bool { _ = "STUB: not implemented"; return false }

func defaultDialOptions() []grpc.DialOption { _ = "STUB: not implemented"; return nil }
