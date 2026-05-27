package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type BqClientProvider struct{}

func NewClientProvider() *BqClientProvider { _ = "STUB: not implemented"; return nil }

func (BqClientProvider) Get(ctx context.Context, account string) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

type BqClient struct {
	bq *bigquery.Client
}

func NewClient(ctx context.Context, svcAccount string) (*BqClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *BqClient) DatasetHandleFrom(ds Dataset) ResourceHandle {
	_ = "STUB: not implemented"
	return *new(ResourceHandle)
}

func (c *BqClient) TableHandleFrom(ds Dataset, name string) TableResourceHandle {
	_ = "STUB: not implemented"
	return *new(TableResourceHandle)
}

func (c *BqClient) ExternalTableHandleFrom(ds Dataset, name string) ResourceHandle {
	_ = "STUB: not implemented"
	return *new(ResourceHandle)
}

func (c *BqClient) ViewHandleFrom(ds Dataset, name string) ResourceHandle {
	_ = "STUB: not implemented"
	return *new(ResourceHandle)
}

func (c *BqClient) Close() { _ = "STUB: not implemented"; return }
