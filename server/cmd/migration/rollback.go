package migration

import (
	"github.com/spf13/cobra"
)

type rollbackCommand struct {
	configFilePath string
	count          int
}

// NewRollbackCommand initializes command for migration rollback
func NewRollbackCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *rollbackCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// nolint:forbidigo

// nolint:forbidigo
