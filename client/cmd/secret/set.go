package secret

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/client/cmd/internal/survey"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type setCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	survey *survey.SecretSetSurvey

	projectName   string
	host          string
	namespaceName string
	filePath      string
	encoded       bool
	updateOnly    bool
	skipConfirm   bool
}

// NewSetCommand initializes command for setting secret
func NewSetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (s *setCommand) injectFlags(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Config filepath flag
	return
}

// Mandatory flags if config is not set

func (s *setCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Load config
	return nil
}

func (s *setCommand) RunE(_ *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *setCommand) registerSecret(req *pb.RegisterSecretRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *setCommand) updateSecret(req *pb.UpdateSecretRequest) error {
	_ = "STUB: not implemented"
	return nil
}
