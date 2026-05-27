package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	refreshTimeout = time.Minute * 60
)

type refreshCommand struct {
	logger         log.Logger
	connection     connection.Connection
	configFilePath string

	verbose                bool
	selectedNamespaceNames []string
	selectedJobNames       []string

	projectName string
	host        string
}

// NewRefreshCommand initializes command for refreshing job specification
func NewRefreshCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *refreshCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (r *refreshCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (r *refreshCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *refreshCommand) refreshJobSpecificationRequest() error {
	_ = "STUB: not implemented"
	return nil
}

func (r *refreshCommand) handleRefreshResponse(stream pb.JobSpecificationService_RefreshJobsClient) error {
	_ = "STUB: not implemented"
	return nil
}
