package airflow

import (
	"context"
	_ "embed"
	"time"

	"github.com/raystack/salt/log"
	"gocloud.dev/blob"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/lib/cron"
)

//go:embed __lib.py
var SharedLib []byte

const (
	EntityAirflow = "Airflow"

	dagStatusBatchURL = "api/v1/dags/~/dagRuns/list"
	dagURL            = "api/v1/dags/%s"
	dagRunClearURL    = "api/v1/dags/%s/clearTaskInstances"
	dagRunCreateURL   = "api/v1/dags/%s/dagRuns"
	airflowDateFormat = "2006-01-02T15:04:05+00:00"

	schedulerHostKey = "SCHEDULER_HOST"

	baseLibFileName = "__lib.py"
	jobsDir         = "dags"
	jobsExtension   = ".py"

	concurrentTicketPerSec = 50
	concurrentLimit        = 100

	metricJobUpload       = "job_upload_total"
	metricJobRemoval      = "job_removal_total"
	metricJobStateSuccess = "success"
	metricJobStateFailed  = "failed"
)

type Bucket interface {
	WriteAll(ctx context.Context, key string, p []byte, opts *blob.WriterOptions) error
	List(opts *blob.ListOptions) *blob.ListIterator
	Delete(ctx context.Context, key string) error
	Close() error
}

type BucketFactory interface {
	New(ctx context.Context, tenant tenant.Tenant) (Bucket, error)
}

type DagCompiler interface {
	Compile(job *scheduler.JobWithDetails) ([]byte, error)
}

type Client interface {
	Invoke(ctx context.Context, r airflowRequest, auth SchedulerAuth) ([]byte, error)
}

type SecretGetter interface {
	Get(ctx context.Context, projName tenant.ProjectName, namespaceName, name string) (*tenant.PlainTextSecret, error)
}

type ProjectGetter interface {
	Get(context.Context, tenant.ProjectName) (*tenant.Project, error)
}

type Scheduler struct {
	l         log.Logger
	bucketFac BucketFactory
	client    Client
	compiler  DagCompiler

	projectGetter ProjectGetter
	secretGetter  SecretGetter
}

func (s *Scheduler) DeployJobs(ctx context.Context, tenant tenant.Tenant, jobs []*scheduler.JobWithDetails) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO list jobs should not refer from the scheduler, rather should list from db and it has nothing to do with scheduler.
func (s *Scheduler) ListJobs(ctx context.Context, t tenant.Tenant) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get all items under namespace directory

func (s *Scheduler) DeleteJobs(ctx context.Context, t tenant.Tenant, jobNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore missing files

// deleteDirectoryIfEmpty remove jobs Folder if it exists
func deleteDirectoryIfEmpty(ctx context.Context, nsDirectoryIdentifier string, bucket Bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scheduler) compileAndUpload(ctx context.Context, job *scheduler.JobWithDetails, bucket Bucket) error {
	_ = "STUB: not implemented"
	return nil
}

func pathFromJobName(prefix, namespace, jobName, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}

func pathForJobDirectory(prefix, namespace string) string { _ = "STUB: not implemented"; return "" }

func jobNameFromPath(filePath, suffix string) string { _ = "STUB: not implemented"; return "" }

func (s *Scheduler) GetJobRuns(ctx context.Context, tnnt tenant.Tenant, jobQuery *scheduler.JobRunsCriteria, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateJobState set the state of jobs as enabled / disabled on scheduler
func (s *Scheduler) UpdateJobState(ctx context.Context, tnnt tenant.Tenant, jobNames []job.Name, state string) error {
	_ = "STUB: not implemented"
	return nil
}

func getDagRunRequest(jobQuery *scheduler.JobRunsCriteria, jobCron *cron.ScheduleSpec) DagRunRequest {
	_ = "STUB: not implemented"
	return *new(DagRunRequest)
}

func (s *Scheduler) getSchedulerAuth(ctx context.Context, tnnt tenant.Tenant) (SchedulerAuth, error) {
	_ = "STUB: not implemented"
	return *new(SchedulerAuth), nil
}

func (s *Scheduler) Clear(ctx context.Context, t tenant.Tenant, jobName scheduler.JobName, executionTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scheduler) ClearBatch(ctx context.Context, tnnt tenant.Tenant, jobName scheduler.JobName, startExecutionTime, endExecutionTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scheduler) CreateRun(ctx context.Context, tnnt tenant.Tenant, jobName scheduler.JobName, executionTime time.Time, dagRunIDPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func NewScheduler(l log.Logger, bucketFac BucketFactory, client Client, compiler DagCompiler, projectGetter ProjectGetter, secretGetter SecretGetter) *Scheduler {
	_ = "STUB: not implemented"
	return nil
}

func raiseSchedulerMetric(jobTenant tenant.Tenant, metricName, status string, metricValue int) {
	_ = "STUB: not implemented"
	return
}
