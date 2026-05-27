package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/extension/model"
)

type renameCommand struct {
	logger log.Logger

	project              *model.RepositoryProject
	reservedCommandNames []string
}

func newRenameCommand(logger log.Logger, project *model.RepositoryProject, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (r *renameCommand) RunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
