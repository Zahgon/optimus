package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
)

type cleanCommand struct {
	logger log.Logger
	survey *survey.ExtensionSurvey
}

func newCleanCommand(logger log.Logger) *cobra.Command { _ = "STUB: not implemented"; return nil }

func (c *cleanCommand) RunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
