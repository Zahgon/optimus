package replay

import (
	"bytes"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type statusCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	projectName string
	host        string
}

// StatusCommand get status for corresponding replay
func StatusCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *statusCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (r *statusCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *statusCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *statusCommand) getReplay(replayID string) (*pb.GetReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getReplay(host, replayID string, connection connection.Connection) (*pb.GetReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stringifyReplayStatus(resp *pb.GetReplayResponse) string { _ = "STUB: not implemented"; return "" }

func stringifyReplayConfig(buff *bytes.Buffer, jobConfig map[string]string) {
	_ = "STUB: not implemented"
	return
}

func stringifyReplayRuns(buff *bytes.Buffer, runs []*pb.ReplayRun) {
	_ = "STUB: not implemented"
	return
}
