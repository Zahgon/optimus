package backup

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type createCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string
	isConfigExist  bool

	namespaceSurvey    *survey.NamespaceSurvey
	backupCreateSurvey *survey.BackupCreateSurvey

	projectName               string
	host                      string
	namespace                 string
	dsBackupConfig            string
	dsBackupConfigUnmarshaled map[string]string // unmarshaled version of datastoreConfig

	resourceNames []string
	description   string
	storeName     string
}

// NewCreateCommand initializes command to create backup
func NewCreateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (c *createCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (c *createCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (c *createCommand) fillAttributes(conf *config.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// use flag or ask namespace name

// use flag or fetched from config

func isStoreNameValid(name string, namespace *config.Namespace) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *createCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *createCommand) runBackupRequest() error { _ = "STUB: not implemented"; return nil }

func (c *createCommand) printBackupResponse(backupResponse *pb.CreateBackupResponse) {
	_ = "STUB: not implemented"
	return
}

func (c *createCommand) prepareInput() error { _ = "STUB: not implemented"; return nil }

func (c *createCommand) prepareDescription() error { _ = "STUB: not implemented"; return nil }

func (c *createCommand) prepareResourceNames() error { _ = "STUB: not implemented"; return nil }
