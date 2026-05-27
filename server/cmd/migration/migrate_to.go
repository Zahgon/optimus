package migration

import (
	"github.com/spf13/cobra"
)

type migrateTo struct {
	configFilePath string
	version        int
}

// NewMigrateToCommand initializes command for migration to a specific version
func NewMigrateToCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (m *migrateTo) RunE(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// nolint:forbidigo

// nolint:forbidigo
