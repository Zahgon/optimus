package bigquery

import (
	"context"
	"time"

	"github.com/raystack/optimus/core/resource"
)

const (
	backupTimePostfixFormat = "2006_01_02_15_04_05"

	configDataset        = "dataset"
	defaultBackupDataset = "optimus_backup"

	configPrefix        = "prefix"
	defaultBackupPrefix = "backup"

	configTTL        = "ttl"
	defaultBackupTTL = "720h"
)

func BackupResources(ctx context.Context, backup *resource.Backup, resources []*resource.Resource, client Client) (*resource.BackupResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateIfDatasetDoesNotExist(ctx context.Context, client Client, dataset Dataset) error {
	_ = "STUB: not implemented"
	return nil
}

func BackupTable(ctx context.Context, backup *resource.Backup, source *resource.Resource, client Client) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CopyTable(ctx context.Context, source, destination TableResourceHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func DestinationDataset(project string, backup *resource.Backup) (Dataset, error) {
	_ = "STUB: not implemented"
	return *new(Dataset), nil
}

func DestinationName(sourceDatasetName, sourceName string, backup *resource.Backup) string {
	_ = "STUB: not implemented"
	return ""
}

func DestinationExpiry(backup *resource.Backup) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
