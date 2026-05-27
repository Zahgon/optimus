package dependencyresolver

import (
	"context"

	pbp "github.com/raystack/optimus/protos/raystack/optimus/plugins/v1beta1"
	"github.com/raystack/optimus/sdk/plugin"
)

// GRPCServer will be used by plugins this is working as proto adapter
type GRPCServer struct {
	// This is the real implementation coming from plugin
	Impl plugin.DependencyResolverMod

	pbp.UnimplementedDependencyResolverModServiceServer
}

func (s *GRPCServer) GetName(ctx context.Context, _ *pbp.GetNameRequest) (*pbp.GetNameResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GRPCServer) GenerateDestination(ctx context.Context, req *pbp.GenerateDestinationRequest) (*pbp.GenerateDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GRPCServer) GenerateDependencies(ctx context.Context, req *pbp.GenerateDependenciesRequest) (*pbp.GenerateDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *GRPCServer) CompileAssets(ctx context.Context, req *pbp.CompileAssetsRequest) (*pbp.CompileAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
