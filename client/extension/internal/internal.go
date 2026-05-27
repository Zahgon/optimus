package internal

import (
	"context"

	"github.com/raystack/optimus/client/extension/model"
)

func buildOwner(metadata *model.Metadata, project *model.RepositoryProject) *model.RepositoryOwner {
	_ = "STUB: not implemented"
	return nil
}

func buildProject(metadata *model.Metadata, release *model.RepositoryRelease) *model.RepositoryProject {
	_ = "STUB: not implemented"
	return nil
}

func install(ctx context.Context, client model.Client, assetOperator model.AssetOperator, metadata *model.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func installAsset(assetOperator model.AssetOperator, asset []byte, localDirPath, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadAsset(ctx context.Context, client model.Client, currentAPIPath, upgradeAPIPath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isInstalled(manifest *model.Manifest, metadata *model.Metadata) bool {
	_ = "STUB: not implemented"
	return false
}

func isTagNameInProject(project *model.RepositoryProject, tagName string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateCommandNameOnReserved(commandName string, reservedCommandNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadRelease(ctx context.Context, client model.Client, currentAPIPath, upgradeAPIPath string) (*model.RepositoryRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findProjectByCommandName(manifest *model.Manifest, commandName string) *model.RepositoryProject {
	_ = "STUB: not implemented"
	return nil
}
