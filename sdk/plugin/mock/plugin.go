package mock

import (
	"context"

	"github.com/raystack/optimus/sdk/plugin"
)

func NewMockBinaryPlugin(name, pluginType string) *plugin.Plugin {
	_ = "STUB: not implemented"
	return nil
}

func NewMockYamlPlugin(name, pluginType string) *plugin.Plugin {
	_ = "STUB: not implemented"
	return nil
}

type MockYamlMod struct {
	Name string
	Type string
}

func (p *MockYamlMod) PluginInfo() *plugin.Info { _ = "STUB: not implemented"; return nil }

func (*MockYamlMod) GetQuestions(context.Context, plugin.GetQuestionsRequest) (*plugin.GetQuestionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*MockYamlMod) ValidateQuestion(context.Context, plugin.ValidateQuestionRequest) (*plugin.ValidateQuestionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*MockYamlMod) DefaultConfig(context.Context, plugin.DefaultConfigRequest) (*plugin.DefaultConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*MockYamlMod) DefaultAssets(context.Context, plugin.DefaultAssetsRequest) (*plugin.DefaultAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MockDependencyMod struct {
	Name string
	Type string
}

func (*MockDependencyMod) GetName(context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*MockDependencyMod) GenerateDestination(context.Context, plugin.GenerateDestinationRequest) (*plugin.GenerateDestinationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*MockDependencyMod) GenerateDependencies(context.Context, plugin.GenerateDependenciesRequest) (*plugin.GenerateDependenciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*MockDependencyMod) CompileAssets(context.Context, plugin.CompileAssetsRequest) (*plugin.CompileAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
