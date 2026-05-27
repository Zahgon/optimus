package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type JobRunService interface {
	JobRunInput(context.Context, tenant.ProjectName, scheduler.JobName, scheduler.RunConfig) (*scheduler.ExecutorInput, error)
	UpdateJobState(context.Context, *scheduler.Event) error
	GetJobRuns(ctx context.Context, projectName tenant.ProjectName, jobName scheduler.JobName, criteria *scheduler.JobRunsCriteria) ([]*scheduler.JobRunStatus, error)
	UploadToScheduler(ctx context.Context, projectName tenant.ProjectName) error
}

type Notifier interface {
	Push(ctx context.Context, event *scheduler.Event) error
}

type JobRunHandler struct {
	l        log.Logger
	service  JobRunService
	notifier Notifier

	pb.UnimplementedJobRunServiceServer
}

func (h JobRunHandler) JobRunInput(ctx context.Context, req *pb.JobRunInputRequest) (*pb.JobRunInputResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JobRun currently gets the job runs from scheduler based on the criteria
// TODO: later should collect the job runs from optimus
func (h JobRunHandler) JobRun(ctx context.Context, req *pb.JobRunRequest) (*pb.JobRunResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCriteriaForJobRun(req *pb.JobRunRequest) (*scheduler.JobRunsCriteria, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h JobRunHandler) UploadToScheduler(_ context.Context, req *pb.UploadToSchedulerRequest) (*pb.UploadToSchedulerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterJobEvent TODO: check in jaeger if this api takes time, then we can make this async
func (h JobRunHandler) RegisterJobEvent(ctx context.Context, req *pb.RegisterJobEventRequest) (*pb.RegisterJobEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobRunHandler(l log.Logger, service JobRunService, notifier Notifier) *JobRunHandler {
	_ = "STUB: not implemented"
	return nil
}
