package models

import (
	"time"
)

type Window interface {
	Validate() error

	GetStartTime(scheduleTime time.Time) (time.Time, error)
	GetEndTime(scheduleTime time.Time) (time.Time, error)
	GetTruncateTo() string
	GetOffset() string
	GetSize() string
	GetVersion() int
}

func NewWindow(version int, truncateTo, offset, size string) (Window, error) {
	_ = "STUB: not implemented"
	return *new(Window), nil
}

// nolint:gomnd

// GetEndRunDate subtract 1 day to make end inclusive
func GetEndRunDate(runTime time.Time, window Window) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func monthsAndNonMonthExpression(durationExpression string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

// duration contains only month

// if duration is negative then use the negative duration for both the splits.
