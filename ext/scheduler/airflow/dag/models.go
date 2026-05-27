package dag

import (
	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	EntitySchedulerAirflow = "schedulerAirflow"
)

type TemplateContext struct {
	JobDetails *scheduler.JobWithDetails

	Tenant          tenant.Tenant
	Version         string
	SLAMissDuration int64
	Hostname        string
	ExecutorTask    string
	ExecutorHook    string

	RuntimeConfig RuntimeConfig
	Task          Task
	Hooks         Hooks
	Priority      int
	Upstreams     Upstreams
}

type Task struct {
	Name       string
	Image      string
	Entrypoint plugin.Entrypoint
}

func PrepareTask(job *scheduler.Job, pluginRepo PluginRepo) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

type Hook struct {
	Name       string
	Image      string
	Entrypoint plugin.Entrypoint
	IsFailHook bool
}

type Hooks struct {
	Pre          []Hook
	Post         []Hook
	Fail         []Hook
	Dependencies map[string]string
}

func (h Hooks) List() []Hook {
	_ = "STUB: not implemented" //nolint: gocritic
	return nil
}

func PrepareHooksForJob(job *scheduler.Job, pluginRepo PluginRepo) (Hooks, error) {
	_ = "STUB: not implemented"
	return *new(Hooks), nil
}

type RuntimeConfig struct {
	Resource *Resource
	Airflow  AirflowConfig
}

func SetupRuntimeConfig(jobDetails *scheduler.JobWithDetails) RuntimeConfig {
	_ = "STUB: not implemented"
	return *new(RuntimeConfig)
}

type Resource struct {
	Request *ResourceConfig
	Limit   *ResourceConfig
}

func ToResource(resource *scheduler.Resource) *Resource { _ = "STUB: not implemented"; return nil }

type ResourceConfig struct {
	CPU    string
	Memory string
}

func ToResourceConfig(config *scheduler.ResourceConfig) *ResourceConfig {
	_ = "STUB: not implemented"
	return nil
}

type AirflowConfig struct {
	Pool  string
	Queue string
}

func ToAirflowConfig(schedulerConf map[string]string) AirflowConfig {
	_ = "STUB: not implemented"
	return *new(AirflowConfig)
}

func SLAMissDuration(job *scheduler.JobWithDetails) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We are ranging and picking one value
