package models

import (
	"errors"

	"github.com/raystack/optimus/sdk/plugin"
)

var ErrUnsupportedPlugin = errors.New("unsupported plugin requested, make sure its correctly installed")

type PluginRepository struct {
	data       map[string]*plugin.Plugin
	sortedKeys []string
}

func (s *PluginRepository) lazySortPluginKeys() {
	_ = "STUB: not implemented"
	// already sorted
	return
}

func (s *PluginRepository) GetByName(name string) (*plugin.Plugin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PluginRepository) GetAll() []*plugin.Plugin { _ = "STUB: not implemented"; return nil }

// sorts keys if not sorted

func (s *PluginRepository) GetTasks() []*plugin.Plugin { _ = "STUB: not implemented"; return nil }

// sorts keys if not sorted

func (s *PluginRepository) GetHooks() []*plugin.Plugin { _ = "STUB: not implemented"; return nil }

func (s *PluginRepository) AddYaml(yamlMod plugin.YamlMod) error {
	_ = "STUB: not implemented"
	return nil
}

// duplicated yaml plugin

func (s *PluginRepository) AddBinary(drMod plugin.DependencyResolverMod) error {
	_ = "STUB: not implemented"
	return nil
}

// any binary plugin should have its yaml version (for the plugin information)

// duplicated binary plugin

func NewPluginRepository() *PluginRepository { _ = "STUB: not implemented"; return nil }
