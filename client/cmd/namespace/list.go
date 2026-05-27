package namespace

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/config"
)

const listTimeout = time.Minute * 15

type listCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string
	clientConfig   *config.ClientConfig

	dirPath     string
	host        string
	projectName string
}

// NewListCommand initializes command for listing namespace
func NewListCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (l *listCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (l *listCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Load config

func (l *listCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *listCommand) listNamespacesFromServer(serverHost, projectName string) ([]*config.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*listCommand) stringifyNamespaces(namespacesFromLocal, namespacesFromServer []*config.Namespace) string {
	_ = "STUB: not implemented"
	return ""
}
