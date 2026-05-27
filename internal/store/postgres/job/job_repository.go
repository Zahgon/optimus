package job

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
)

const (
	jobColumnsToStore = `name, version, owner, description, labels, schedule, alert, static_upstreams, http_upstreams, 
	task_name, task_config, window_spec, assets, hooks, metadata, destination, sources, project_name, namespace_name, created_at, updated_at`

	jobColumns = `id, ` + jobColumnsToStore + `, deleted_at`
)

type JobRepository struct {
	db *pgxpool.Pool
}

func NewJobRepository(pool *pgxpool.Pool) *JobRepository { _ = "STUB: not implemented"; return nil }

func (j JobRepository) Add(ctx context.Context, jobs []*job.Job) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) insertJobSpec(ctx context.Context, jobEntity *job.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) triggerInsert(ctx context.Context, jobEntity *job.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) Update(ctx context.Context, jobs []*job.Job) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) UpdateState(ctx context.Context, jobTenant tenant.Tenant, jobNames []job.Name, jobState job.State, remark string) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) SyncState(ctx context.Context, jobTenant tenant.Tenant, disabledJobNames, enabledJobNames []job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) ChangeJobNamespace(ctx context.Context, jobName job.Name, tenant, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented"
	return nil
}

func changeJobNamespace(ctx context.Context, tx pgx.Tx, jobName job.Name, tenant, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented"
	return nil
}

func changeJobUpstreamNamespace(ctx context.Context, tx pgx.Tx, jobName job.Name, tenant, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented"
	return nil
}

func changeJobRunNamespace(ctx context.Context, tx pgx.Tx, jobName job.Name, tenant, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) preCheckUpdate(ctx context.Context, jobEntity *job.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) triggerUpdate(ctx context.Context, jobEntity *job.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) get(ctx context.Context, projectName tenant.ProjectName, jobName job.Name, onlyActiveJob bool) (*Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) ResolveUpstreams(ctx context.Context, projectName tenant.ProjectName, jobNames []job.Name) (map[job.Name][]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) toJobNameWithUpstreams(storeJobsWithUpstreams []*JobWithUpstream) (map[job.Name][]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupUpstreamsPerJobFullName(upstreams []*JobWithUpstream) map[string][]*JobWithUpstream {
	_ = "STUB: not implemented"
	return nil
}

func (JobRepository) toUpstreams(storeUpstreams []*JobWithUpstream) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetByJobName(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) (*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetAllByProjectName(ctx context.Context, projectName tenant.ProjectName) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetAllByResourceDestination(ctx context.Context, resourceDestination job.ResourceURN) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func specToJob(spec *Spec) (*job.Job, error) { _ = "STUB: not implemented"; return nil, nil }

type JobWithUpstream struct {
	JobName               string         `json:"job_name"`
	ProjectName           string         `json:"project_name"`
	UpstreamJobName       sql.NullString `json:"upstream_job_name"`
	UpstreamResourceURN   sql.NullString `json:"upstream_resource_urn"`
	UpstreamProjectName   sql.NullString `json:"upstream_project_name"`
	UpstreamNamespaceName sql.NullString `json:"upstream_namespace_name"`
	UpstreamTaskName      sql.NullString `json:"upstream_task_name"`
	UpstreamHost          sql.NullString `json:"upstream_host"`
	UpstreamType          string         `json:"upstream_type"`
	UpstreamState         string         `json:"upstream_state"`
	UpstreamExternal      sql.NullBool   `json:"upstream_external"`
}

func (j *JobWithUpstream) getJobFullName() string { _ = "STUB: not implemented"; return "" }

func (j JobRepository) ReplaceUpstreams(ctx context.Context, jobsWithUpstreams []*job.WithUpstream) error {
	_ = "STUB: not implemented"
	return nil
}

func (JobRepository) insertUpstreams(ctx context.Context, tx pgx.Tx, storageJobUpstreams []*JobWithUpstream) error {
	_ = "STUB: not implemented"
	return nil
}

func (JobRepository) deleteUpstreamsByJobNames(ctx context.Context, tx pgx.Tx, jobUpstreams []string) error {
	_ = "STUB: not implemented"
	return nil
}

func toJobUpstream(jobWithUpstream *job.WithUpstream) []*JobWithUpstream {
	_ = "STUB: not implemented"
	return nil
}

// TODO: re-check this implementation as project and namespace name is not supposed to be empty within a tenant

func toNullString(val string) sql.NullString {
	_ = "STUB: not implemented"
	return *new(sql.NullString)
}

func toNullBool(val bool) sql.NullBool { _ = "STUB: not implemented"; return *new(sql.NullBool) }

type ProjectAndJobNames struct {
	ProjectName string `json:"project_name"`
	JobName     string `json:"job_name"`
}

func (j JobRepository) Delete(ctx context.Context, projectName tenant.ProjectName, jobName job.Name, cleanHistory bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) hardDelete(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) softDelete(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRepository) GetAllByTenant(ctx context.Context, jobTenant tenant.Tenant) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetUpstreams(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetDownstreamByDestination(ctx context.Context, projectName tenant.ProjectName, destination job.ResourceURN) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Downstream struct {
	JobName       string `json:"job_name"`
	ProjectName   string `json:"project_name"`
	NamespaceName string `json:"namespace_name"`
	TaskName      string `json:"task_name"`
}

func (j JobRepository) GetDownstreamByJobName(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromStoreDownstream(storeDownstreamList []Downstream) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j JobRepository) GetDownstreamBySources(ctx context.Context, sources []job.ResourceURN) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
