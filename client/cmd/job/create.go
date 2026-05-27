package job

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/internal/models"
)

type createCommand struct {
	logger          log.Logger
	configFilePath  string
	clientConfig    *config.ClientConfig
	namespaceSurvey *survey.NamespaceSurvey
	jobCreateSurvey *survey.JobCreateSurvey
	pluginRepo      *models.PluginRepository
}

// NewCreateCommand initializes job create command
func NewCreateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Config filepath flag

func (c *createCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load mandatory config
	return nil
}

func (c *createCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*createCommand) PostRunE(*cobra.Command, []string) error {
	_ = "STUB: not implemented"
	return nil
}
