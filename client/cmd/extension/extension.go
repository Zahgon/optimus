package extension

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/extension"
	"github.com/raystack/optimus/client/extension/model"
)

// UpdateWithExtension updates input command with the available extensions
func UpdateWithExtension(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func getReservedCommandNames(cmd *cobra.Command) []string { _ = "STUB: not implemented"; return nil }

func extensionCommand(logger log.Logger, reservedCommandNames []string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func generateManagementCommands(logger log.Logger, reservedCommandNames []string) []*cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func loadManifest() *model.Manifest { _ = "STUB: not implemented"; return nil }

func getExtensionManager(verbose bool, reservedCommandNames ...string) (*extension.Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
