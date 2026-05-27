package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const jobStatusTimeout = time.Second * 30

type runListCommand struct {
	logger         log.Logger
	connection     *connection.Insecure
	configFilePath string

	startDate   string
	endDate     string
	projectName string
	host        string
}

// NewRunListCommand initializes run list command
func NewRunListCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *runListCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (r *runListCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (r *runListCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runListCommand) callJobRun(jobRunRequest *pb.JobRunRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *runListCommand) createJobRunRequest(jobName, startDate, endDate string) (*pb.JobRunRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*runListCommand) validateDateArgs(startDate, endDate string) error {
	_ = "STUB: not implemented"
	return nil
}
