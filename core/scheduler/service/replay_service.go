package service

import (
	"github.com/google/uuid"
	"github.com/raystack/salt/log"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/lib/cron"
)

const (
	getReplaysDayLimit = 30 // TODO: make it configurable via cli

	metricJobReplay = "jobrun_replay_requests_total"
)

type ReplayRepository interface {
	RegisterReplay(ctx context.Context, replay *scheduler.Replay, runs []*scheduler.JobRunStatus) (uuid.UUID, error)
	UpdateReplay(ctx context.Context, replayID uuid.UUID, state scheduler.ReplayState, runs []*scheduler.JobRunStatus, message string) error
	UpdateReplayStatus(ctx context.Context, replayID uuid.UUID, state scheduler.ReplayState, message string) error

	GetReplayToExecute(context.Context) (*scheduler.ReplayWithRun, error)
	GetReplayRequestsByStatus(ctx context.Context, statusList []scheduler.ReplayState) ([]*scheduler.Replay, error)
	GetReplaysByProject(ctx context.Context, projectName tenant.ProjectName, dayLimits int) ([]*scheduler.Replay, error)
	GetReplayByID(ctx context.Context, replayID uuid.UUID) (*scheduler.ReplayWithRun, error)
}

type ReplayValidator interface {
	Validate(ctx context.Context, replayRequest *scheduler.Replay, jobCron *cron.ScheduleSpec) error
}

type ReplayService struct {
	replayRepo ReplayRepository
	jobRepo    JobRepository

	validator ReplayValidator

	logger log.Logger
}

func (r ReplayService) CreateReplay(ctx context.Context, tenant tenant.Tenant, jobName scheduler.JobName, config *scheduler.ReplayConfig) (replayID uuid.UUID, err error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

func (r ReplayService) GetReplayList(ctx context.Context, projectName tenant.ProjectName) (replays []*scheduler.Replay, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayService) GetReplayByID(ctx context.Context, replayID uuid.UUID) (*scheduler.ReplayWithRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReplayService(replayRepo ReplayRepository, jobRepo JobRepository, validator ReplayValidator, logger log.Logger) *ReplayService {
	_ = "STUB: not implemented"
	return nil
}
