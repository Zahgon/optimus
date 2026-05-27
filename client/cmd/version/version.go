package version

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/connection"
	"github.com/raystack/optimus/internal/models"
)

const versionTimeout = time.Second * 2

type versionCommand struct {
	logger     log.Logger
	connection connection.Connection

	configFilePath string

	isWithServer bool
	host         string

	pluginRepo *models.PluginRepository
}

// NewVersionCommand initializes command to get version
func NewVersionCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (v *versionCommand) injectFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// Config filepath flag

// Mandatory flags if with-server is set but config is not set

func (v *versionCommand) PreRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *versionCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Print client version
	return nil
}

// Print server version

// Print version update if new version is exist

func (*versionCommand) PostRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *versionCommand) printAllPluginInfos() { _ = "STUB: not implemented"; return }

// getVersionRequest send a version request to service
func (v *versionCommand) getVersionRequest(clientVer, host string) (ver string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
