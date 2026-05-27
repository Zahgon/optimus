package resource

import (
	"context"
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	uploadAllTimeout = time.Minute * 60
)

type uploadAllCommand struct {
	logger     log.Logger
	connection connection.Connection

	clientConfig *config.ClientConfig

	selectedNamespaceNames []string
	verbose                bool
	configFilePath         string

	batchSize int
}

// NewUploadAllCommand initializes command for uploading all resources
func NewUploadAllCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (u *uploadAllCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadAllCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadAllCommand) uploadAll(selectedNamespaces []*config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadAllCommand) uploadAllResources(ctx context.Context, conn *grpc.ClientConn, selectedNamespaces []*config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadAllCommand) sendNamespaceResourceRequest(stream pb.ResourceService_DeployResourceSpecificationClient,
	namespace *config.Namespace, progressFn func(totalCount int),
) error {
	_ = "STUB: not implemented"
	return nil
}

func readResourceSpecs(repoFS afero.Fs) ([]*model.ResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *uploadAllCommand) getResourceDeploymentRequest(namespaceName, storeName string,
	resources []*model.ResourceSpec,
) (*pb.DeployResourceSpecificationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *uploadAllCommand) getResourceStreamClient(ctx context.Context, conn *grpc.ClientConn) (pb.ResourceService_DeployResourceSpecificationClient, error) {
	_ = "STUB: not implemented"
	return *new(pb.ResourceService_DeployResourceSpecificationClient), nil
}

// TODO: create a new api for upload-all and remove deploy

func (u *uploadAllCommand) processResourceDeploymentResponse(stream pb.ResourceService_DeployResourceSpecificationClient) error {
	_ = "STUB: not implemented"
	return nil
}
