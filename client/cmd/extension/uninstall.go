package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/client/extension/model"
)

type uninstallCommand struct {
	logger log.Logger
	survey *survey.ExtensionSurvey

	project              *model.RepositoryProject
	reservedCommandNames []string
}

func newUninstallCommand(logger log.Logger, project *model.RepositoryProject, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (r *uninstallCommand) RunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
