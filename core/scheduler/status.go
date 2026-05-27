package scheduler

import (
	"time"

	"github.com/raystack/optimus/internal/lib/cron"
)

const (
	StatePending State = "pending"

	StateAccepted State = "accepted"
	StateRunning  State = "running"
	StateQueued   State = "queued"

	StateRetry State = "retried"

	StateSuccess State = "success"
	StateFailed  State = "failed"

	StateWaitUpstream State = "wait_upstream"
	StateInProgress   State = "in_progress"
)

var TaskEndStates = []State{StateSuccess, StateFailed, StateRetry}

type State string

func StateFromString(state string) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (j State) String() string { _ = "STUB: not implemented"; return "" }

type JobRunStatus struct {
	ScheduledAt time.Time
	State       State
}

func JobRunStatusFrom(scheduledAt time.Time, state string) (JobRunStatus, error) {
	_ = "STUB: not implemented"
	return *new(JobRunStatus), nil
}

func (j JobRunStatus) GetLogicalTime(jobCron *cron.ScheduleSpec) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

type JobRunStatusList []*JobRunStatus

func (j JobRunStatusList) GetSortedRunsByStates(states []State) []*JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRunStatusList) GetSortedRunsByScheduledAt() []*JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRunStatusList) MergeWithUpdatedRuns(updatedRunMap map[time.Time]State) []*JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func (j JobRunStatusList) ToRunStatusMap() map[time.Time]State {
	_ = "STUB: not implemented"
	return nil
}

// JobRunsCriteria represents the filter condition to get run status from scheduler
type JobRunsCriteria struct {
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Filter      []string
	OnlyLastRun bool
}

func (c *JobRunsCriteria) ExecutionStart(cron *cron.ScheduleSpec) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *JobRunsCriteria) ExecutionEndDate(jobCron *cron.ScheduleSpec) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// when the current time matches one of the schedule times execution time means previous schedule.

// else it is previous to previous schedule.
