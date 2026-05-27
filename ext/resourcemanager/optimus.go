package resourcemanager

import (
	"context"
	"net/http"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/core/job"
)

// ResourceManager is repository for external job spec
type ResourceManager interface {
	GetOptimusUpstreams(ctx context.Context, unresolvedDependency *job.Upstream) ([]*job.Upstream, error)
}

type OptimusResourceManager struct {
	name   string
	config config.ResourceManagerConfigOptimus

	httpClient *http.Client
}

// NewOptimusResourceManager initializes job spec repository for Optimus neighbor
func NewOptimusResourceManager(resourceManagerConfig config.ResourceManager) (*OptimusResourceManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OptimusResourceManager) GetOptimusUpstreams(ctx context.Context, unresolvedDependency *job.Upstream) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OptimusResourceManager) constructGetJobSpecificationsRequest(ctx context.Context, unresolvedDependency *job.Upstream) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OptimusResourceManager) toOptimusDependencies(responses []*jobSpecificationResponse, unresolvedDependency *job.Upstream) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OptimusResourceManager) toOptimusDependency(response *jobSpecificationResponse, unresolvedDependency *job.Upstream) (*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
