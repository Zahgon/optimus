package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"

	"github.com/raystack/optimus/core/resource"
)

const (
	expirationTimeKey = "expiration_time"

	skipLeadingRowsKey = "skip_leading_rows"
	rangeKey           = "range"
)

type ExternalTableHandle struct {
	bqExternalTable BqTable
}

func (et ExternalTableHandle) Create(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (et ExternalTableHandle) Update(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func (et ExternalTableHandle) Exists(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// There can be connection issue, we return false for now

func NewExternalTableHandle(bq BqTable) *ExternalTableHandle { _ = "STUB: not implemented"; return nil }

func bqExternalDataConfigTo(es *ExternalSource, schema Schema) (*bigquery.ExternalDataConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bqGoogleSheetsOptionsTo(m map[string]any) *bigquery.GoogleSheetsOptions {
	_ = "STUB: not implemented"
	return nil

	// grpc structpb.Struct cast numbers to float64
}

func toBQSchema(schema Schema) bigquery.Schema {
	_ = "STUB: not implemented"
	return *new(bigquery.Schema)
}
