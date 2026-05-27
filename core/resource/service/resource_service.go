package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/event/moderator"
	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/writer"
)

type ResourceRepository interface {
	Create(ctx context.Context, res *resource.Resource) error
	Update(ctx context.Context, res *resource.Resource) error
	ChangeNamespace(ctx context.Context, res *resource.Resource, newTenant tenant.Tenant) error
	ReadByFullName(ctx context.Context, tnnt tenant.Tenant, store resource.Store, fullName string) (*resource.Resource, error)
	ReadAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Resource, error)
	GetResources(ctx context.Context, tnnt tenant.Tenant, store resource.Store, names []string) ([]*resource.Resource, error)
}

type ResourceManager interface {
	CreateResource(ctx context.Context, res *resource.Resource) error
	UpdateResource(ctx context.Context, res *resource.Resource) error
	SyncResource(ctx context.Context, res *resource.Resource) error
	BatchUpdate(ctx context.Context, store resource.Store, resources []*resource.Resource) error
	Validate(res *resource.Resource) error
	GetURN(res *resource.Resource) (string, error)
}

type DownstreamRefresher interface {
	RefreshResourceDownstream(ctx context.Context, resourceURNs []job.ResourceURN, logWriter writer.LogWriter) error
}

type TenantDetailsGetter interface {
	GetDetails(ctx context.Context, tnnt tenant.Tenant) (*tenant.WithDetails, error)
}

type EventHandler interface {
	HandleEvent(moderator.Event)
}

type ResourceService struct {
	repo      ResourceRepository
	mgr       ResourceManager
	refresher DownstreamRefresher

	logger       log.Logger
	eventHandler EventHandler
}

func NewResourceService(
	logger log.Logger,
	repo ResourceRepository, downstreamRefresher DownstreamRefresher, mgr ResourceManager,
	eventHandler EventHandler,
) *ResourceService {
	_ = "STUB: not implemented"
	return nil
}

func (rs ResourceService) Create(ctx context.Context, incoming *resource.Resource) error {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil
}

// Note: return in case resource already exists

func (rs ResourceService) Update(ctx context.Context, incoming *resource.Resource, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil
}

func (rs ResourceService) ChangeNamespace(ctx context.Context, datastore resource.Store, resourceFullName string, oldTenant, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil
}

func (rs ResourceService) Get(ctx context.Context, tnnt tenant.Tenant, store resource.Store, resourceFullName string) (*resource.Resource, error) {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil, nil
}

func (rs ResourceService) GetAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Resource, error) {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil, nil
}

func (rs ResourceService) SyncResources(ctx context.Context, tnnt tenant.Tenant, store resource.Store, names []string) (*resource.SyncResponse, error) {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil, nil
}

func (rs ResourceService) Deploy(ctx context.Context, tnnt tenant.Tenant, store resource.Store, incomings []*resource.Resource, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil
}

func (rs ResourceService) getResourcesToBatchUpdate(ctx context.Context, incomings []*resource.Resource, existingMappedByFullName map[string]*resource.Resource) ([]*resource.Resource, error) {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil, nil
}

func (rs ResourceService) raiseCreateEvent(res *resource.Resource) {
	_ = "STUB: not implemented" // nolint:gocritic
	return
}

func (rs ResourceService) raiseUpdateEvent(res *resource.Resource) {
	_ = "STUB: not implemented" // nolint:gocritic
	return
}

func (rs ResourceService) handleRefreshDownstream( // nolint:gocritic
	ctx context.Context,
	incomings []*resource.Resource,
	existingMappedByFullName map[string]*resource.Resource,
	logWriter writer.LogWriter,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (ResourceService) isToRefreshDownstream(incoming, existing *resource.Resource) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: this is not ideal solution, we need to see how to get these 'special' fields

func createFullNameToResourceMap(resources []*resource.Resource) map[string]*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}
