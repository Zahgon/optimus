package internal

import (
	"github.com/raystack/optimus/client/extension/model"
)

type uninstallResource struct {
	manifest *model.Manifest
	project  *model.RepositoryProject
	releases []*model.RepositoryRelease

	localDirPath string
	tagNames     []string
}

// UninstallManager is an extension manager to manage uninstallation process
type UninstallManager struct {
	manifester    model.Manifester
	assetOperator model.AssetOperator

	verbose bool
}

// NewUninstallManager initializes uninstall manager
func NewUninstallManager(
	manifester model.Manifester,
	assetOperator model.AssetOperator,
	verbose bool,
) (*UninstallManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Uninstall uninstalls extension based on the command name and the tag
func (u *UninstallManager) Uninstall(commandName, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *UninstallManager) rebuildManifest(resource *uninstallResource) *model.Manifest {
	_ = "STUB: not implemented"
	return nil
}

func (*UninstallManager) setReleasesForProject(project *model.RepositoryProject, releases []*model.RepositoryRelease) {
	_ = "STUB: not implemented"
	return
}

func (*UninstallManager) removeReleases(sourceReleases, releasesToBeRemoved []*model.RepositoryRelease) []*model.RepositoryRelease {
	_ = "STUB: not implemented"
	return nil
}

func (u *UninstallManager) uninstall(resource *uninstallResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *UninstallManager) setupResource(commandName, tagName string) (*uninstallResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*UninstallManager) findReleasesFromProject(project *model.RepositoryProject, tagName string) ([]*model.RepositoryRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*UninstallManager) validateInput(commandName, _ string) error {
	_ = "STUB: not implemented"
	return nil
}
