package tenant

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/core/tenant/dto"
)

type SecretRepository struct {
	db *pgxpool.Pool
}

const (
	secretColumns = `id, name, value, project_name, namespace_name, created_at, updated_at`

	getAllSecretsInProject = `SELECT ` + secretColumns + `
FROM secret s WHERE project_name = $1`
)

type Secret struct {
	ID uuid.UUID

	Name  string
	Value string

	ProjectName   string
	NamespaceName sql.NullString

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSecret(secret *tenant.Secret) Secret {
	_ = "STUB: not implemented"
	// base64 for storing safely in db
	return *new(Secret)
}

func (s *Secret) ToTenantSecret() (*tenant.Secret, error) {
	_ = "STUB: not implemented"
	// decode base64
	return nil, nil
}

func (s *Secret) ToSecretInfo() (*dto.SecretInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s SecretRepository) Save(ctx context.Context, tenantSecret *tenant.Secret) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretRepository) Update(ctx context.Context, tenantSecret *tenant.Secret) error {
	_ = "STUB: not implemented"
	return nil
}

// Get is scoped to the tenant provided in the argument
func (s SecretRepository) Get(ctx context.Context, projName tenant.ProjectName, nsName string, name tenant.SecretName) (*tenant.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get is scoped only at project level, used for db operations
func (s SecretRepository) get(ctx context.Context, projName tenant.ProjectName, name tenant.SecretName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretRepository) GetAll(ctx context.Context, projName tenant.ProjectName, nsName string) ([]*tenant.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete will not support soft delete, once deleted it has to be created again
func (s SecretRepository) Delete(ctx context.Context, projName tenant.ProjectName, nsName string, name tenant.SecretName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretRepository) GetSecretsInfo(ctx context.Context, projName tenant.ProjectName) ([]*dto.SecretInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSecretRepository(pool *pgxpool.Pool) *SecretRepository {
	_ = "STUB: not implemented"
	return nil
}
