package bigquery

import (
	"context"
	"time"

	bq "cloud.google.com/go/bigquery"
	"go.opentelemetry.io/otel/trace"

	"github.com/raystack/optimus/core/resource"
	"github.com/raystack/optimus/core/tenant"
)

const (
	accountKey = "DATASTORE_BIGQUERY"
	store      = "BigqueryStore"

	ConcurrentTicketPerSec = 5
	ConcurrentLimit        = 20
)

type ResourceHandle interface {
	Create(ctx context.Context, res *resource.Resource) error
	Update(ctx context.Context, res *resource.Resource) error
	Exists(ctx context.Context) bool
}

type TableResourceHandle interface {
	ResourceHandle
	GetBQTable() (*bq.Table, error)
	CopierFrom(source TableResourceHandle) (TableCopier, error)
	UpdateExpiry(ctx context.Context, name string, expiry time.Time) error
}

type Client interface {
	DatasetHandleFrom(dataset Dataset) ResourceHandle
	TableHandleFrom(dataset Dataset, name string) TableResourceHandle
	ExternalTableHandleFrom(dataset Dataset, name string) ResourceHandle
	ViewHandleFrom(dataset Dataset, name string) ResourceHandle
	Close()
}

type ClientProvider interface {
	Get(ctx context.Context, account string) (Client, error)
}

type SecretProvider interface {
	GetSecret(ctx context.Context, tnnt tenant.Tenant, key string) (*tenant.PlainTextSecret, error)
}

type Store struct {
	secretProvider SecretProvider
	clientProvider ClientProvider
}

func (s Store) Create(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Store) Update(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (s Store) BatchUpdate(ctx context.Context, resources []*resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (Store) Validate(r *resource.Resource) error { _ = "STUB: not implemented"; return nil }

func (Store) GetURN(res *resource.Resource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s Store) Backup(ctx context.Context, backup *resource.Backup, resources []*resource.Resource) (*resource.BackupResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func startChildSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func NewBigqueryDataStore(secretProvider SecretProvider, clientProvider ClientProvider) *Store {
	_ = "STUB: not implemented"
	return nil
}
