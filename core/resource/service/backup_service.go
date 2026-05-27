package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
)

const (
	// recentBackupWindowMonths contains the window interval to consider for recent backups
	recentBackupWindowMonths = -3

	metricBackupRequest        = "resource_backup_requests_total"
	backupRequestStatusSuccess = "success"
	backupRequestStatusFailed  = "failed"
)

type BackupRepository interface {
	GetByID(ctx context.Context, id resource.BackupID) (*resource.Backup, error)
	GetAll(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Backup, error)
	Create(ctx context.Context, backup *resource.Backup) error
}

type ResourceProvider interface {
	GetResources(ctx context.Context, tnnt tenant.Tenant, store resource.Store, names []string) ([]*resource.Resource, error)
}

type BackupManager interface {
	Backup(ctx context.Context, backup *resource.Backup, resources []*resource.Resource) (*resource.BackupResult, error)
}

type BackupService struct {
	repo BackupRepository

	resources     ResourceProvider
	backupManager BackupManager

	logger log.Logger
}

func (s BackupService) Create(ctx context.Context, backup *resource.Backup) (*resource.BackupResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s BackupService) Get(ctx context.Context, backupID resource.BackupID) (*resource.Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s BackupService) List(ctx context.Context, tnnt tenant.Tenant, store resource.Store) ([]*resource.Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findMissingResources(names []string, resources []*resource.Resource) []resource.IgnoredResource {
	_ = "STUB: not implemented"
	return nil
}

func NewBackupService(repo BackupRepository, resources ResourceProvider, manager BackupManager, logger log.Logger) *BackupService {
	_ = "STUB: not implemented"
	return nil
}

func raiseBackupRequestMetrics(jobTenant tenant.Tenant, backupResult *resource.BackupResult) {
	_ = "STUB: not implemented"
	return
}

func raiseBackupRequestMetric(jobTenant tenant.Tenant, resourceName, state string) {
	_ = "STUB: not implemented"
	return
}
