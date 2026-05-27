package github

import (
	"github.com/raystack/optimus/client/extension/model"
)

const apiPrefix = "https://api.github.com/repos"

// Release defines github repository release
type Release struct {
	TagName    string   `json:"tag_name"`
	Draft      bool     `json:"draft"`
	Prerelease bool     `json:"prerelease"`
	Assets     []*Asset `json:"assets"`
}

func (r *Release) toRepositoryRelease(apiPath string) *model.RepositoryRelease {
	_ = "STUB: not implemented"
	return nil
}

func (r *Release) getCurrentAPIPath(apiPath string) string { _ = "STUB: not implemented"; return "" }

func (*Release) getUpgradeAPIPath(apiPath string) string { _ = "STUB: not implemented"; return "" }

// Asset defines github release asset
type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func (a *Asset) toRepositoryAsset() *model.RepositoryAsset { _ = "STUB: not implemented"; return nil }
