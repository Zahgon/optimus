package yaml

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/raystack/optimus/internal/models"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	Prefix = "optimus-plugin-"
	Suffix = ".yaml"
)

type PluginSpec struct {
	plugin.Info                  `yaml:",inline,omitempty"`
	plugin.GetQuestionsResponse  `yaml:",inline,omitempty"`
	plugin.DefaultAssetsResponse `yaml:",inline,omitempty"`
	plugin.DefaultConfigResponse `yaml:",inline,omitempty"`
}

func (p *PluginSpec) PluginInfo() *plugin.Info { _ = "STUB: not implemented"; return nil }

func (p *PluginSpec) GetQuestions(context.Context, plugin.GetQuestionsRequest) (*plugin.GetQuestionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*PluginSpec) ValidateQuestion(_ context.Context, req plugin.ValidateQuestionRequest) (*plugin.ValidateQuestionResponse, error) {
	_ = "STUB: not implemented" //nolint
	return nil, nil
}

//nolint: nilerr

func (p *PluginSpec) DefaultConfig(_ context.Context, req plugin.DefaultConfigRequest) (*plugin.DefaultConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		// config from survey answers
		nil
}

// nolint:gocritic

// adding defaultconfig (static, macros & referential config) from yaml

func (p *PluginSpec) DefaultAssets(context.Context, plugin.DefaultAssetsRequest) (*plugin.DefaultAssetsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPluginSpec(pluginPath string) (*PluginSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default values

// standardize script value

// if error in loading, initializing or adding to pluginsrepo , skipping that particular plugin
// NOTE: binary plugins are loaded after yaml plugins loaded
func Init(pluginsRepo *models.PluginRepository, discoveredYamlPlugins []string, pluginLogger hclog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
