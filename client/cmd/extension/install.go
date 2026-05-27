package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
)

type installCommand struct {
	logger log.Logger

	reservedCommandNames []string
}

func newInstallCommand(logger log.Logger, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (i *installCommand) RunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
