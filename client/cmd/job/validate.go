package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const validateTimeout = time.Minute * 5

type validateCommand struct {
	logger     log.Logger
	connection *connection.Insecure

	configFilePath string
	clientConfig   *config.ClientConfig

	verbose       bool
	namespaceName string
}

// NewValidateCommand initializes command for validating job specification
func NewValidateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Config filepath flag

func (v *validateCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented" // Load mandatory config
	return nil
}

func (v *validateCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *validateCommand) validateJobSpecificationRequest(jobSpecs []*model.JobSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *validateCommand) getCheckJobSpecificationsResponse(stream pb.JobSpecificationService_CheckJobSpecificationsClient) error {
	_ = "STUB: not implemented"
	return nil
}
