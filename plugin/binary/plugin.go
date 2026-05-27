package binary

import (
	"fmt"
	"runtime"

	"github.com/hashicorp/go-hclog"

	"github.com/raystack/optimus/internal/models"
)

var (
	Prefix = "optimus-"
	Suffix = fmt.Sprintf("_%s_%s", runtime.GOOS, runtime.GOARCH)
)

func Init(pluginsRepo *models.PluginRepository, discoveredBinaryPlugins []string, pluginLogger hclog.Logger, args ...string) error {
	_ = "STUB: not implemented"
	// pluginMap is the map of plugins we can dispense.
	return nil
}

// we are core, start by launching the plugin processes

// connect via GRPC

// create a client with dependency resolver mod
