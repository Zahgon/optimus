package dependencyresolver

import (
	"context"
	"time"

	"github.com/hashicorp/go-hclog"

	pbp "github.com/raystack/optimus/protos/raystack/optimus/plugins/v1beta1"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	PluginGRPCMaxRetry = 3
	BackoffDuration    = 200 * time.Millisecond
)

// GRPCClient will be used by core to talk over grpc with plugins
type GRPCClient struct {
	client pbp.DependencyResolverModServiceClient
	logger hclog.Logger
}

func (m *GRPCClient) GetName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *GRPCClient) GenerateDestination(ctx context.Context, request plugin.GenerateDestinationRequest) (*plugin.GenerateDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *GRPCClient) GenerateDependencies(ctx context.Context, request plugin.GenerateDependenciesRequest) (*plugin.GenerateDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *GRPCClient) CompileAssets(ctx context.Context, request plugin.CompileAssetsRequest) (*plugin.CompileAssetsResponse, error) {
	_ = "STUB: not implemented" //nolint: gocritic
	return nil, nil
}

// propagateMetadata is based on UnaryClientInterceptor, here we cannot use interceptor as it is not
// available as a callOption for the grpc call. We need to manually inject the metadata to context
// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/google.golang.org/grpc/otelgrpc/interceptor.go#L67
func propagateMetadata(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (m *GRPCClient) makeFatalOnConnErr(err error) { _ = "STUB: not implemented"; return }
