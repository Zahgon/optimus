package backup

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	backupTimeout = time.Minute * 15
)

// NewBackupCommand initializes
func NewBackupCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }
