package scheduler

import (
	"time"
)

const (
	ExecutorTask ExecutorType = "task"
	ExecutorHook ExecutorType = "hook"
)

type ExecutorType string

func (e ExecutorType) String() string { _ = "STUB: not implemented"; return "" }

func ExecutorTypeFrom(val string) (ExecutorType, error) {
	_ = "STUB: not implemented"
	return *new(ExecutorType), nil
}

type Executor struct {
	Name string
	Type ExecutorType
}

func ExecutorFrom(name string, executorType ExecutorType) (Executor, error) {
	_ = "STUB: not implemented"
	return *new(Executor), nil
}

func ExecutorFromEnum(name, enum string) (Executor, error) {
	_ = "STUB: not implemented"
	return *new(Executor), nil
}

type RunConfig struct {
	Executor Executor

	ScheduledAt time.Time
	JobRunID    JobRunID
}

func RunConfigFrom(executor Executor, scheduledAt time.Time, runID string) (RunConfig, error) {
	_ = "STUB: not implemented"
	return *new(RunConfig), nil
}

// runID can be empty or a valid uuid

type ConfigMap map[string]string

type ExecutorInput struct {
	Configs ConfigMap
	Secrets ConfigMap
	Files   ConfigMap
}
