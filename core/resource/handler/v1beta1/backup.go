package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type BackupService interface {
	Create(context.Context, *resource.Backup) (*resource.BackupResult, error)
	Get(context.Context, resource.BackupID) (*resource.Backup, error)
	List(context.Context, tenant.Tenant, resource.Store) ([]*resource.Backup, error)
}

type BackupHandler struct {
	l       log.Logger
	service BackupService

	pb.UnimplementedBackupServiceServer
}

func (b BackupHandler) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*pb.CreateBackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b BackupHandler) ListBackups(ctx context.Context, req *pb.ListBackupsRequest) (*pb.ListBackupsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b BackupHandler) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.GetBackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toBackupSpec(detail *resource.Backup) *pb.BackupSpec { _ = "STUB: not implemented"; return nil }

func toIgnoredResources(ignored []resource.IgnoredResource) []*pb.IgnoredResource {
	_ = "STUB: not implemented"
	return nil
}

func NewBackupHandler(l log.Logger, service BackupService) *BackupHandler {
	_ = "STUB: not implemented"
	return nil
}
