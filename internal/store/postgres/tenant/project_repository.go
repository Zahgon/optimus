package tenant

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/tenant"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

const (
	projectColumns = `id, name, config, created_at, updated_at`
)

type Project struct {
	ID     uuid.UUID
	Name   string
	Config map[string]string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Project) toTenantProject() (*tenant.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo ProjectRepository) Save(ctx context.Context, tenantProject *tenant.Project) error {
	_ = "STUB: not implemented"
	return nil
}

func (repo ProjectRepository) GetByName(ctx context.Context, name tenant.ProjectName) (*tenant.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo ProjectRepository) get(ctx context.Context, name tenant.ProjectName) (Project, error) {
	_ = "STUB: not implemented"
	return *new(Project), nil
}

func (repo ProjectRepository) GetAll(ctx context.Context) ([]*tenant.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewProjectRepository(pool *pgxpool.Pool) *ProjectRepository {
	_ = "STUB: not implemented"
	return nil
}
