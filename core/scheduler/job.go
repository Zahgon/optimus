package scheduler

import (
	"time"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/models"
)

type (
	JobName      string
	OperatorType string
)

func (o OperatorType) String() string { _ = "STUB: not implemented"; return "" }

const (
	EntityJobRun = "jobRun"

	OperatorTask   OperatorType = "task"
	OperatorSensor OperatorType = "sensor"
	OperatorHook   OperatorType = "hook"

	UpstreamTypeStatic   = "static"
	UpstreamTypeInferred = "inferred"
)

func JobNameFrom(name string) (JobName, error) {
	_ = "STUB: not implemented"
	return *new(JobName), nil
}

func (n JobName) String() string { _ = "STUB: not implemented"; return "" }

type Job struct {
	Name   JobName
	Tenant tenant.Tenant

	Destination string
	Task        *Task
	Hooks       []*Hook
	Window      models.Window
	Assets      map[string]string
}

func (j *Job) GetHook(hookName string) (*Hook, error) { _ = "STUB: not implemented"; return nil, nil }

type Task struct {
	Name   string
	Config map[string]string
}

type Hook struct {
	Name   string
	Config map[string]string
}

// JobWithDetails contains the details for a job
type JobWithDetails struct {
	Name JobName

	Job           *Job
	JobMetadata   *JobMetadata
	Schedule      *Schedule
	Retry         Retry
	Alerts        []Alert
	RuntimeConfig RuntimeConfig
	Priority      int
	Upstreams     Upstreams
}

func (j *JobWithDetails) GetName() string { _ = "STUB: not implemented"; return "" }

func GroupJobsByTenant(j []*JobWithDetails) map[tenant.Tenant][]*JobWithDetails {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobWithDetails) SLADuration() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

type JobMetadata struct {
	Version     int
	Owner       string
	Description string
	Labels      map[string]string
}

type Schedule struct {
	DependsOnPast bool
	StartDate     time.Time
	EndDate       *time.Time
	Interval      string
}

func (j *JobWithDetails) GetLabelsAsString() string { _ = "STUB: not implemented"; return "" }

func (j *JobWithDetails) GetUniqueLabelValues() []string { _ = "STUB: not implemented"; return nil }

type Retry struct {
	ExponentialBackoff bool
	Count              int
	Delay              int32
}

type Alert struct {
	On       JobEventCategory
	Channels []string
	Config   map[string]string
}

type RuntimeConfig struct {
	Resource  *Resource
	Scheduler map[string]string
}

type Resource struct {
	Request *ResourceConfig
	Limit   *ResourceConfig
}

type ResourceConfig struct {
	CPU    string
	Memory string
}

type Upstreams struct {
	HTTP         []*HTTPUpstreams
	UpstreamJobs []*JobUpstream
}

type HTTPUpstreams struct {
	Name    string
	URL     string
	Headers map[string]string
	Params  map[string]string
}

type JobUpstream struct {
	JobName        string
	Host           string
	TaskName       string        // TODO: remove after airflow migration
	DestinationURN string        //- bigquery://pilot.playground.table
	Tenant         tenant.Tenant // Current or external tenant
	Type           string
	External       bool
	State          string
}
