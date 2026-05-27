package service

import (
	"context"
	"time"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	typeEnv = "env"
)

type FilesCompiler interface {
	Compile(fileMap map[string]string, context map[string]any) (map[string]string, error)
}

type PluginRepo interface {
	GetByName(name string) (*plugin.Plugin, error)
}

type JobRunAssetsCompiler struct {
	compiler   FilesCompiler
	pluginRepo PluginRepo

	logger log.Logger
}

func NewJobAssetsCompiler(engine FilesCompiler, pluginRepo PluginRepo, logger log.Logger) *JobRunAssetsCompiler {
	_ = "STUB: not implemented"
	return nil
}

func (c *JobRunAssetsCompiler) CompileJobRunAssets(ctx context.Context, job *scheduler.Job, systemEnvVars map[string]string, scheduledAt time.Time, contextForTask map[string]interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if task needs to override the compilation behaviour

// TODO: deprecate after changing type for plugin
func toJobRunSpecData(mapping map[string]string) []plugin.JobRunSpecData {
	_ = "STUB: not implemented"
	return nil
}

// TODO: deprecate
func toPluginAssets(assets map[string]string) plugin.Assets {
	_ = "STUB: not implemented"
	return *new(plugin.Assets)
}

// TODO: deprecate
func toPluginConfig(conf map[string]string) plugin.Configs {
	_ = "STUB: not implemented"
	return *new(plugin.Configs)
}
