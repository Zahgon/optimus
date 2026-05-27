package secret

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type listCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	projectName string
	host        string
}

// NewListCommand initializes command for listing secret
func NewListCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (l *listCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (l *listCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (l *listCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *listCommand) listSecret(req *pb.ListSecretsRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (*listCommand) stringifyListOfSecrets(listSecretsResponse *pb.ListSecretsResponse) string {
	_ = "STUB: not implemented"
	return ""
}
