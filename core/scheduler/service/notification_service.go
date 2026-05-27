package service

import (
	"context"
	"io"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/scheduler"
)

const (
	NotificationSchemeSlack     = "slack"
	NotificationSchemePagerDuty = "pagerduty"
)

type Notifier interface {
	io.Closer
	Notify(ctx context.Context, attr scheduler.NotifyAttrs) error
}

type NotifyService struct {
	notifyChannels map[string]Notifier
	jobRepo        JobRepository
	tenantService  TenantService
	l              log.Logger
}

func (n *NotifyService) Push(ctx context.Context, event *scheduler.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NotifyService) Close() error { _ = "STUB: not implemented"; return nil }

func NewNotifyService(l log.Logger, jobRepo JobRepository, tenantService TenantService, notifyChan map[string]Notifier) *NotifyService {
	_ = "STUB: not implemented"
	return nil
}
