package setup

import (
	"context"

	"github.com/raystack/optimus/sdk/plugin"
)

type MockPluginBQ struct{}

func (MockPluginBQ) GetName(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (MockPluginBQ) GenerateDestination(_ context.Context, request plugin.GenerateDestinationRequest) (*plugin.GenerateDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (MockPluginBQ) GenerateDependencies(_ context.Context, req plugin.GenerateDependenciesRequest) (*plugin.GenerateDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (MockPluginBQ) CompileAssets(_ context.Context, _ plugin.CompileAssetsRequest) (*plugin.CompileAssetsResponse, error) {
	_ = "STUB: not implemented"
	// TODO: implement mock
	return nil, nil
}
