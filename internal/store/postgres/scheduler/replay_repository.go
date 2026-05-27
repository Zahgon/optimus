package scheduler

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

const (
	replayColumnsToStore = `job_name, namespace_name, project_name, start_time, end_time, description, parallel, job_config, status, message`
	replayColumns        = `id, ` + replayColumnsToStore + `, created_at`

	replayRunColumns       = `replay_id, scheduled_at, status`
	replayRunDetailColumns = `id as replay_id, job_name, namespace_name, project_name, start_time, end_time, description, 
parallel, job_config, r.status as replay_status, r.message as replay_message, scheduled_at, run.status as run_status, r.created_at as replay_created_at`

	updateReplayRequest = `UPDATE replay_request SET status = $1, message = $2, updated_at = NOW() WHERE id = $3`
)

type ReplayRepository struct {
	db *pgxpool.Pool
}

type replayRequest struct {
	ID uuid.UUID

	JobName       string
	NamespaceName string
	ProjectName   string

	StartTime   time.Time
	EndTime     time.Time
	Description string
	Parallel    bool
	JobConfig   map[string]string

	Status  string
	Message string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *replayRequest) toSchedulerReplayRequest() (*scheduler.Replay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type replayRun struct {
	ID uuid.UUID

	JobName       string
	NamespaceName string
	ProjectName   string

	StartTime   time.Time
	EndTime     time.Time
	Description string
	Parallel    bool
	JobConfig   map[string]string

	ReplayStatus string
	Message      string

	ScheduledTime time.Time
	RunStatus     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *replayRun) toReplayRequest() (*scheduler.Replay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *replayRun) toJobRunStatus() (*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayRepository) RegisterReplay(ctx context.Context, replay *scheduler.Replay, runs []*scheduler.JobRunStatus) (uuid.UUID, error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

// TODO: consider to store message of each run

func (r ReplayRepository) GetReplayToExecute(ctx context.Context) (*scheduler.ReplayWithRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Avoid having In Progress, but instead use row lock (for update)

func (r ReplayRepository) GetReplayRequestsByStatus(ctx context.Context, statusList []scheduler.ReplayState) ([]*scheduler.Replay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayRepository) GetReplaysByProject(ctx context.Context, projectName tenant.ProjectName, dayLimits int) ([]*scheduler.Replay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayRepository) GetReplayByID(ctx context.Context, replayID uuid.UUID) (*scheduler.ReplayWithRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toReplay(replayRuns []*replayRun) (*scheduler.ReplayWithRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayRepository) UpdateReplayStatus(ctx context.Context, id uuid.UUID, replayStatus scheduler.ReplayState, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ReplayRepository) UpdateReplay(ctx context.Context, id uuid.UUID, replayStatus scheduler.ReplayState, runs []*scheduler.JobRunStatus, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ReplayRepository) GetReplayJobConfig(ctx context.Context, jobTenant tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReplayRepository) updateReplayRequest(ctx context.Context, id uuid.UUID, replayStatus scheduler.ReplayState, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r ReplayRepository) updateReplayRuns(ctx context.Context, id uuid.UUID, runs []*scheduler.JobRunStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (ReplayRepository) insertReplay(ctx context.Context, tx pgx.Tx, replay *scheduler.Replay) error {
	_ = "STUB: not implemented"
	return nil
}

func (ReplayRepository) getReplayRequest(ctx context.Context, tx pgx.Tx, replay *scheduler.Replay) (replayRequest, error) {
	_ = "STUB: not implemented"
	return *new(replayRequest), nil
}

func (r ReplayRepository) getReplayRequestByID(ctx context.Context, replayID uuid.UUID) (replayRequest, error) {
	_ = "STUB: not implemented"
	return *new(replayRequest), nil
}

func (r ReplayRepository) getReplayRuns(ctx context.Context, replayID uuid.UUID) ([]replayRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ReplayRepository) getExecutableReplayRuns(ctx context.Context, tx pgx.Tx) ([]*replayRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ReplayRepository) insertReplayRuns(ctx context.Context, tx pgx.Tx, replayID uuid.UUID, runs []*scheduler.JobRunStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReplayRepository(db *pgxpool.Pool) *ReplayRepository { _ = "STUB: not implemented"; return nil }
