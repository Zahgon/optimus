package job

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/internal/models"
)

type addHookCommand struct {
	logger           log.Logger
	configFilePath   string
	clientConfig     *config.ClientConfig
	jobSurvey        *survey.JobSurvey
	jobAddHookSurvey *survey.JobAddHookSurvey
	namespaceSurvey  *survey.NamespaceSurvey
	pluginRepo       *models.PluginRepository
}

// NewAddHookCommand initializes command for adding hook
func NewAddHookCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Config filepath flag

func (a *addHookCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load mandatory config
	return nil
}

func (a *addHookCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*addHookCommand) PostRunE(*cobra.Command, []string) error {
	_ = "STUB: not implemented"
	return nil
}
