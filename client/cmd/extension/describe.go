package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/extension/model"
)

type describeCommand struct {
	logger  log.Logger
	project *model.RepositoryProject
}

func newDescribeCommand(logger log.Logger, project *model.RepositoryProject) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func (d *describeCommand) RunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *describeCommand) buildHeader() string { _ = "STUB: not implemented"; return "" }

func (d *describeCommand) buildVerboseTable() string { _ = "STUB: not implemented"; return "" }

func (d *describeCommand) buildSimpleTable() string { _ = "STUB: not implemented"; return "" }
