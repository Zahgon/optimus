package event

import (
	"github.com/raystack/optimus/core/scheduler"
	pbInt "github.com/raystack/optimus/protos/raystack/optimus/integration/v1beta1"
)

type JobRunWaitUpstream struct {
	Event

	JobRun *scheduler.JobRun
}

func (j *JobRunWaitUpstream) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobRunInProgress struct {
	Event

	JobRun *scheduler.JobRun
}

func (j *JobRunInProgress) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobRunSuccess struct {
	Event

	JobRun *scheduler.JobRun
}

func (j *JobRunSuccess) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type JobRunFailed struct {
	Event

	JobRun *scheduler.JobRun
}

func (j *JobRunFailed) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewJobRunWaitUpstreamEvent(jobRun *scheduler.JobRun) (*JobRunWaitUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobRunInProgressEvent(jobRun *scheduler.JobRun) (*JobRunInProgress, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobRunSuccessEvent(jobRun *scheduler.JobRun) (*JobRunSuccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobRunFailedEvent(jobRun *scheduler.JobRun) (*JobRunFailed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toOptimusChangeEvent(j *scheduler.JobRun, e Event, eventType pbInt.OptimusChangeEvent_EventType) *pbInt.OptimusChangeEvent {
	_ = "STUB: not implemented"
	return nil
}
