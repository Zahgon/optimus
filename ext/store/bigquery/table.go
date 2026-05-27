package bigquery

import (
	"context"
	"time"

	"cloud.google.com/go/bigquery"

	"github.com/raystack/optimus/core/resource"
)

type BqTable interface {
	Create(context.Context, *bigquery.TableMetadata) error
	Update(context.Context, bigquery.TableMetadataToUpdate, string, ...bigquery.TableUpdateOption) (*bigquery.TableMetadata, error)
	Metadata(ctx context.Context, opts ...bigquery.TableMetadataOption) (*bigquery.TableMetadata, error)
	CopierFrom(srcs ...*bigquery.Table) *bigquery.Copier
}

type TableCopier interface {
	Run(ctx context.Context) (CopyJob, error)
}

type TableHandle struct {
	bqTable BqTable
}

func (t TableHandle) Create(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TableHandle) Update(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// updating range based partition after creation is not supported

func (t TableHandle) Exists(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// There can be connection issue, we return false for now

func (t TableHandle) CopierFrom(source TableResourceHandle) (TableCopier, error) {
	_ = "STUB: not implemented"
	return *new(TableCopier), nil
}

func (t TableHandle) UpdateExpiry(ctx context.Context, name string, expiry time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (t TableHandle) GetBQTable() (*bigquery.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTableHandle(bq BqTable) *TableHandle { _ = "STUB: not implemented"; return nil }

func toBQTableMetadata(t *Table, res *resource.Resource) (*bigquery.TableMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toBQRangePartitioning(t *Partition) *bigquery.RangePartitioning {
	_ = "STUB: not implemented"
	return nil
}

func toBQTimePartitioning(t *Partition) *bigquery.TimePartitioning {
	_ = "STUB: not implemented"
	return nil
}

func toBQClustering(ct *Cluster) *bigquery.Clustering { _ = "STUB: not implemented"; return nil }
