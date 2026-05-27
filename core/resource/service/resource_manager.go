package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/resource"
)

type DataStore interface {
	Create(context.Context, *resource.Resource) error
	Update(context.Context, *resource.Resource) error
	BatchUpdate(context.Context, []*resource.Resource) error
	Validate(*resource.Resource) error
	GetURN(res *resource.Resource) (string, error)
	Backup(context.Context, *resource.Backup, []*resource.Resource) (*resource.BackupResult, error)
}

type ResourceStatusRepo interface {
	UpdateStatus(ctx context.Context, res ...*resource.Resource) error
}

type ResourceMgr struct {
	datastoreMap map[resource.Store]DataStore

	repo ResourceStatusRepo

	logger log.Logger
}

func (m *ResourceMgr) CreateResource(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ResourceMgr) UpdateResource(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ResourceMgr) SyncResource(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ResourceMgr) Validate(res *resource.Resource) error { _ = "STUB: not implemented"; return nil }

func (m *ResourceMgr) GetURN(res *resource.Resource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *ResourceMgr) BatchUpdate(ctx context.Context, store resource.Store, resources []*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ResourceMgr) Backup(ctx context.Context, details *resource.Backup, resources []*resource.Resource) (*resource.BackupResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ResourceMgr) RegisterDatastore(store resource.Store, dataStore DataStore) {
	_ = "STUB: not implemented"
	return
}

func NewResourceManager(repo ResourceStatusRepo, logger log.Logger) *ResourceMgr {
	_ = "STUB: not implemented"
	return nil
}
