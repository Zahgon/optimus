package tenant

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/tenant"
)

type NamespaceRepository struct {
	db *pgxpool.Pool
}

const (
	namespaceColumns = `id, name, config, project_name, created_at, updated_at`
)

type Namespace struct {
	ID     uuid.UUID
	Name   string
	Config map[string]string

	ProjectName string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n *Namespace) toTenantNamespace() (*tenant.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NamespaceRepository) Save(ctx context.Context, namespace *tenant.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamespaceRepository) GetByName(ctx context.Context, projectName tenant.ProjectName, name tenant.NamespaceName) (*tenant.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NamespaceRepository) get(ctx context.Context, projName tenant.ProjectName, name tenant.NamespaceName) (Namespace, error) {
	_ = "STUB: not implemented"
	return *new(Namespace), nil
}

func (n *NamespaceRepository) GetAll(ctx context.Context, projectName tenant.ProjectName) ([]*tenant.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNamespaceRepository(pool *pgxpool.Pool) *NamespaceRepository {
	_ = "STUB: not implemented"
	return nil
}
