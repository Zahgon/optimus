package scheduler

import (
	"time"

	"github.com/raystack/optimus/core/tenant"
)

type (
	EventName        string
	JobEventType     string
	JobEventCategory string
)

const (
	EntityEvent = "event"

	ISODateFormat = "2006-01-02T15:04:05Z"

	EventCategorySLAMiss    JobEventCategory = "sla_miss"
	EventCategoryJobFailure JobEventCategory = "failure"

	SLAMissEvent    JobEventType = "sla_miss"
	JobFailureEvent JobEventType = "failure"
	JobSuccessEvent JobEventType = "job_success"

	TaskStartEvent   JobEventType = "task_start"
	TaskRetryEvent   JobEventType = "task_retry"
	TaskFailEvent    JobEventType = "task_fail"
	TaskSuccessEvent JobEventType = "task_success"

	HookStartEvent   JobEventType = "hook_start"
	HookRetryEvent   JobEventType = "hook_retry"
	HookFailEvent    JobEventType = "hook_fail"
	HookSuccessEvent JobEventType = "hook_success"

	SensorStartEvent   JobEventType = "sensor_start"
	SensorRetryEvent   JobEventType = "sensor_retry"
	SensorFailEvent    JobEventType = "sensor_fail"
	SensorSuccessEvent JobEventType = "sensor_success"
)

func FromStringToEventType(name string) (JobEventType, error) {
	_ = "STUB: not implemented"
	return *new(JobEventType), nil
}

type SLAObject struct {
	JobName        JobName
	JobScheduledAt time.Time
}

func (s *SLAObject) String() string { _ = "STUB: not implemented"; return "" }

type Event struct {
	JobName        JobName
	Tenant         tenant.Tenant
	Type           JobEventType
	EventTime      time.Time
	OperatorName   string
	Status         State
	JobScheduledAt time.Time
	Values         map[string]any
	SLAObjectList  []*SLAObject
}

func (event JobEventType) IsOfType(category JobEventCategory) bool {
	_ = "STUB: not implemented"
	return false
}

func (event JobEventType) String() string { _ = "STUB: not implemented"; return "" }

func EventFrom(eventTypeName string, eventValues map[string]any, jobName JobName, tenent tenant.Tenant) (*Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
