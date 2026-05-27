package service

import (
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/internal/lib/cron"
)

var replayStatusToValidate = []scheduler.ReplayState{
	scheduler.ReplayStateCreated, scheduler.ReplayStateInProgress,
	scheduler.ReplayStatePartialReplayed, scheduler.ReplayStateReplayed,
}

type Validator struct {
	replayRepository ReplayRepository
	scheduler        ReplayScheduler
	jobRepo          JobRepository
}

func NewValidator(replayRepository ReplayRepository, scheduler ReplayScheduler, jobRepo JobRepository) *Validator {
	_ = "STUB: not implemented"
	return nil
}

func (v Validator) Validate(ctx context.Context, replayRequest *scheduler.Replay, jobCron *cron.ScheduleSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v Validator) validateDateRange(ctx context.Context, replayRequest *scheduler.Replay) error {
	_ = "STUB: not implemented"
	return nil
}

// time bound for end date

func (v Validator) validateConflictedReplay(ctx context.Context, replayRequest *scheduler.Replay) error {
	_ = "STUB: not implemented"
	return nil
}

// Check any intersection of date range

func (v Validator) validateConflictedRun(ctx context.Context, replayRequest *scheduler.Replay, jobCron *cron.ScheduleSpec) error {
	_ = "STUB: not implemented"
	return nil
}
