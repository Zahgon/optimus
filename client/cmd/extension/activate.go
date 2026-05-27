package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/extension/model"
)

type activateCommand struct {
	logger log.Logger

	project              *model.RepositoryProject
	reservedCommandNames []string
}

func newActivateCommand(logger log.Logger, project *model.RepositoryProject, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (a *activateCommand) RunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
