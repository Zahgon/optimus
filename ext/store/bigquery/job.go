package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type BQJob interface {
	Wait(context.Context) (*bigquery.JobStatus, error)
}

type JobHandle struct {
	bqJob BQJob
}

func NewJob(job BQJob) *JobHandle { _ = "STUB: not implemented"; return nil }

func (j JobHandle) Wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
