package plugin

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/config"
)

type installCommand struct {
	logger         log.Logger
	serverConfig   *config.ServerConfig
	configFilePath string `default:"config.yaml"`
}

// NewInstallCommand initializes plugin install command
func NewInstallCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (i *installCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *installCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
