package extension

import (
	"context"

	"github.com/raystack/optimus/client/extension/model"
)

// Manager defines the extension management
type Manager struct {
	manifester    model.Manifester
	assetOperator model.AssetOperator

	verbose              bool
	reservedCommandNames []string
}

// NewManager initializes new manager
func NewManager(
	manifester model.Manifester,
	assetOperator model.AssetOperator,
	verbose bool,
	reservedCommandNames ...string,
) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Activate runs an extension activation process
func (m *Manager) Activate(commandName, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Rename runs an extension renaming process
func (m *Manager) Rename(sourceCommandName, targetCommandName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run runs an extension execution process
func (m *Manager) Run(commandName string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Uninstall runs an extension uninstall process
func (m *Manager) Uninstall(commandName, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Upgrade runs an extension upgrade process
func (m *Manager) Upgrade(ctx context.Context, commandName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Install runs an extension installation process
func (m *Manager) Install(ctx context.Context, remotePath, commandName string) error {
	_ = "STUB: not implemented"
	return nil
}
