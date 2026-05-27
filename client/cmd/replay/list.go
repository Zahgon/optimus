package replay

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

func ListCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

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

func (l *listCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *listCommand) listReplay(req *pb.ListReplayRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func stringifyListOfReplays(resp *pb.ListReplayResponse) string {
	_ = "STUB: not implemented"
	return ""
}
