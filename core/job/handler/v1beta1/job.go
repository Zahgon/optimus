package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/job/service/filter"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/writer"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	metricReplaceAllDuration = "job_replace_all_duration_seconds"
	metricRefreshDuration    = "job_refresh_duration_seconds"
	metricValidationDuration = "job_validation_duration_seconds"
)

type JobHandler struct {
	l          log.Logger
	jobService JobService

	pb.UnimplementedJobSpecificationServiceServer
}

func NewJobHandler(jobService JobService, logger log.Logger) *JobHandler {
	_ = "STUB: not implemented"
	return nil
}

type JobService interface {
	Add(ctx context.Context, jobTenant tenant.Tenant, jobs []*job.Spec) error
	Update(ctx context.Context, jobTenant tenant.Tenant, jobs []*job.Spec) error
	SyncState(ctx context.Context, jobTenant tenant.Tenant, disabledJobNames, enabledJobNames []job.Name) error
	UpdateState(ctx context.Context, jobTenant tenant.Tenant, jobNames []job.Name, jobState job.State, remark string) error
	ChangeNamespace(ctx context.Context, jobSourceTenant, jobNewTenant tenant.Tenant, jobName job.Name) error
	Delete(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name, cleanFlag, forceFlag bool) (affectedDownstream []job.FullName, err error)
	Get(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name) (jobSpec *job.Job, err error)
	GetTaskInfo(ctx context.Context, task job.Task) (*plugin.Info, error)
	GetByFilter(ctx context.Context, filters ...filter.FilterOpt) (jobSpecs []*job.Job, err error)
	ReplaceAll(ctx context.Context, jobTenant tenant.Tenant, jobs []*job.Spec, jobNamesWithInvalidSpec []job.Name, logWriter writer.LogWriter) error
	Refresh(ctx context.Context, projectName tenant.ProjectName, namespaceNames, jobNames []string, logWriter writer.LogWriter) error
	Validate(ctx context.Context, jobTenant tenant.Tenant, jobSpecs []*job.Spec, jobNamesWithInvalidSpec []job.Name, logWriter writer.LogWriter) error

	GetJobBasicInfo(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name, spec *job.Spec) (*job.Job, writer.BufferedLogger)
	GetUpstreamsToInspect(ctx context.Context, subjectJob *job.Job, localJob bool) ([]*job.Upstream, error)
	GetDownstream(ctx context.Context, job *job.Job, localJob bool) ([]*job.Downstream, error)
}

func (jh *JobHandler) AddJobSpecifications(ctx context.Context, jobSpecRequest *pb.AddJobSpecificationsRequest) (*pb.AddJobSpecificationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) DeleteJobSpecification(ctx context.Context, deleteRequest *pb.DeleteJobSpecificationRequest) (*pb.DeleteJobSpecificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) ChangeJobNamespace(ctx context.Context, changeRequest *pb.ChangeJobNamespaceRequest) (*pb.ChangeJobNamespaceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) UpdateJobSpecifications(ctx context.Context, jobSpecRequest *pb.UpdateJobSpecificationsRequest) (*pb.UpdateJobSpecificationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) GetJobSpecification(ctx context.Context, req *pb.GetJobSpecificationRequest) (*pb.GetJobSpecificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: return 404 if job is not found

func (jh *JobHandler) GetJobSpecifications(ctx context.Context, req *pb.GetJobSpecificationsRequest) (*pb.GetJobSpecificationsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: return 404 if job is not found

func (jh *JobHandler) ListJobSpecification(ctx context.Context, req *pb.ListJobSpecificationRequest) (*pb.ListJobSpecificationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: make a stream response

func (jh *JobHandler) GetWindow(_ context.Context, req *pb.GetWindowRequest) (*pb.GetWindowResponse, error) {
	_ = "STUB: not implemented"
	// TODO: the default version to be deprecated & made mandatory in future releases
	return nil, nil
}

func (jh *JobHandler) ReplaceAllJobSpecifications(stream pb.JobSpecificationService_ReplaceAllJobSpecificationsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (jh *JobHandler) RefreshJobs(request *pb.RefreshJobsRequest, stream pb.JobSpecificationService_RefreshJobsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (jh *JobHandler) CheckJobSpecifications(req *pb.CheckJobSpecificationsRequest, stream pb.JobSpecificationService_CheckJobSpecificationsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (jh *JobHandler) GetJobTask(ctx context.Context, req *pb.GetJobTaskRequest) (*pb.GetJobTaskResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: investigate destination type

func (jh *JobHandler) UpdateJobsState(ctx context.Context, req *pb.UpdateJobsStateRequest) (*pb.UpdateJobsStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) SyncJobsState(ctx context.Context, req *pb.SyncJobsStateRequest) (*pb.SyncJobsStateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jh *JobHandler) JobInspect(ctx context.Context, req *pb.JobInspectRequest) (*pb.JobInspectResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func raiseJobEventMetric(jobTenant tenant.Tenant, state string, metricValue int) {
	_ = "STUB: not implemented"
	return
}
