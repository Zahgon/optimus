package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/raystack/optimus/sdk/plugin"
)

type YamlMod struct {
	mock.Mock `hash:"-"`
}

func (repo *YamlMod) PluginInfo() *plugin.Info { _ = "STUB: not implemented"; return nil }

func (repo *YamlMod) DefaultConfig(ctx context.Context, inp plugin.DefaultConfigRequest) (*plugin.DefaultConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *YamlMod) DefaultAssets(ctx context.Context, inp plugin.DefaultAssetsRequest) (*plugin.DefaultAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *YamlMod) GetQuestions(ctx context.Context, inp plugin.GetQuestionsRequest) (*plugin.GetQuestionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *YamlMod) ValidateQuestion(ctx context.Context, inp plugin.ValidateQuestionRequest) (*plugin.ValidateQuestionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DependencyResolverMod struct {
	mock.Mock `hash:"-"`
}

func (repo *DependencyResolverMod) GetName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (repo *DependencyResolverMod) GenerateDestination(ctx context.Context, inp plugin.GenerateDestinationRequest) (*plugin.GenerateDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *DependencyResolverMod) GenerateDependencies(ctx context.Context, inp plugin.GenerateDependenciesRequest) (*plugin.GenerateDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo *DependencyResolverMod) CompileAssets(ctx context.Context, inp plugin.CompileAssetsRequest) (*plugin.CompileAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
