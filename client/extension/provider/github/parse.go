package github

import (
	"github.com/raystack/optimus/client/extension/factory"
	"github.com/raystack/optimus/client/extension/model"
)

// Parse parses remote path to get its metadata according to github convention
func Parse(remotePath string) (*model.Metadata, error) { _ = "STUB: not implemented"; return nil, nil }

func extractCommandName(repoName string) string { _ = "STUB: not implemented"; return "" }

func composeLocalDirPath(ownerName, repoName string) string { _ = "STUB: not implemented"; return "" }

func composeUpgradeAPIPath(ownerName, repoName string) string { _ = "STUB: not implemented"; return "" }

func composeCurrentAPIPath(ownerName, repoName, tagName string) string {
	_ = "STUB: not implemented"
	return ""
}

func extractTag(cleanedRemotePath string) string { _ = "STUB: not implemented"; return "" }

func extractRepoName(cleanedRemotePath string) string { _ = "STUB: not implemented"; return "" }

func extractOwner(cleanedRemotePath string) string { _ = "STUB: not implemented"; return "" }

func removeURLPrefix(remotePath string) string { _ = "STUB: not implemented"; return "" }

func validate(remotePath string) error { _ = "STUB: not implemented"; return nil }

func init() { //nolint:gochecknoinits
	factory.ParseRegistry = append(factory.ParseRegistry, Parse)
}
