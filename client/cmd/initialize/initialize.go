package initialize

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
)

type initializeCommand struct {
	logger     log.Logger
	initSurvey *survey.InititalizeSurvey

	dirPath string
}

// NewInitializeCommand initializes command to interactively initialize client config
func NewInitializeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (i *initializeCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *initializeCommand) initClientConfig(clientConfig *config.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *initializeCommand) setupDirPathForClientConfig(clientConfig *config.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *initializeCommand) getClientConfigPath() string { _ = "STUB: not implemented"; return "" }
