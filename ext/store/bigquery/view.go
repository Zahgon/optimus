package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"

	"github.com/raystack/optimus/core/resource"
)

type ViewHandle struct {
	bqView BqTable
}

func (v ViewHandle) Create(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (v ViewHandle) Update(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (v ViewHandle) Exists(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// There can be connection issue, we return false for now

func NewViewHandle(bq BqTable) *ViewHandle { _ = "STUB: not implemented"; return nil }

func getMetadataToCreate(desc string, extraConf map[string]any, labels map[string]string) (*bigquery.TableMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMetadataToUpdate(description string, extraConf map[string]any, labels map[string]string) (bigquery.TableMetadataToUpdate, error) {
	_ = "STUB: not implemented"
	return *new(bigquery.TableMetadataToUpdate), nil
}
