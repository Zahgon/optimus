package scheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

const (
	columnsToStore = `job_name, namespace_name, project_name, scheduled_at, start_time, end_time, status, sla_definition, sla_alert`
	jobRunColumns  = `id, ` + columnsToStore + `, monitoring`
)

type JobRunRepository struct {
	db *pgxpool.Pool
}

type jobRun struct {
	ID uuid.UUID

	JobName       string
	NamespaceName string
	ProjectName   string

	ScheduledAt time.Time
	StartTime   time.Time
	EndTime     time.Time

	Status        string
	SLAAlert      bool
	SLADefinition int64

	CreatedAt time.Time
	UpdatedAt time.Time

	Monitoring json.RawMessage
}

func (j *jobRun) toJobRun() (*scheduler.JobRun, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *JobRunRepository) GetByID(ctx context.Context, id scheduler.JobRunID) (*scheduler.JobRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRunRepository) GetByScheduledAt(ctx context.Context, t tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time) (*scheduler.JobRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobRunRepository) UpdateState(ctx context.Context, jobRunID uuid.UUID, status scheduler.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobRunRepository) Update(ctx context.Context, jobRunID uuid.UUID, endTime time.Time, status scheduler.State) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobRunRepository) UpdateSLA(ctx context.Context, slaObjects []*scheduler.SLAObject) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobRunRepository) UpdateMonitoring(ctx context.Context, jobRunID uuid.UUID, monitoringValues map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobRunRepository) Create(ctx context.Context, t tenant.Tenant, jobName scheduler.JobName, scheduledAt time.Time, slaDefinitionInSec int64) error {
	_ = "STUB: not implemented"
	return nil
}

func NewJobRunRepository(pool *pgxpool.Pool) *JobRunRepository {
	_ = "STUB: not implemented"
	return nil
}
