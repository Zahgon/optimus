package resolver

import (
	"context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

const (
	// maxPriorityWeight - is the maximus weight a DAG will be given.
	maxPriorityWeight = 10000

	// priorityWeightGap - while giving weights to the DAG, what's the GAP
	// do we want to consider. PriorityWeightGap = 1 means, weights will be 1, 2, 3 etc.
	priorityWeightGap = 10
)

type SimpleResolver struct{}

func NewSimpleResolver() *SimpleResolver { _ = "STUB: not implemented"; return nil }

func (SimpleResolver) Resolve(_ context.Context, details []*scheduler.JobWithDetails) error {
	_ = "STUB: not implemented" // nolint:unparam
	return nil
}

func numberOfUpstreams(upstream scheduler.Upstreams, tnnt tenant.Tenant) int {
	_ = "STUB: not implemented"
	return 0
}
