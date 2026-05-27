package plugin

import (
	"github.com/hashicorp/go-hclog"

	"github.com/raystack/optimus/internal/models"
)

func Initialize(pluginLogger hclog.Logger, arg ...string) (*models.PluginRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch yaml plugins first, it holds detailed information about the plugin

// fetch binary plugins. Any binary plugin which doesn't have its yaml version will be failed

// discoverPluginsGivenFilePattern look for plugin with the specific pattern in following folders
// order to search is top to down
// ./
// <exec>/
// <exec>/.optimus/plugins
// $HOME/.optimus/plugins
// /usr/bin
// /usr/local/bin
//
// for duplicate plugins(even with different versions for now), only the first found will be used
// sample plugin name:
// - optimus-myplugin_linux_amd64 | with suffix: optimus- and prefix: _linux_amd64
// - optimus-plugin-myplugin.yaml | with suffix: optimus-plugin and prefix: .yaml
func discoverPluginsGivenFilePattern(pluginLogger hclog.Logger, prefix, suffix string) []string {
	_ = "STUB: not implemented"
	return nil
}

// look in the same directory as the executable

// add user home directory

//nolint: gomnd

// get plugin name

// check for duplicate binaries, could be different versions
// if we have already discovered one, ignore rest

// Factory returns a new plugin instance
type Factory func(log hclog.Logger) interface{}

// Serve is used to serve a new Nomad plugin
func Serve(f Factory) { _ = "STUB: not implemented"; return }

func servePlugin(optimusPlugin interface{}, logger hclog.Logger) { _ = "STUB: not implemented"; return }
