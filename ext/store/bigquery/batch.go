package bigquery

import (
	"context"

	"github.com/kushsharma/parallel"

	"github.com/raystack/optimus/core/resource"
)

type Batch struct {
	Dataset        Dataset
	DatasetDetails *resource.Resource

	provider ClientProvider

	Tables         []*resource.Resource
	ExternalTables []*resource.Resource
	Views          []*resource.Resource
}

func (b *Batch) QueueJobs(ctx context.Context, account string, runner *parallel.Runner) error {
	_ = "STUB: not implemented"
	return nil
}

func createOrUpdate(ctx context.Context, handle ResourceHandle, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func create(ctx context.Context, handle ResourceHandle, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func update(ctx context.Context, handle ResourceHandle, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batch) DatasetOrDefault() (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BatchesFrom(resources []*resource.Resource, provider ClientProvider) (map[string]*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
