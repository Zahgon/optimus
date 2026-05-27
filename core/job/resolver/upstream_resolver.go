package resolver

import (
	"context"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/writer"
)

const (
	ConcurrentTicketPerSec = 50
	ConcurrentLimit        = 100
)

type UpstreamResolver struct {
	jobRepository            JobRepository
	externalUpstreamResolver ExternalUpstreamResolver
	internalUpstreamResolver InternalUpstreamResolver
}

func NewUpstreamResolver(jobRepository JobRepository, externalUpstreamResolver ExternalUpstreamResolver, internalUpstreamResolver InternalUpstreamResolver) *UpstreamResolver {
	_ = "STUB: not implemented"
	return nil
}

type ExternalUpstreamResolver interface {
	Resolve(ctx context.Context, jobWithUpstream *job.WithUpstream, lw writer.LogWriter) (*job.WithUpstream, error)
	BulkResolve(context.Context, []*job.WithUpstream, writer.LogWriter) ([]*job.WithUpstream, error)
}

type InternalUpstreamResolver interface {
	Resolve(context.Context, *job.WithUpstream) (*job.WithUpstream, error)
	BulkResolve(context.Context, tenant.ProjectName, []*job.WithUpstream) ([]*job.WithUpstream, error)
}

type JobRepository interface {
	ResolveUpstreams(ctx context.Context, projectName tenant.ProjectName, jobNames []job.Name) (map[job.Name][]*job.Upstream, error)

	GetAllByResourceDestination(ctx context.Context, resourceDestination job.ResourceURN) ([]*job.Job, error)
	GetByJobName(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) (*job.Job, error)
}

func (u UpstreamResolver) BulkResolve(ctx context.Context, projectName tenant.ProjectName, jobs []*job.Job, logWriter writer.LogWriter) ([]*job.WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u UpstreamResolver) Resolve(ctx context.Context, subjectJob *job.Job, logWriter writer.LogWriter) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (UpstreamResolver) getUnresolvedUpstreamsErrors(jobsWithUpstreams []*job.WithUpstream, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented"
	return nil
}
