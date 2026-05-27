package resolver

import (
	"context"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/ext/resourcemanager"
	"github.com/raystack/optimus/internal/writer"
)

type extUpstreamResolver struct {
	optimusResourceManagers []resourcemanager.ResourceManager
}

// NewExternalUpstreamResolver creates a new instance of externalUpstreamResolver
func NewExternalUpstreamResolver(resourceManagerConfigs []config.ResourceManager) (*extUpstreamResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResourceManager interface {
	GetOptimusUpstreams(ctx context.Context, unresolvedDependency *job.Upstream) ([]*job.Upstream, error)
}

func (e *extUpstreamResolver) Resolve(ctx context.Context, jobWithUpstream *job.WithUpstream, lw writer.LogWriter) (*job.WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *extUpstreamResolver) BulkResolve(ctx context.Context, jobsWithUpstream []*job.WithUpstream, lw writer.LogWriter) ([]*job.WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *extUpstreamResolver) fetchOptimusUpstreams(ctx context.Context, unresolvedUpstream *job.Upstream) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTestExternalUpstreamResolver(
	optimusResourceManagers []resourcemanager.ResourceManager,
) ExternalUpstreamResolver {
	_ = "STUB: not implemented"
	return *new(ExternalUpstreamResolver)
}
