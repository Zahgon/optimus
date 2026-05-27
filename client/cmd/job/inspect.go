package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	inspectTimeout         = time.Minute * 1
	optimusServerFetchFlag = "server"
	MASKED                 = "<masked>"
)

type inspectCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	projectName   string
	namespaceName string
	host          string

	clientConfig    *config.ClientConfig
	jobSurvey       *survey.JobSurvey
	namespaceSurvey *survey.NamespaceSurvey
}

// NewInspectCommand initializes command for inspecting job specification
func NewInspectCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (e *inspectCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load mandatory config
	return nil
}

func (e *inspectCommand) RunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *inspectCommand) getJobSpecByName(args []string, namespaceJobPath string) (*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *inspectCommand) loadConfig() error { _ = "STUB: not implemented"; return nil }

func (e *inspectCommand) inspectJobSpecification(jobSpec *model.JobSpec, serverFetch bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *inspectCommand) printLogs(logs []*pb.Log) { _ = "STUB: not implemented"; return }

func getRunsDateArray(jobRunProtos []*pb.JobRun) []string { _ = "STUB: not implemented"; return nil }

func (e *inspectCommand) displayUpstreamSection(upstreams *pb.JobInspectResponse_UpstreamSection) {
	_ = "STUB: not implemented"
	return
}

func (e *inspectCommand) displayDownstreamSection(downStreams *pb.JobInspectResponse_DownstreamSection) {
	_ = "STUB: not implemented"
	return
}

func (e *inspectCommand) displayBasicInfoSection(basicInfoSection *pb.JobInspectResponse_BasicInfoSection) {
	_ = "STUB: not implemented"
	return
}

func (e *inspectCommand) processJobInspectResponse(resp *pb.JobInspectResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *inspectCommand) yamlPrint(input interface{}) { _ = "STUB: not implemented"; return }
