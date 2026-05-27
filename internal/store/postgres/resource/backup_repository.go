package resource

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
)

const (
	backupToStoreColumns = `store, project_name, namespace_name, description, resource_names, config, created_at, updated_at`
	backupColumns        = `id, ` + backupToStoreColumns
)

type Backup struct {
	ID uuid.UUID

	Store         string
	ProjectName   string
	NamespaceName string

	Description   string
	ResourceNames pq.StringArray

	Config map[string]string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBackup(b *resource.Backup) Backup { _ = "STUB: not implemented"; return *new(Backup) }

func (b Backup) ToResourceBackup() (*resource.Backup, error) {
	_ = "STUB: not implemented" //nolint: gocritic
	return nil, nil
}

type BackupRepository struct {
	db *pgxpool.Pool
}

func (repo BackupRepository) Create(ctx context.Context, resourceBackup *resource.Backup) error {
	_ = "STUB: not implemented"
	return nil
}

func (repo BackupRepository) GetByID(ctx context.Context, id resource.BackupID) (*resource.Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (repo BackupRepository) GetAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBackupRepository(pool *pgxpool.Pool) *BackupRepository {
	_ = "STUB: not implemented"
	return nil
}
