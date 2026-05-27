package internal

import (
	"github.com/raystack/optimus/client/extension/model"
)

// RenameManager is an extension manater to manage command rename process
type RenameManager struct {
	manifester model.Manifester

	reservedCommandNames []string
	verbose              bool
}

// NewRenameManager initializes rename manager
func NewRenameManager(
	manifester model.Manifester,
	verbose bool,
	reservedCommandNames ...string,
) (*RenameManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rename renames an existing command name into a targeted command name
func (r *RenameManager) Rename(sourceCommandName, targetCommandName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenameManager) validateInput(sourceCommandName, targetCommandName string) error {
	_ = "STUB: not implemented"
	return nil
}
