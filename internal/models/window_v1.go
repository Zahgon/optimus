package models

import (
	"regexp"
	"time"
)

const (
	HoursInDay   = 24
	HoursInMonth = HoursInDay * 30
)

var monthExp = regexp.MustCompile("(\\+|-)?([0-9]+)(M)") //nolint:gosimple

type windowV1 struct {
	truncateTo string
	offset     string
	size       string
}

func (w windowV1) Validate() error { _ = "STUB: not implemented"; return nil }

// nolint:dogsled

func (windowV1) GetVersion() int { _ = "STUB: not implemented"; return 0 }

func (w windowV1) GetTruncateTo() string { _ = "STUB: not implemented"; return "" }

func (w windowV1) GetOffset() string { _ = "STUB: not implemented"; return "" }

func (w windowV1) GetSize() string { _ = "STUB: not implemented"; return "" }

func (w windowV1) GetStartTime(scheduledAt time.Time) (startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (w windowV1) GetEndTime(scheduledAt time.Time) (endTime time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

type JobSpecTaskWindow struct {
	Size       time.Duration
	Offset     time.Duration
	TruncateTo string
}

func (w *windowV1) prepareWindow() (JobSpecTaskWindow, error) {
	_ = "STUB: not implemented"
	return *new(JobSpecTaskWindow), nil
}

// check if string contains monthly notation

// treat as normal duration

// check if string contains monthly notation

// treat as normal duration

func (*JobSpecTaskWindow) getWindowDate(today time.Time, windowSize, windowOffset time.Duration, windowTruncateTo string) (time.Time, time.Time) {
	_ = "STUB: not implemented"
	return *

	// apply truncation to end
	new(time.Time), *new(time.Time)
}

// remove time upto hours

// remove time upto day

// shift current window to nearest Sunday

// handle monthly windows separately as every month is not of same size

// shift current window to nearest month start and end

// truncate the date

// then add the month offset
// for handling offset, treat 30 days as 1 month

// then find the last day of this month

// final end is computed

// truncate days/hours from window start as well

// for handling size, treat 30 days as 1 month, and as we have already truncated current month
// subtract 1 from this

// final start is computed

func (windowV1) inHrs(hrs int) string { _ = "STUB: not implemented"; return "" }

// check if string contains monthly notation
func (windowV1) tryParsingInMonths(str string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// replace month notation with days first, treating 1M as 30 days

// check if there is remaining time that we can still parse
