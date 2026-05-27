package moderator

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/raystack/salt/log"
)

var eventQueueCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "publisher_events_created_total",
	Help: "Events created and to be sent to writer",
})

type Event interface {
	Bytes() ([]byte, error)
}

type Handler interface {
	HandleEvent(e Event)
}

type NoOpHandler struct{}

func (NoOpHandler) HandleEvent(_ Event) { _ = "STUB: not implemented"; return }

type EventHandler struct {
	messageChan chan<- []byte
	logger      log.Logger
}

func NewEventHandler(messageChan chan<- []byte, logger log.Logger) *EventHandler {
	_ = "STUB: not implemented"
	return nil
}

func (e EventHandler) HandleEvent(event Event) { _ = "STUB: not implemented"; return }
