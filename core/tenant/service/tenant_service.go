package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/tenant"
)

type ProjectGetter interface {
	Get(context.Context, tenant.ProjectName) (*tenant.Project, error)
}

type NamespaceGetter interface {
	Get(context.Context, tenant.ProjectName, tenant.NamespaceName) (*tenant.Namespace, error)
}

type SecretsGetter interface {
	Get(ctx context.Context, projName tenant.ProjectName, namespaceName, name string) (*tenant.PlainTextSecret, error)
	GetAll(ctx context.Context, projName tenant.ProjectName, namespaceName string) ([]*tenant.PlainTextSecret, error)
}

type TenantService struct {
	projGetter      ProjectGetter
	namespaceGetter NamespaceGetter
	secretsGetter   SecretsGetter

	logger log.Logger
}

func (t TenantService) GetDetails(ctx context.Context, tnnt tenant.Tenant) (*tenant.WithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TenantService) GetProject(ctx context.Context, name tenant.ProjectName) (*tenant.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TenantService) GetSecrets(ctx context.Context, tnnt tenant.Tenant) ([]*tenant.PlainTextSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TenantService) GetSecret(ctx context.Context, tnnt tenant.Tenant, name string) (*tenant.PlainTextSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTenantService(projGetter ProjectGetter, nsGetter NamespaceGetter, secretsGetter SecretsGetter, logger log.Logger) *TenantService {
	_ = "STUB: not implemented"
	return nil
}
