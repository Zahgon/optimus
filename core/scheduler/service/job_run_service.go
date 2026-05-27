package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/event/moderator"
	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/lib/cron"
)

type metricType string

func (m metricType) String() string { _ = "STUB: not implemented"; return "" }

const (
	scheduleDelay metricType = "schedule_delay"

	metricJobRunEvents = "jobrun_events_total"
)

type JobRepository interface {
	GetJob(ctx context.Context, name tenant.ProjectName, jobName scheduler.JobName) (*scheduler.Job, error)
	GetJobDetails(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName) (*scheduler.JobWithDetails, error)
	GetAll(ctx context.Context, projectName tenant.ProjectName) ([]*scheduler.JobWithDetails, error)
	GetJobs(ctx context.Context, projectName tenant.ProjectName, jobs []string) ([]*scheduler.JobWithDetails, error)
}

type JobRunRepository interface {
	GetByID(ctx context.Context, id scheduler.JobRunID) (*scheduler.JobRun, error)
	GetByScheduledAt(ctx context.Context, tenant tenant.Tenant, name scheduler.JobName, scheduledAt time.Time) (*scheduler.JobRun, error)
	Create(ctx context.Context, tenant tenant.Tenant, name scheduler.JobName, scheduledAt time.Time, slaDefinitionInSec int64) error
	Update(ctx context.Context, jobRunID uuid.UUID, endTime time.Time, jobRunStatus scheduler.State) error
	UpdateState(ctx context.Context, jobRunID uuid.UUID, jobRunStatus scheduler.State) error
	UpdateSLA(ctx context.Context, slaObjects []*scheduler.SLAObject) error
	UpdateMonitoring(ctx context.Context, jobRunID uuid.UUID, monitoring map[string]any) error
}

type JobReplayRepository interface {
	GetReplayJobConfig(ctx context.Context, jobTenant tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) (map[string]string, error)
}

type OperatorRunRepository interface {
	GetOperatorRun(ctx context.Context, operatorName string, operator scheduler.OperatorType, jobRunID uuid.UUID) (*scheduler.OperatorRun, error)
	CreateOperatorRun(ctx context.Context, operatorName string, operator scheduler.OperatorType, jobRunID uuid.UUID, startTime time.Time) error
	UpdateOperatorRun(ctx context.Context, operator scheduler.OperatorType, jobRunID uuid.UUID, eventTime time.Time, state scheduler.State) error
}

type JobInputCompiler interface {
	Compile(ctx context.Context, job *scheduler.Job, config scheduler.RunConfig, executedAt time.Time) (*scheduler.ExecutorInput, error)
}

type PriorityResolver interface {
	Resolve(context.Context, []*scheduler.JobWithDetails) error
}

type Scheduler interface {
	GetJobRuns(ctx context.Context, t tenant.Tenant, criteria *scheduler.JobRunsCriteria, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error)
	DeployJobs(ctx context.Context, t tenant.Tenant, jobs []*scheduler.JobWithDetails) error
	ListJobs(ctx context.Context, t tenant.Tenant) ([]string, error)
	DeleteJobs(ctx context.Context, t tenant.Tenant, jobsToDelete []string) error
}

type EventHandler interface {
	HandleEvent(moderator.Event)
}

type JobRunService struct {
	l                log.Logger
	repo             JobRunRepository
	replayRepo       JobReplayRepository
	operatorRunRepo  OperatorRunRepository
	eventHandler     EventHandler
	scheduler        Scheduler
	jobRepo          JobRepository
	priorityResolver PriorityResolver
	compiler         JobInputCompiler
}

func (s *JobRunService) JobRunInput(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName, config scheduler.RunConfig) (*scheduler.ExecutorInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Use scheduled_at instead of executed_at for computations, for deterministic calculations
// Todo: later, always return scheduleTime, for scheduleTimes greater than a given date

// Fallback for executed_at to scheduled_at

// Additional task config from existing replay

func (s *JobRunService) GetJobRuns(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName, criteria *scheduler.JobRunsCriteria) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getExpectedRuns(spec *cron.ScheduleSpec, startTime, endTime time.Time) []*scheduler.JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func mergeRuns(expected, actual []*scheduler.JobRunStatus) []*scheduler.JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func actualRunMap(runs []*scheduler.JobRunStatus) map[string]scheduler.JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func filterRuns(runs []*scheduler.JobRunStatus, filter map[string]struct{}) []*scheduler.JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func createFilterSet(filter []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func validateJobQuery(jobQuery *scheduler.JobRunsCriteria, jobWithDetails *scheduler.JobWithDetails) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) registerNewJobRun(ctx context.Context, tenant tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) getJobRunByScheduledAt(ctx context.Context, tenant tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) (*scheduler.JobRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: consider moving below call outside as the caller is a 'getter'

func (s *JobRunService) updateJobRun(ctx context.Context, event *scheduler.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (*JobRunService) getMonitoringValues(event *scheduler.Event) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) updateJobRunSLA(ctx context.Context, event *scheduler.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func operatorStartToJobState(operatorType scheduler.OperatorType) (scheduler.State, error) {
	_ = "STUB: not implemented"
	return *new(scheduler.State), nil
}

func (s *JobRunService) raiseJobRunStateChangeEvent(jobRun *scheduler.JobRun) {
	_ = "STUB: not implemented"
	return
}

func (s *JobRunService) createOperatorRun(ctx context.Context, event *scheduler.Event, operatorType scheduler.OperatorType) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) getOperatorRun(ctx context.Context, event *scheduler.Event, operatorType scheduler.OperatorType, jobRunID uuid.UUID) (*scheduler.OperatorRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: consider moving below call outside as the caller is a 'getter'

func (s *JobRunService) updateOperatorRun(ctx context.Context, event *scheduler.Event, operatorType scheduler.OperatorType) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) trackEvent(event *scheduler.Event) { _ = "STUB: not implemented"; return }

func (s *JobRunService) UpdateJobState(ctx context.Context, event *scheduler.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func NewJobRunService(logger log.Logger, jobRepo JobRepository, jobRunRepo JobRunRepository, replayRepo JobReplayRepository,
	operatorRunRepo OperatorRunRepository, scheduler Scheduler, resolver PriorityResolver, compiler JobInputCompiler, eventHandler EventHandler,
) *JobRunService {
	_ = "STUB: not implemented"
	return nil
}
