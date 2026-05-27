package internal

import (
	"context"

	"github.com/raystack/optimus/client/extension/model"
)

type upgradeResource struct {
	client         model.Client
	manifest       *model.Manifest
	metadata       *model.Metadata
	currentRelease *model.RepositoryRelease
	upgradeRelease *model.RepositoryRelease
}

// UpgradeManager is an extension manager to manage upgrade process
type UpgradeManager struct {
	manifester    model.Manifester
	assetOperator model.AssetOperator

	verbose bool
}

// NewUpgradeManager initializes upgrade manager
func NewUpgradeManager(
	manifester model.Manifester,
	assetOperator model.AssetOperator,
	verbose bool,
) (*UpgradeManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Upgrade upgrades extension specified by the command name
func (u *UpgradeManager) Upgrade(ctx context.Context, commandName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpgradeManager) rebuildManifest(resource *upgradeResource) *model.Manifest {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpgradeManager) upgradeOwnerWithResource(owner *model.RepositoryOwner, resource *upgradeResource) {
	_ = "STUB: not implemented"
	return
}

func (*UpgradeManager) upgradeProjectWithRelease(project *model.RepositoryProject, upgradeRelease *model.RepositoryRelease) {
	_ = "STUB: not implemented"
	return
}

func (u *UpgradeManager) setupResource(ctx context.Context, commandName string) (*upgradeResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*UpgradeManager) getCurrentRelease(project *model.RepositoryProject) *model.RepositoryRelease {
	_ = "STUB: not implemented"
	return nil
}

func (*UpgradeManager) validateInput(commandName string) error {
	_ = "STUB: not implemented"
	return nil
}
