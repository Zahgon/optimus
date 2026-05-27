package models

import (
	"time"
)

type windowV2 struct {
	truncateTo string
	offset     string
	size       string
}

func (windowV2) GetVersion() int {
	_ = "STUB: not implemented"
	//nolint:gomnd
	return 0
}

func (w windowV2) Validate() error { _ = "STUB: not implemented"; return nil }

func (w windowV2) GetStartTime(scheduleTime time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (w windowV2) GetEndTime(scheduleTime time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (w windowV2) GetTruncateTo() string { _ = "STUB: not implemented"; return "" }

func (w windowV2) GetOffset() string { _ = "STUB: not implemented"; return "" }

func (w windowV2) GetSize() string { _ = "STUB: not implemented"; return "" }

func (w windowV2) validateTruncateTo() error { _ = "STUB: not implemented"; return nil }

// TODO: perhaps we can avoid using util, in hope we can remove this package

func (w windowV2) validateOffset() error { _ = "STUB: not implemented"; return nil }

func (w windowV2) validateSize() error { _ = "STUB: not implemented"; return nil }

func (w windowV2) truncateTime(scheduleTime time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// remove time upto hours

// remove time upto day

// weekday with start of the week as Monday

func (w windowV2) adjustOffset(truncatedTime time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (w windowV2) getStartTime(endTime time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// not expecting this, if this happens due to bad code just return inputTime
