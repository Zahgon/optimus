package event

import (
	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
	pbInt "github.com/raystack/optimus/protos/raystack/optimus/integration/v1beta1"
)

type JobCreated struct {
	Event

	Job *job.Job
}

func NewJobCreatedEvent(job *job.Job) (*JobCreated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobCreated) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobUpdated struct {
	Event

	Job *job.Job
}

func NewJobUpdateEvent(job *job.Job) (*JobUpdated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobUpdated) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobDeleted struct {
	Event

	JobName   job.Name
	JobTenant tenant.Tenant
}

func NewJobDeleteEvent(tnnt tenant.Tenant, jobName job.Name) (*JobDeleted, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobDeleted) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobStateChange struct {
	Event

	JobName   job.Name
	JobTenant tenant.Tenant
	State     job.State
}

func NewJobStateChangeEvent(tnnt tenant.Tenant, jobName job.Name, state job.State) (*JobStateChange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobStateChange) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func jobEventToBytes(event Event, job *job.Job, eventType pbInt.OptimusChangeEvent_EventType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
