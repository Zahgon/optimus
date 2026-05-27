package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"

	"github.com/raystack/optimus/core/resource"
)

const (
	locationKey        = "location"
	tableExpirationKey = "table_expiration"
)

type BqDataset interface {
	Create(context.Context, *bigquery.DatasetMetadata) error
	Update(context.Context, bigquery.DatasetMetadataToUpdate, string) (*bigquery.DatasetMetadata, error)
	Metadata(context.Context) (*bigquery.DatasetMetadata, error)
}

type DatasetHandle struct {
	bqDataset BqDataset
}

func (d DatasetHandle) Create(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DatasetHandle) Update(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DatasetHandle) Exists(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// There can be connection issue, we return false for now

func NewDatasetHandle(ds BqDataset) *DatasetHandle { _ = "STUB: not implemented"; return nil }

func toBQDatasetMetadata(details *DatasetDetails, res *resource.Resource) *bigquery.DatasetMetadata {
	_ = "STUB: not implemented"
	return nil
}

// structpb from proto returns a number value as float64

func ConfigAs[T any](mapping map[string]any, key string) T {
	_ = "STUB: not implemented"
	return *new(T)
}
