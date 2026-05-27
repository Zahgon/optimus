package slack

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	api "github.com/slack-go/slack"

	"github.com/raystack/optimus/core/scheduler"
)

const (
	DefaultEventBatchInterval = time.Second * 10
	MaxSLAEventsToProcess     = 6
)

var (
	notifierType      = "slack"
	slackQueueCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationQueue,
		ConstLabels: map[string]string{"type": notifierType},
	})
	slackWorkerBatchCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationWorkerBatch,
		ConstLabels: map[string]string{"type": notifierType},
	})
	slackWorkerSendErrCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name:        scheduler.MetricNotificationWorkerSendErr,
		ConstLabels: map[string]string{"type": notifierType},
	})
)

type Notifier struct {
	io.Closer

	slackURL      string
	routeMsgBatch map[route][]event // channelID -> [][][][][]
	wg            sync.WaitGroup
	mu            sync.Mutex
	workerErrChan chan error

	eventBatchInterval time.Duration
}

type route struct {
	receiverID string
	authToken  string
}

type event struct {
	authToken string
	owner     string
	meta      *scheduler.Event
}

func (s *Notifier) Notify(ctx context.Context, attr scheduler.NotifyAttrs) error {
	_ = "STUB: not implemented" //nolint: gocritic
	return nil
}

// channel

// user

// user group

// user email

// fail if unable to find the receiver ID

func (s *Notifier) queueNotification(receiverIDs []string, oauthSecret string, attr scheduler.NotifyAttrs) {
	_ = "STUB: not implemented" //nolint: gocritic
	return
}

// accumulate messages
func buildMessageBlocks(events []event, workerErrChan chan error) []api.Block {
	_ = "STUB: not implemented"
	return nil

	// core details related to event
}

//nolint: gocritic

// skip further SLA events

// event log url button

// event job url button

// build context footer

// api.NewOptionBlockObject("", optionText, nil))

// Build context section

func (s *Notifier) Worker(ctx context.Context) { _ = "STUB: not implemented"; return }

// iterate over all queued routeMsgBatch and

//nolint: gocritic

// clear events from map as they are processed

// send messages in batches of 5 secs

func (s *Notifier) Close() error {
	_ = "STUB: not implemented" // nolint: unparam
	// drain batches
	return nil
}

func NewNotifier(ctx context.Context, slackURL string, eventBatchInterval time.Duration, errHandler func(error)) *Notifier {
	_ = "STUB: not implemented"
	return nil
}
