package pagerduty

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/raystack/optimus/core/scheduler"
)

const (
	DefaultEventBatchInterval = time.Second * 10
)

var (
	notifierType          = "pagerduty"
	pagerdutyQueueCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationQueue,
		ConstLabels: map[string]string{"type": notifierType},
	})
	pagerdutyWorkerBatchCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationWorkerBatch,
		ConstLabels: map[string]string{"type": notifierType},
	})
	pagerdutyWorkerSendErrCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationWorkerSendErr,
		ConstLabels: map[string]string{"type": notifierType},
	})
)

type Notifier struct {
	io.Closer
	msgQueue           []Event
	wg                 sync.WaitGroup
	mu                 sync.Mutex
	workerErrChan      chan error
	pdService          PagerDutyService
	eventBatchInterval time.Duration
}

type Event struct {
	routingKey string
	owner      string
	meta       *scheduler.Event
}

func NewEvent(routingKey, owner string, meta *scheduler.Event) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

func (s *Notifier) Notify(_ context.Context, attr scheduler.NotifyAttrs) error {
	_ = "STUB: not implemented" //nolint:unparam
	return nil
}

func (s *Notifier) queueNotification(routingKey string, attr scheduler.NotifyAttrs) {
	_ = "STUB: not implemented"
	return
}

func (s *Notifier) Worker(ctx context.Context) { _ = "STUB: not implemented"; return }

// empty the queue

func (s *Notifier) Close() error {
	_ = "STUB: not implemented" // nolint: unparam
	// drain batches
	return nil
}

func NewNotifier(ctx context.Context, eventBatchInterval time.Duration, errHandler func(error), pdService PagerDutyService) *Notifier {
	_ = "STUB: not implemented"
	return nil
}
