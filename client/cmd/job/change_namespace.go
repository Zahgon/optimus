package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/config"
)

const (
	changeNamespaceTimeout = time.Minute * 1
)

type changeNamespaceCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string
	clientConfig   *config.ClientConfig

	project          string
	oldNamespaceName string
	newNamespaceName string
	host             string
}

// NewChangeNamespaceCommand initializes job namespace change command
func NewChangeNamespaceCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Config filepath flag

func (c *changeNamespaceCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Mandatory flags
	return
}

// Mandatory flags if config is not set

func (c *changeNamespaceCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load mandatory config
	return nil
}

func (c *changeNamespaceCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *changeNamespaceCommand) sendChangeNamespaceRequest(jobName string) error {
	_ = "STUB: not implemented"
	return nil
}

// fetch Instance by calling the optimus API

func (c *changeNamespaceCommand) PostRunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *changeNamespaceCommand) getNamespaceConfig(namespaceName string) (*config.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
