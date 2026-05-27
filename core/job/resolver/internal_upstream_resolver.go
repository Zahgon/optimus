package resolver

import (
	"golang.org/x/net/context"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
)

type internalUpstreamResolver struct {
	jobRepository JobRepository
}

func NewInternalUpstreamResolver(jobRepository JobRepository) *internalUpstreamResolver {
	_ = "STUB: not implemented"
	return nil
}

func (i internalUpstreamResolver) Resolve(ctx context.Context, jobWithUnresolvedUpstream *job.WithUpstream) (*job.WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i internalUpstreamResolver) BulkResolve(ctx context.Context, projectName tenant.ProjectName, jobsWithUnresolvedUpstream []*job.WithUpstream) ([]*job.WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i internalUpstreamResolver) resolveInferredUpstream(ctx context.Context, sources []job.ResourceURN) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i internalUpstreamResolver) resolveStaticUpstream(ctx context.Context, projectName tenant.ProjectName, upstreamSpec *job.UpstreamSpec) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
