package namespace

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/config"
)

const describeTimeout = time.Minute * 15

type describeCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	dirPath       string
	host          string
	projectName   string
	namespaceName string
}

// NewDescribeCommand initializes command to describe namespace
func NewDescribeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (d *describeCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (d *describeCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Load config

func (d *describeCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *describeCommand) getNamespace() (*config.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*describeCommand) stringifyNamespace(namespace *config.Namespace) string {
	_ = "STUB: not implemented"
	return ""
}
