package internal

import (
	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/internal/models"
)

// InitPlugins triggers initialization of all available plugins
func InitPlugins(logLevel config.LogLevel) (*models.PluginRepository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// discover and load plugins.

func CleanupPlugins() { _ = "STUB: not implemented"; return }
