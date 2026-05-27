package resource

import (
	"github.com/raystack/optimus/core/tenant"
)

const (
	EntityResource       = "resource"
	nameSectionSeparator = "."
)

type Metadata struct {
	Version     int32
	Description string
	Labels      map[string]string
}

type Name string

func NameFrom(name string) (Name, error) { _ = "STUB: not implemented"; return *new(Name), nil }

func (n Name) String() string { _ = "STUB: not implemented"; return "" }

type Resource struct {
	name Name

	kind  string
	store Store
	urn   string

	tenant tenant.Tenant

	spec     map[string]any
	metadata *Metadata

	status Status
}

func NewResource(fullName, kind string, store Store, tnnt tenant.Tenant, meta *Metadata, spec map[string]any) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Resource) Name() Name { _ = "STUB: not implemented"; return *new(Name) }

func (r *Resource) FullName() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) URN() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) UpdateURN(urn string) error { _ = "STUB: not implemented"; return nil }

func (r *Resource) UpdateTenant(tnnt tenant.Tenant) { _ = "STUB: not implemented"; return }

func (r *Resource) Metadata() *Metadata { _ = "STUB: not implemented"; return nil }

func (r *Resource) NameSections() []string { _ = "STUB: not implemented"; return nil }

func (r *Resource) Kind() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) Tenant() tenant.Tenant { _ = "STUB: not implemented"; return *new(tenant.Tenant) }

func (r *Resource) Store() Store { _ = "STUB: not implemented"; return *new(Store) }

func (r *Resource) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (r *Resource) Spec() map[string]any { _ = "STUB: not implemented"; return nil }

func (r *Resource) Equal(incoming *Resource) bool { _ = "STUB: not implemented"; return false }

type FromExistingOpt func(r *Resource)

func ReplaceStatus(status Status) FromExistingOpt {
	_ = "STUB: not implemented"
	return *new(FromExistingOpt)
}

func FromExisting(existing *Resource, opts ...FromExistingOpt) *Resource {
	_ = "STUB: not implemented"
	return nil
}
