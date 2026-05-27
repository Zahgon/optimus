package scheduler

import (
	"time"

	"github.com/google/uuid"

	"github.com/raystack/optimus/core/tenant"
)

type JobRunID uuid.UUID

func JobRunIDFromString(runID string) (JobRunID, error) {
	_ = "STUB: not implemented"
	return *new(JobRunID), nil
}

func (i JobRunID) UUID() uuid.UUID { _ = "STUB: not implemented"; return *new(uuid.UUID) }

func (i JobRunID) IsEmpty() bool { _ = "STUB: not implemented"; return false }

type JobRun struct {
	ID uuid.UUID

	JobName     JobName
	Tenant      tenant.Tenant
	State       State
	ScheduledAt time.Time
	StartTime   time.Time
	SLAAlert    bool
	EndTime     time.Time

	Monitoring map[string]any
}

type OperatorRun struct {
	ID           uuid.UUID
	Name         string
	JobRunID     uuid.UUID
	OperatorType OperatorType
	Status       State
	StartTime    time.Time
	EndTime      time.Time
}

type NotifyAttrs struct {
	Owner    string
	JobEvent *Event
	Route    string
	Secret   string
}

const (
	MetricNotificationQueue         = "notification_queue_total"
	MetricNotificationWorkerBatch   = "notification_worker_batch_total"
	MetricNotificationWorkerSendErr = "notification_worker_send_err_total"
)
