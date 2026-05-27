package internal

import (
	"github.com/raystack/optimus/client/extension/model"
)

type runResource struct {
	localDirPath string
	tagName      string

	args []string
}

// RunManager is an extension manager to manage run operation
type RunManager struct {
	manifester    model.Manifester
	assetOperator model.AssetOperator

	verbose bool
}

// NewRunManager initializes run manager
func NewRunManager(
	manifester model.Manifester,
	assetOperator model.AssetOperator,
	verbose bool,
) (*RunManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run executes an installed extension
func (r *RunManager) Run(commandName string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RunManager) run(resource *runResource) error { _ = "STUB: not implemented"; return nil }

func (r *RunManager) setupResource(commandName string, args ...string) (*runResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*RunManager) validateInput(commandName string, _ ...string) error {
	_ = "STUB: not implemented"
	return nil
}
