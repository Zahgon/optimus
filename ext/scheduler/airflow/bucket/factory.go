package bucket

import (
	"context"
	"net/url"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/ext/scheduler/airflow"
)

const (
	storagePathKey = "STORAGE_PATH"
)

type Factory struct {
	secretsGetter airflow.SecretGetter
	projectGetter airflow.ProjectGetter
}

func (f *Factory) New(ctx context.Context, tnnt tenant.Tenant) (airflow.Bucket, error) {
	_ = "STUB: not implemented"
	return *new(airflow.Bucket), nil
}

func (f *Factory) storageURL(ctx context.Context, tnnt tenant.Tenant) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFactory(projectGetter airflow.ProjectGetter, secretsGetter airflow.SecretGetter) *Factory {
	_ = "STUB: not implemented"
	return nil
}
