package job

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
)

const (
	fetchTenantTimeout = time.Minute
	fetchJobTimeout    = time.Minute * 15
)

type exportCommand struct {
	logger     log.Logger
	connection connection.Connection

	writer local.SpecWriter[*model.JobSpec]

	configFilePath string
	outputDirPath  string
	host           string

	projectName   string
	namespaceName string
	jobName       string
}

// NewExportCommand initializes command for exporting job specification to yaml file
func NewExportCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (e *exportCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *exportCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *exportCommand) downloadAll() bool { _ = "STUB: not implemented"; return false }

func (e *exportCommand) downloadByProjectName(projectName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *exportCommand) downloadByProjectNameAndNamespaceName(projectName, namespaceName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *exportCommand) downloadSpecificJob(projectName, namespaceName, jobName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *exportCommand) writeJobs(projectName, namespaceName string, jobs []*model.JobSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *exportCommand) fetchNamespaceJobsByProjectName(projectName string) (map[string][]*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchJobsByProjectAndNamespaceName(projectName, namespaceName string) ([]*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchSpecificJob(projectName, namespaceName, jobName string) (*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchProjectNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) validate() error { _ = "STUB: not implemented"; return nil }
