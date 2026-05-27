package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/extension/model"
)

type upgradeCommand struct {
	logger log.Logger

	project              *model.RepositoryProject
	reservedCommandNames []string
}

func newUpgradeCommand(logger log.Logger, project *model.RepositoryProject, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (u *upgradeCommand) RunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
