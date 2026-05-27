package scheduler

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/models"
)

const (
	jobColumns = `id, name, version, owner, description, labels, schedule, alert, static_upstreams, http_upstreams,
				  task_name, task_config, window_spec, assets, hooks, metadata, destination, sources, project_name, namespace_name, created_at, updated_at`
	upstreamColumns = `
    job_name, project_name, upstream_job_name, upstream_project_name, upstream_host,
    upstream_namespace_name, upstream_resource_urn, upstream_task_name, upstream_type, upstream_external, upstream_state`
)

type JobRepository struct {
	db *pgxpool.Pool
}

type Schedule struct {
	StartDate     time.Time
	EndDate       *time.Time
	Interval      string
	DependsOnPast bool
	Retry         *Retry
}
type Retry struct {
	Count              int   `json:"count"`
	Delay              int32 `json:"delay"`
	ExponentialBackoff bool
}

type JobUpstreams struct {
	JobID                 uuid.UUID
	JobName               string
	ProjectName           string
	UpstreamJobID         uuid.UUID
	UpstreamJobName       sql.NullString
	UpstreamResourceUrn   sql.NullString
	UpstreamProjectName   sql.NullString
	UpstreamNamespaceName sql.NullString
	UpstreamTaskName      sql.NullString
	UpstreamHost          sql.NullString
	UpstreamType          string
	UpstreamState         string
	UpstreamExternal      sql.NullBool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (j *JobUpstreams) toJobUpstreams() (*scheduler.JobUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Job struct {
	ID          uuid.UUID
	Name        string
	Version     int
	Owner       string
	Description string
	Labels      map[string]string

	Schedule   json.RawMessage
	WindowSpec json.RawMessage

	Alert json.RawMessage

	StaticUpstreams pq.StringArray
	HTTPUpstreams   json.RawMessage

	TaskName   string
	TaskConfig map[string]string

	Hooks json.RawMessage

	Assets map[string]string

	Metadata json.RawMessage

	Destination string
	Sources     pq.StringArray

	ProjectName   string `json:"project_name"`
	NamespaceName string `json:"namespace_name"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
type Window struct {
	WindowSize       string
	WindowOffset     string
	WindowTruncateTo string
}

func fromStorageWindow(raw []byte, jobVersion int) (models.Window, error) {
	_ = "STUB: not implemented"
	return *new(models.Window), nil
}

type Metadata struct {
	Resource  *MetadataResource
	Scheduler map[string]string
}

type MetadataResource struct {
	Request *MetadataResourceConfig
	Limit   *MetadataResourceConfig
}

type MetadataResourceConfig struct {
	CPU    string
	Memory string
}

func fromStorageMetadata(metadata json.RawMessage) (scheduler.RuntimeConfig, error) {
	_ = "STUB: not implemented"
	return *new(scheduler.RuntimeConfig), nil
}

func (j *Job) toJob() (*scheduler.Job, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *Job) toJobWithDetails() (*scheduler.JobWithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FromRow(row pgx.Row) (*Job, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *JobRepository) GetJob(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName) (*scheduler.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRepository) GetJobDetails(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName) (*scheduler.JobWithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupUpstreamsByJobName(jobUpstreams []*JobUpstreams) (map[string][]*scheduler.JobUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRepository) getJobsUpstreams(ctx context.Context, projectName tenant.ProjectName, jobNames []string) (map[string][]*scheduler.JobUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRepository) GetAll(ctx context.Context, projectName tenant.ProjectName) ([]*scheduler.JobWithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRepository) GetJobs(ctx context.Context, projectName tenant.ProjectName, jobs []string) ([]*scheduler.JobWithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobProviderRepository(pool *pgxpool.Pool) *JobRepository {
	_ = "STUB: not implemented"
	return nil
}
