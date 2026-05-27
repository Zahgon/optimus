package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/tenant"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type ProjectHandler struct {
	l              log.Logger
	projectService ProjectService

	pb.UnimplementedProjectServiceServer
}

type ProjectService interface {
	Save(context.Context, *tenant.Project) error
	Get(context.Context, tenant.ProjectName) (*tenant.Project, error)
	GetAll(context.Context) ([]*tenant.Project, error)
}

type TenantService interface {
	GetDetails(ctx context.Context, tnnt tenant.Tenant) (*tenant.WithDetails, error)
	GetSecrets(ctx context.Context, tnnt tenant.Tenant) ([]*tenant.PlainTextSecret, error)
	GetSecret(ctx context.Context, tnnt tenant.Tenant, name string) (*tenant.PlainTextSecret, error)
}

func (ph *ProjectHandler) RegisterProject(ctx context.Context, req *pb.RegisterProjectRequest) (*pb.RegisterProjectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO update the proto to remove the success & Message

func (ph *ProjectHandler) ListProjects(ctx context.Context, _ *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ph *ProjectHandler) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewProjectHandler(l log.Logger, projectService ProjectService) *ProjectHandler {
	_ = "STUB: not implemented"
	return nil
}

func fromProjectProto(conf *pb.ProjectSpecification) (*tenant.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toProjectProto(project *tenant.Project) *pb.ProjectSpecification {
	_ = "STUB: not implemented"
	return nil
}
