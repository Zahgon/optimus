package bucket

import (
	"context"
	"net/url"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/ext/scheduler/airflow"
)

const (
	scope = "https://www.googleapis.com/auth/cloud-platform"
)

func (f *Factory) GetGCSBucket(ctx context.Context, tnnt tenant.Tenant, parsedURL *url.URL) (airflow.Bucket, error) {
	_ = "STUB: not implemented"
	return *new(airflow.Bucket), nil
}
