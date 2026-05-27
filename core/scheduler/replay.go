package scheduler

import (
	"time"

	"github.com/google/uuid"

	"github.com/raystack/optimus/core/tenant"
)

const (
	// initial state
	ReplayStateCreated ReplayState = "created"

	// running state
	ReplayStateInProgress      ReplayState = "in progress"
	ReplayStatePartialReplayed ReplayState = "partial replayed"
	ReplayStateReplayed        ReplayState = "replayed"

	// terminal state
	ReplayStateInvalid ReplayState = "invalid"
	ReplayStateSuccess ReplayState = "success"
	ReplayStateFailed  ReplayState = "failed"

	// state on presentation layer
	ReplayUserStateCreated    ReplayUserState = "created"
	ReplayUserStateInProgress ReplayUserState = "in progress"
	ReplayUserStateInvalid    ReplayUserState = "invalid"
	ReplayUserStateSuccess    ReplayUserState = "success"
	ReplayUserStateFailed     ReplayUserState = "failed"

	EntityReplay = "replay"
)

type (
	ReplayState     string // contract status for business layer
	ReplayUserState string // contract status for presentation layer
)

func ReplayStateFromString(state string) (ReplayState, error) {
	_ = "STUB: not implemented"
	return *new(ReplayState), nil
}

func (j ReplayState) String() string { _ = "STUB: not implemented"; return "" }

func (j ReplayUserState) String() string { _ = "STUB: not implemented"; return "" }

type Replay struct {
	id uuid.UUID

	jobName JobName
	tenant  tenant.Tenant
	config  *ReplayConfig

	state   ReplayState
	message string

	createdAt time.Time
}

func (r *Replay) ID() uuid.UUID { _ = "STUB: not implemented"; return *new(uuid.UUID) }

func (r *Replay) JobName() JobName { _ = "STUB: not implemented"; return *new(JobName) }

func (r *Replay) Tenant() tenant.Tenant { _ = "STUB: not implemented"; return *new(tenant.Tenant) }

func (r *Replay) Config() *ReplayConfig { _ = "STUB: not implemented"; return nil }

func (r *Replay) State() ReplayState { _ = "STUB: not implemented"; return *new(ReplayState) }

func (r *Replay) UserState() ReplayUserState {
	_ = "STUB: not implemented"
	return *new(ReplayUserState)
}

func (r *Replay) Message() string { _ = "STUB: not implemented"; return "" }

func (r *Replay) CreatedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func NewReplayRequest(jobName JobName, tenant tenant.Tenant, config *ReplayConfig, state ReplayState) *Replay {
	_ = "STUB: not implemented"
	return nil
}

func NewReplay(id uuid.UUID, jobName JobName, tenant tenant.Tenant, config *ReplayConfig, state ReplayState, createdAt time.Time) *Replay {
	_ = "STUB: not implemented"
	return nil
}

type ReplayWithRun struct {
	Replay *Replay
	Runs   []*JobRunStatus // TODO: JobRunStatus does not have `message/log`
}

func (r *ReplayWithRun) GetFirstExecutableRun() *JobRunStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReplayWithRun) GetLastExecutableRun() *JobRunStatus { _ = "STUB: not implemented"; return nil }

type ReplayConfig struct {
	StartTime   time.Time
	EndTime     time.Time
	Parallel    bool
	JobConfig   map[string]string
	Description string
}

func NewReplayConfig(startTime, endTime time.Time, parallel bool, jobConfig map[string]string, description string) *ReplayConfig {
	_ = "STUB: not implemented"
	return nil
}
