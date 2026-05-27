package replay

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	replayTimeout        = time.Minute * 1
	ISOTimeLayout        = time.RFC3339
	pollIntervalInSecond = 30
)

var (
	supportedISOTimeLayouts = [...]string{time.RFC3339, "2006-01-02"}
	terminalStatuses        = map[string]bool{"success": true, "failed": true, "invalid": true}
)

type createCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	parallel    bool
	description string
	jobConfig   string

	projectName   string
	namespaceName string
	host          string
}

// CreateCommand initializes command for creating a replay request
func CreateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

//nolint: gomnd

func (r *createCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (r *createCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *createCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint: gomnd

func (r *createCommand) waitForReplayState(replayID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *createCommand) getReplay(replayID string) (*pb.GetReplayResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *createCommand) createReplayRequest(jobName, startTimeStr, endTimeStr, jobConfig string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getTimeProto(timeStr string) (*timestamppb.Timestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
