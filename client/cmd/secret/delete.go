package secret

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type deleteCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	projectName   string
	host          string
	namespaceName string
}

// NewDeleteCommand initializes command to delete secret
func NewDeleteCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (d *deleteCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (d *deleteCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (d *deleteCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *deleteCommand) deleteSecret(req *pb.DeleteSecretRequest) error {
	_ = "STUB: not implemented"
	return nil
}
