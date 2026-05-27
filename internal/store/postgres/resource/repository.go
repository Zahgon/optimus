package resource

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
)

const (
	columnsToStore  = `full_name, kind, store, status, urn, project_name, namespace_name, metadata, spec, created_at, updated_at`
	resourceColumns = `id, ` + columnsToStore
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository { _ = "STUB: not implemented"; return nil }

func (r Repository) Create(ctx context.Context, resourceModel *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (r Repository) Update(ctx context.Context, resourceModel *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (r Repository) ChangeNamespace(ctx context.Context, res *resource.Resource, newTenant tenant.Tenant) error {
	_ = "STUB: not implemented"
	return nil
}

func (r Repository) ReadByFullName(ctx context.Context, tnnt tenant.Tenant, store resource.Store, fullName string) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Repository) ReadAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Repository) GetResources(ctx context.Context, tnnt tenant.Tenant, store resource.Store, names []string) ([]*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Repository) UpdateStatus(ctx context.Context, resources ...*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}
