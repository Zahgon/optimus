package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/writer"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	metricResourceEvents             = "resource_events_total"
	metricResourcesUploadAllDuration = "resource_upload_all_duration_seconds_total"
)

type ResourceService interface {
	Create(ctx context.Context, res *resource.Resource) error
	Update(ctx context.Context, res *resource.Resource, logWriter writer.LogWriter) error
	ChangeNamespace(ctx context.Context, datastore resource.Store, resourceFullName string, oldTenant, newTenant tenant.Tenant) error
	Get(ctx context.Context, tnnt tenant.Tenant, store resource.Store, resourceName string) (*resource.Resource, error)
	GetAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Resource, error)
	Deploy(ctx context.Context, tnnt tenant.Tenant, store resource.Store, resources []*resource.Resource, logWriter writer.LogWriter) error
	SyncResources(ctx context.Context, tnnt tenant.Tenant, store resource.Store, names []string) (*resource.SyncResponse, error)
}

type ResourceHandler struct {
	l       log.Logger
	service ResourceService

	pb.UnimplementedResourceServiceServer
}

func (rh ResourceHandler) DeployResourceSpecification(stream pb.ResourceService_DeployResourceSpecificationServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (rh ResourceHandler) ListResourceSpecification(ctx context.Context, req *pb.ListResourceSpecificationRequest) (*pb.ListResourceSpecificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh ResourceHandler) CreateResource(ctx context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh ResourceHandler) ReadResource(ctx context.Context, req *pb.ReadResourceRequest) (*pb.ReadResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh ResourceHandler) UpdateResource(ctx context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh ResourceHandler) ChangeResourceNamespace(ctx context.Context, req *pb.ChangeResourceNamespaceRequest) (*pb.ChangeResourceNamespaceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rh ResourceHandler) ApplyResources(ctx context.Context, req *pb.ApplyResourcesRequest) (*pb.ApplyResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeError(logWriter writer.LogWriter, err error) { _ = "STUB: not implemented"; return }

func writeResourcesStatus(resources []*resource.Resource, writeFn func(msg string)) {
	_ = "STUB: not implemented"
	return
}

func getResourcesByStatuses(resources []*resource.Resource, statuses ...resource.Status) []*resource.Resource {
	_ = "STUB: not implemented"
	return nil
}

func fromResourceProto(rs *pb.ResourceSpecification, tnnt tenant.Tenant, store resource.Store) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toResourceProto(res *resource.Resource) (*pb.ResourceSpecification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func raiseResourceDatastoreEventMetric(jobTenant tenant.Tenant, datastoreName, resourceKind, state string) {
	_ = "STUB: not implemented"
	return
}

func NewResourceHandler(l log.Logger, resourceService ResourceService) *ResourceHandler {
	_ = "STUB: not implemented"
	return nil
}
