package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

// CronIntervalValidator return a nil value when a valid cron string is passed
// used in gopkg.in/validator.v2
func CronIntervalValidator(val interface{}, _ string) error { _ = "STUB: not implemented"; return nil }

// an empty schedule is a valid schedule

// validatorFactory, name abbreviated so that
// the global implementation can be called 'validatorFactory'
type VFactory struct{}

func (*VFactory) NewFromRegex(re, message string) survey.Validator {
	_ = "STUB: not implemented"
	return *new(survey.Validator)
}

var ValidatorFactory = new(VFactory)

// ValidateCronInterval return a nil value when a valid cron string is passed
func ValidateCronInterval(val interface{}) error { _ = "STUB: not implemented"; return nil }
