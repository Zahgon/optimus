package backup

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/cmd/internal/survey"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	configTTL = "ttl"
)

type listCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	namespaceSurvey *survey.NamespaceSurvey

	projectName   string
	namespaceName string
	host          string
	storeName     string
}

// NewListCommand initialize command to list backup
func NewListCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (l *listCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (l *listCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

// use flag or ask namespace name

func (l *listCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*listCommand) stringifyBackupListResponse(listBackupsResponse *pb.ListBackupsResponse) string {
	_ = "STUB: not implemented"
	return ""
}
