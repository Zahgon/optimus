package internal

import (
	"context"

	"github.com/raystack/optimus/client/extension/model"
)

type installResource struct {
	client   model.Client
	manifest *model.Manifest
	metadata *model.Metadata
	release  *model.RepositoryRelease
}

// InstallManager is an extension manager to manage installation process
type InstallManager struct {
	manifester    model.Manifester
	assetOperator model.AssetOperator

	verbose              bool
	reservedCommandNames []string
}

// NewInstallManager initializes install manager
func NewInstallManager(
	manifester model.Manifester,
	assetOperator model.AssetOperator,
	verbose bool,
	reservedCommandNames ...string,
) (*InstallManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Install installs extension
func (i *InstallManager) Install(ctx context.Context, remotePath, commandName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*InstallManager) rebuildManifest(resource *installResource) *model.Manifest {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallManager) validateResource(resource *installResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (*InstallManager) validateCommandNameOnManifest(manifest *model.Manifest, metadata *model.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallManager) setupInstallResource(ctx context.Context, remotePath, commandName string) (*installResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*InstallManager) extractMetadata(remotePath string) (*model.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*InstallManager) validateInput(remotePath string) error {
	_ = "STUB: not implemented"
	return nil
}
