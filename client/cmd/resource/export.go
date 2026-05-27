package resource

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
)

const (
	fetchTenantTimeout   = time.Minute
	fetchResourceTimeout = time.Minute * 15
)

type exportCommand struct {
	logger     log.Logger
	connection connection.Connection

	writer local.SpecWriter[*model.ResourceSpec]

	configFilePath string
	outputDirPath  string
	host           string

	projectName   string
	namespaceName string
	resourceName  string

	storeName string
}

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

func (e *exportCommand) downloadSpecificResource(projectName, namespaceName, resourceName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *exportCommand) writeResources(projectName, namespaceName string, resources []*model.ResourceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *exportCommand) fetchAllResources(projectName, namespaceName string) ([]*model.ResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchSpecificResource(projectName, namespaceName, resourceName string) (*model.ResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchNamespaceNames(projectName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) fetchProjectNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exportCommand) validate() error { _ = "STUB: not implemented"; return nil }
