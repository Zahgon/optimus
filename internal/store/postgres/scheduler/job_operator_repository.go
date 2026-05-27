package scheduler

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/scheduler"
)

const (
	sensorRunTableName = "sensor_run"
	taskRunTableName   = "task_run"
	hookRunTableName   = "hook_run"

	jobOperatorColumnsToStore = `name, job_run_id, status, start_time, end_time`
	jobOperatorColumns        = `id, ` + jobOperatorColumnsToStore
)

type OperatorRunRepository struct {
	db *pgxpool.Pool
}

type operatorRun struct {
	ID       uuid.UUID
	JobRunID uuid.UUID

	Name         string
	OperatorType string
	Status       string

	StartTime time.Time
	EndTime   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	// TODO:  add a remarks colum to capture failure reason
	DeletedAt sql.NullTime
}

func operatorTypeToTableName(operatorType scheduler.OperatorType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *operatorRun) toOperatorRun() (*scheduler.OperatorRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OperatorRunRepository) GetOperatorRun(ctx context.Context, name string, operatorType scheduler.OperatorType, jobRunID uuid.UUID) (*scheduler.OperatorRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OperatorRunRepository) CreateOperatorRun(ctx context.Context, name string, operatorType scheduler.OperatorType, jobRunID uuid.UUID, startTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *OperatorRunRepository) UpdateOperatorRun(ctx context.Context, operatorType scheduler.OperatorType, operatorRunID uuid.UUID, eventTime time.Time, state scheduler.State) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOperatorRunRepository(pool *pgxpool.Pool) *OperatorRunRepository {
	_ = "STUB: not implemented"
	return nil
}
