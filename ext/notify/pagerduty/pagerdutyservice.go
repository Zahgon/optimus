package pagerduty

import (
	"context"
)

type PagerDutyService interface {
	SendAlert(context.Context, Event) error
}

type PagerDutyServiceImpl struct{}

type customDetails struct {
	Owner     string `json:"owner"`
	Namespace string `json:"namespace"`
	LogURL    string `json:"log_url"`
	JobURL    string `json:"job_url"`
	Exception string `json:"exception"`
	Message   string `json:"message"`
}

func buildPayloadCustomDetails(evt Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*PagerDutyServiceImpl) SendAlert(ctx context.Context, evt Event) error {
	_ = "STUB: not implemented"
	return nil
}
