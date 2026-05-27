package compiler

import (
	"text/template"
	"time"
)

func OptimusFuncMap() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }

func Date(timeStr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func Replace(old, newStr, name string) string { _ = "STUB: not implemented"; return "" }

func Trunc(c int, s string) string { _ = "STUB: not implemented"; return "" }

func date(fmt string, date interface{}) string {
	_ = "STUB: not implemented"
	// Cannot have a reliable test, depends on local machine time
	return ""
}

func dateInZone(fmt string, date interface{}, zone string) string {
	_ = "STUB: not implemented"
	return ""
}

func DateModify(fmt string, date time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func toDate(fmt, str string) time.Time {
	_ = "STUB: not implemented"
	// Cannot have a reliable test, depends on local machine time
	return *new(time.Time)
}

func UnixEpoch(date time.Time) string { _ = "STUB: not implemented"; return "" }

//nolint

func List(v ...string) []string { _ = "STUB: not implemented"; return nil }

func Join(sep string, v []string) string { _ = "STUB: not implemented"; return "" }
