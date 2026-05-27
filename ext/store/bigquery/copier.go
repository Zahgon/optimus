package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type BQCopier interface {
	Run(context.Context) (*bigquery.Job, error)
}

type CopyJob interface {
	Wait(ctx context.Context) error
}

type Copier struct {
	bqCopier BQCopier
}

func NewCopier(bqCopier BQCopier) *Copier { _ = "STUB: not implemented"; return nil }

func (c Copier) Run(ctx context.Context) (CopyJob, error) {
	_ = "STUB: not implemented"
	return *new(CopyJob), nil
}
