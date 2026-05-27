package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/tenant"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type NamespaceService interface {
	Save(ctx context.Context, namespace *tenant.Namespace) error
	Get(context.Context, tenant.ProjectName, tenant.NamespaceName) (*tenant.Namespace, error)
	GetAll(context.Context, tenant.ProjectName) ([]*tenant.Namespace, error)
}

type NamespaceHandler struct {
	l         log.Logger
	nsService NamespaceService

	pb.UnimplementedNamespaceServiceServer
}

func (nh *NamespaceHandler) RegisterProjectNamespace(ctx context.Context, req *pb.RegisterProjectNamespaceRequest) (
	*pb.RegisterProjectNamespaceResponse, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nh *NamespaceHandler) ListProjectNamespaces(ctx context.Context, req *pb.ListProjectNamespacesRequest) (
	*pb.ListProjectNamespacesResponse, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nh *NamespaceHandler) GetNamespace(ctx context.Context, request *pb.GetNamespaceRequest) (
	*pb.GetNamespaceResponse, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNamespaceHandler(l log.Logger, nsService NamespaceService) *NamespaceHandler {
	_ = "STUB: not implemented"
	return nil
}

func fromNamespaceProto(conf *pb.NamespaceSpecification, projName tenant.ProjectName) (*tenant.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toNamespaceProto(ns *tenant.Namespace) *pb.NamespaceSpecification {
	_ = "STUB: not implemented"
	return nil
}
