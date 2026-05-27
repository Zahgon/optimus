package service

import (
	"errors"
	"time"

	"github.com/raystack/salt/log"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	projectConfigPrefix = "GLOBAL__"

	configKeyDstart        = "DSTART"
	configKeyDend          = "DEND"
	configKeyExecutionTime = "EXECUTION_TIME"
	configKeyDestination   = "JOB_DESTINATION"

	TimeISOFormat = time.RFC3339
)

var (
	ErrUpstreamModNotFound = errors.New("upstream mod not found for plugin")
	ErrYamlModNotExist     = errors.New("yaml mod not found for plugin")
)

type PluginRepo interface {
	GetByName(string) (*plugin.Plugin, error)
}

type Engine interface {
	Compile(templateMap map[string]string, context map[string]any) (map[string]string, error)
	CompileString(input string, context map[string]any) (string, error)
}

type JobPluginService struct {
	pluginRepo PluginRepo
	engine     Engine

	now func() time.Time

	logger log.Logger
}

func NewJobPluginService(pluginRepo PluginRepo, engine Engine, logger log.Logger) *JobPluginService {
	_ = "STUB: not implemented"
	return nil
}

func (p JobPluginService) Info(_ context.Context, taskName job.TaskName) (*plugin.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p JobPluginService) GenerateDestination(ctx context.Context, tnnt *tenant.WithDetails, task job.Task) (job.ResourceURN, error) {
	_ = "STUB: not implemented"
	return *new(job.ResourceURN), nil
}

func (p JobPluginService) GenerateUpstreams(ctx context.Context, jobTenant *tenant.WithDetails, spec *job.Spec, dryRun bool) ([]job.ResourceURN, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: this now will always be a same time for start of service, is it correct ?

func (p JobPluginService) compileConfig(configs job.Config, tnnt *tenant.WithDetails) plugin.Configs {
	_ = "STUB: not implemented"
	return *new(plugin.Configs)
}

func (p JobPluginService) compileAsset(ctx context.Context, taskPlugin *plugin.Plugin, spec *job.Spec, scheduledAt time.Time) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
