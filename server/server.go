package server

import (
	"net/http"
	"time"

	"github.com/raystack/salt/log"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/config"
)

const (
	shutdownWait       = 30 * time.Second
	GRPCMaxRecvMsgSize = 128 << 20 // 128MB
	GRPCMaxSendMsgSize = 128 << 20 // 128MB

	DialTimeout      = time.Second * 5
	BootstrapTimeout = time.Second * 10
)

func checkRequiredConfigs(conf config.Serve) error { _ = "STUB: not implemented"; return nil }

func setupGRPCServer(l log.Logger) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	return nil, nil
}

// Shared options for the logger, with a custom gRPC code to log level function.

// Make sure that log statements internal to gRPC library are logged using the logrus logger as well.

func prepareHTTPProxy(grpcAddr string, grpcServer *grpc.Server) (*http.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// prepare http proxy

// gRPC dialup options to proxy http connections

// base router

//nolint: gomnd

// FIXME: Creating issues for grpc connection

// grpcHandlerFunc routes http1 calls to baseMux and http2 with grpc header to grpcServer.
// Using a single port for proxying both http1 & 2 protocols will degrade http performance
// but for our use-case the convenience per performance tradeoff is better suited
// if in the future, this does become a bottleneck(which I highly doubt), we can break the service
// into two ports, default port for grpc and default+1 for grpc-gateway proxy.
// We can also use something like a connection multiplexer
// https://github.com/soheilhy/cmux to achieve the same.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
