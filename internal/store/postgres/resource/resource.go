package resource

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/raystack/optimus/core/resource"
)

type Resource struct {
	ID uuid.UUID

	FullName string
	Kind     string
	Store    string

	ProjectName   string
	NamespaceName string

	Metadata json.RawMessage
	Spec     map[string]any

	URN string

	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromResourceToModel(r *resource.Resource) *Resource { _ = "STUB: not implemented"; return nil }

func FromModelToResource(r *Resource) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
