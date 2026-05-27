package cmd

import (
	"github.com/spf13/cobra"
)

type serveCommand struct {
	configFilePath string
	installPlugins bool
}

// NewServeCommand initializes command to start server
func NewServeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (s *serveCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// TODO: find a way to load the config in one place
	return nil
}
