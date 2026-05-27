package scheduler

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	uploadTimeout = time.Minute * 30
)

type uploadCommand struct {
	logger     log.Logger
	connection connection.Connection

	clientConfig *config.ClientConfig

	configFilePath string
}

// UploadCommand initializes command for scheduler DAG deployment
func UploadCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (u *uploadCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *uploadCommand) sendUploadAllRequest(projectName string) (*pb.UploadToSchedulerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
