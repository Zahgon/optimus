package plugin

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
)

type validateCommand struct {
	logger log.Logger
	path   string
}

func NewValidateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (v *validateCommand) validateFile(pluginPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *validateCommand) validateDir(pluginPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *validateCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
