package event

import (
	"github.com/raystack/optimus/core/resource"
	pbInt "github.com/raystack/optimus/protos/raystack/optimus/integration/v1beta1"
)

type ResourceCreated struct {
	Event

	Resource *resource.Resource
}

func NewResourceCreatedEvent(rsc *resource.Resource) (*ResourceCreated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResourceCreated) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type ResourceUpdated struct {
	Event

	Resource *resource.Resource
}

func NewResourceUpdatedEvent(rsc *resource.Resource) (*ResourceUpdated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResourceUpdated) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func resourceEventToBytes(event Event, rsc *resource.Resource, eventType pbInt.OptimusChangeEvent_EventType) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
