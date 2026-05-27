package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/raystack/salt/log"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/lib/cron"
)

const (
	prefixReplayed = "replayed"
)

type ReplayScheduler interface {
	Clear(ctx context.Context, t tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) error
	ClearBatch(ctx context.Context, t tenant.Tenant, jobName scheduler.JobName, startTime, endTime time.Time) error

	CreateRun(ctx context.Context, tnnt tenant.Tenant, jobName scheduler.JobName, executionTime time.Time, dagRunIDPrefix string) error
	GetJobRuns(ctx context.Context, t tenant.Tenant, criteria *scheduler.JobRunsCriteria, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error)
}

type ReplayWorker struct {
	l log.Logger

	replayRepo ReplayRepository
	scheduler  ReplayScheduler

	jobRepo JobRepository

	config config.ReplayConfig
}

func NewReplayWorker(l log.Logger, replayRepo ReplayRepository, scheduler ReplayScheduler, jobRepo JobRepository, config config.ReplayConfig) *ReplayWorker {
	_ = "STUB: not implemented"
	return nil
}

type JobReplayRunService interface {
	GetJobRuns(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName, criteria *scheduler.JobRunsCriteria) ([]*scheduler.JobRunStatus, error)
}

func (w ReplayWorker) Process(replayReq *scheduler.ReplayWithRun) {
	_ = "STUB: not implemented"
	return
}

func (w ReplayWorker) createMissingRuns(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetch runs within range of replay range

// check each runs if there's no existing run from the above

// create any missing runs

func (w ReplayWorker) processNewReplayRequest(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (w ReplayWorker) processNewReplayRequestParallel(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w ReplayWorker) processNewReplayRequestSequential(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w ReplayWorker) processPartialReplayedRequest(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (w ReplayWorker) processReplayedRequest(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func identifyUpdatedRunStatus(existingJobRuns, incomingJobRuns []*scheduler.JobRunStatus) map[time.Time]scheduler.State {
	_ = "STUB: not implemented"
	return nil
}

func (w ReplayWorker) getJobCron(ctx context.Context, replayReq *scheduler.ReplayWithRun) (*cron.ScheduleSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w ReplayWorker) fetchRuns(ctx context.Context, replayReq *scheduler.ReplayWithRun, jobCron *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w ReplayWorker) updateReplayAsFailed(ctx context.Context, replayID uuid.UUID, message string) {
	_ = "STUB: not implemented"
	return
}

func raiseReplayMetric(t tenant.Tenant, jobName scheduler.JobName, state scheduler.ReplayState) {
	_ = "STUB: not implemented"
	return
}
