package job

import (
	"context"
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	replaceAllTimeout = time.Minute * 60
)

type replaceAllCommand struct {
	logger     log.Logger
	connection connection.Connection

	clientConfig *config.ClientConfig

	selectedNamespaceNames []string
	verbose                bool
	configFilePath         string
}

// NewReplaceAllCommand initializes command for ReplaceAll
func NewReplaceAllCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *replaceAllCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *replaceAllCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *replaceAllCommand) replaceAll(selectedNamespaces []*config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *replaceAllCommand) replaceAllJobs(conn *grpc.ClientConn, selectedNamespaces []*config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *replaceAllCommand) sendNamespaceJobRequest(
	stream pb.JobSpecificationService_ReplaceAllJobSpecificationsClient,
	namespace *config.Namespace,
	progressFn func(totalCount int),
) error {
	_ = "STUB: not implemented"
	return nil
}

func (*replaceAllCommand) getReplaceAllRequest(projectName string, namespace *config.Namespace) (*pb.ReplaceAllJobSpecificationsRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *replaceAllCommand) getJobStreamClient(ctx context.Context, conn *grpc.ClientConn) (pb.JobSpecificationService_ReplaceAllJobSpecificationsClient, error) {
	_ = "STUB: not implemented"
	return *new(pb.JobSpecificationService_ReplaceAllJobSpecificationsClient), nil
}

func (r *replaceAllCommand) processJobReplaceAllResponses(stream pb.JobSpecificationService_ReplaceAllJobSpecificationsClient) error {
	_ = "STUB: not implemented"
	return nil
}
