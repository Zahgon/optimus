package plugin

import (
	getter "github.com/hashicorp/go-getter"
	"github.com/hashicorp/go-hclog"

	"github.com/raystack/optimus/config"
)

var (
	// for plugin installation, discovery & sync on both client and server
	PluginsDir         = ".plugins"
	PluginsArchiveName = "yaml-plugins.zip"
)

type IPluginManager interface {
	Install(dst string, sources ...string) error
	Archive(name string) error
	UnArchive(src, dest string) error
}

func NewPluginManager() *PluginManager { _ = "STUB: not implemented"; return nil }

type PluginManager struct {
	logger hclog.Logger
	client *getter.Client
}

func (p *PluginManager) installOne(dst, src string) error { _ = "STUB: not implemented"; return nil }

func (p *PluginManager) Install(dst string, sources ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func sanitizeArchivePath(d, t string) (v string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *PluginManager) UnArchive(src, dest string) error { _ = "STUB: not implemented"; return nil }

func (p *PluginManager) Archive(archiveName string) error { _ = "STUB: not implemented"; return nil }

// used during server start
// also exposed as cmd
func InstallPlugins(conf *config.ServerConfig) error { _ = "STUB: not implemented"; return nil }
