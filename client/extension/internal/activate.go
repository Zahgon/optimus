package internal

import (
	"github.com/raystack/optimus/client/extension/model"
)

// ActivateManager is an extension manater to manage tag activation process
type ActivateManager struct {
	manifester model.Manifester

	verbose bool
}

// NewActivateManager initializes activate manager
func NewActivateManager(manifester model.Manifester, verbose bool) (*ActivateManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Activate activates the tag for an extension specified by the command name
func (a *ActivateManager) Activate(commandName, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*ActivateManager) activateTagInProject(project *model.RepositoryProject, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*ActivateManager) validateActivateInput(commandName, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}
