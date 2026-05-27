package resource

import (
	"time"

	"github.com/google/uuid"

	"github.com/raystack/optimus/core/tenant"
)

const (
	EntityBackup = "backup"
)

type BackupID uuid.UUID

func BackupIDFrom(id string) (BackupID, error) {
	_ = "STUB: not implemented"
	return *new(BackupID), nil
}

func (i BackupID) String() string { _ = "STUB: not implemented"; return "" }

func (i BackupID) IsInvalid() bool { _ = "STUB: not implemented"; return false }

func (i BackupID) UUID() uuid.UUID { _ = "STUB: not implemented"; return *new(uuid.UUID) }

type IgnoredResource struct {
	Name   string
	Reason string
}

type BackupResult struct {
	ID               BackupID
	ResourceNames    []string
	IgnoredResources []IgnoredResource
}

type Backup struct {
	id BackupID

	store  Store
	tenant tenant.Tenant

	resourceNames []string
	description   string
	createdAt     time.Time
	config        map[string]string
}

func NewBackup(store Store, t tenant.Tenant, resNames []string, desc string, createdAt time.Time, conf map[string]string) (*Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backup) GetConfigOrDefaultFor(key, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *Backup) UpdateID(id uuid.UUID) error { _ = "STUB: not implemented"; return nil }

func (b *Backup) ID() BackupID { _ = "STUB: not implemented"; return *new(BackupID) }

func (b *Backup) Store() Store { _ = "STUB: not implemented"; return *new(Store) }

func (b *Backup) Tenant() tenant.Tenant { _ = "STUB: not implemented"; return *new(tenant.Tenant) }

func (b *Backup) ResourceNames() []string { _ = "STUB: not implemented"; return nil }

func (b *Backup) Description() string { _ = "STUB: not implemented"; return "" }

func (b *Backup) CreatedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *Backup) Config() map[string]string { _ = "STUB: not implemented"; return nil }
