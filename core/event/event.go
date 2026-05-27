package event

import (
	"time"

	"github.com/google/uuid"
)

const eventsEntity = "events"

type Event struct {
	ID         uuid.UUID
	OccurredAt time.Time
}

func NewBaseEvent() (Event, error) { _ = "STUB: not implemented"; return *new(Event), nil }
