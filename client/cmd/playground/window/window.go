package window

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
)

type command struct {
	log log.Logger
}

// NewCommand initializes command for window playground
func NewCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (j *command) RunE(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
