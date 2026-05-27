package service

import (
	"context"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

func (s *JobRunService) UploadToScheduler(ctx context.Context, projectName tenant.ProjectName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) deployJobsPerNamespace(ctx context.Context, t tenant.Tenant, jobs []*scheduler.JobWithDetails) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) cleanPerNamespace(ctx context.Context, t tenant.Tenant, jobs []*scheduler.JobWithDetails) error {
	_ = "STUB: not implemented"
	// get all stored job names
	return nil
}

func (s *JobRunService) UploadJobs(ctx context.Context, tnnt tenant.Tenant, toUpdate, toDelete []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *JobRunService) resolveAndDeployJobs(ctx context.Context, tnnt tenant.Tenant, toUpdate []string) error {
	_ = "STUB: not implemented"
	return nil
}
