package resource

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	applyTimeout  = time.Minute * 5
	successStatus = "success"
)

type applyCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string
	clientConfig   *config.ClientConfig

	namespaceSurvey *survey.NamespaceSurvey
	namespaceName   string
	projectName     string
	storeName       string

	verbose       bool
	resourceNames []string
}

// NewApplyCommand initializes command for applying resources from optimus to datastore
func NewApplyCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (a *applyCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *applyCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// use flag or ask namespace name

func (a *applyCommand) apply() error { _ = "STUB: not implemented"; return nil }

func (a *applyCommand) printApplyStatus(responses *pb.ApplyResourcesResponse) {
	_ = "STUB: not implemented"
	return
}
