package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/event/moderator"
	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/job/service/filter"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/errors"
	"github.com/raystack/optimus/internal/lib/tree"
	"github.com/raystack/optimus/internal/writer"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	ConcurrentTicketPerSec = 50
	ConcurrentLimit        = 100
)

type JobService struct {
	jobRepo        JobRepository
	upstreamRepo   UpstreamRepository
	downstreamRepo DownstreamRepository

	pluginService    PluginService
	upstreamResolver UpstreamResolver
	eventHandler     EventHandler
	scheduler        Scheduler

	tenantDetailsGetter TenantDetailsGetter

	jobDeploymentService JobDeploymentService

	logger log.Logger
}

func NewJobService(
	jobRepo JobRepository, upstreamRepo UpstreamRepository, downstreamRepo DownstreamRepository,
	pluginService PluginService, upstreamResolver UpstreamResolver,
	tenantDetailsGetter TenantDetailsGetter, eventHandler EventHandler, logger log.Logger,
	jobDeploymentService JobDeploymentService, scheduler Scheduler,
) *JobService {
	_ = "STUB: not implemented"
	return nil
}

type PluginService interface {
	Info(context.Context, job.TaskName) (*plugin.Info, error)
	GenerateDestination(context.Context, *tenant.WithDetails, job.Task) (job.ResourceURN, error)
	GenerateUpstreams(ctx context.Context, jobTenant *tenant.WithDetails, spec *job.Spec, dryRun bool) ([]job.ResourceURN, error)
}

type TenantDetailsGetter interface {
	GetDetails(ctx context.Context, jobTenant tenant.Tenant) (*tenant.WithDetails, error)
}

type JobDeploymentService interface {
	UploadJobs(ctx context.Context, jobTenant tenant.Tenant, toUpdate, toDelete []string) error
}

type JobRepository interface {
	// TODO: remove `savedJobs` since the method's main purpose is to add, not to get
	Add(context.Context, []*job.Job) (addedJobs []*job.Job, err error)
	Update(context.Context, []*job.Job) (updatedJobs []*job.Job, err error)
	Delete(ctx context.Context, projectName tenant.ProjectName, jobName job.Name, cleanHistory bool) error

	ChangeJobNamespace(ctx context.Context, jobName job.Name, tenant, newTenant tenant.Tenant) error

	GetByJobName(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) (*job.Job, error)
	GetAllByResourceDestination(ctx context.Context, resourceDestination job.ResourceURN) ([]*job.Job, error)
	GetAllByTenant(ctx context.Context, jobTenant tenant.Tenant) ([]*job.Job, error)
	GetAllByProjectName(ctx context.Context, projectName tenant.ProjectName) ([]*job.Job, error)
	SyncState(ctx context.Context, jobTenant tenant.Tenant, disabledJobNames, enabledJobNames []job.Name) error
	UpdateState(ctx context.Context, jobTenant tenant.Tenant, jobNames []job.Name, jobState job.State, remark string) error
}

type UpstreamRepository interface {
	ResolveUpstreams(context.Context, tenant.ProjectName, []job.Name) (map[job.Name][]*job.Upstream, error)
	ReplaceUpstreams(context.Context, []*job.WithUpstream) error
	GetUpstreams(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) ([]*job.Upstream, error)
}

type DownstreamRepository interface {
	GetDownstreamByDestination(ctx context.Context, projectName tenant.ProjectName, destination job.ResourceURN) ([]*job.Downstream, error)
	GetDownstreamByJobName(ctx context.Context, projectName tenant.ProjectName, jobName job.Name) ([]*job.Downstream, error)
	GetDownstreamBySources(ctx context.Context, sources []job.ResourceURN) ([]*job.Downstream, error)
}

type EventHandler interface {
	HandleEvent(moderator.Event)
}

type UpstreamResolver interface {
	BulkResolve(ctx context.Context, projectName tenant.ProjectName, jobs []*job.Job, logWriter writer.LogWriter) (jobWithUpstreams []*job.WithUpstream, err error)
	Resolve(ctx context.Context, subjectJob *job.Job, logWriter writer.LogWriter) ([]*job.Upstream, error)
}

type Scheduler interface {
	UpdateJobState(ctx context.Context, tnnt tenant.Tenant, jobName []job.Name, state string) error
}

func (j *JobService) Add(ctx context.Context, jobTenant tenant.Tenant, specs []*job.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) Update(ctx context.Context, jobTenant tenant.Tenant, specs []*job.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) UpdateState(ctx context.Context, jobTenant tenant.Tenant, jobNames []job.Name, jobState job.State, remark string) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) SyncState(ctx context.Context, jobTenant tenant.Tenant, disabledJobNames, enabledJobNames []job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) Delete(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name, cleanFlag, forceFlag bool) (affectedDownstream []job.FullName, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) ChangeNamespace(ctx context.Context, jobTenant, jobNewTenant tenant.Tenant, jobName job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) Get(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name) (*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) GetTaskInfo(ctx context.Context, task job.Task) (*plugin.Info, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) GetByFilter(ctx context.Context, filters ...filter.FilterOpt) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when resource destination exist, filter by destination

// when project name and job names exist, filter by project and job names

// when project name and job name exist, filter by project name and job name

// when project name and namespace names exist, filter by tenant

// when project name and namespace name exist, filter by tenant

// when project name exist, filter by project name

func (j *JobService) ReplaceAll(ctx context.Context, jobTenant tenant.Tenant, specs []*job.Spec, jobNamesWithInvalidSpec []job.Name, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) uploadJobs(ctx context.Context, jobTenant tenant.Tenant, addedJobs, updatedJobs []*job.Job, deletedJobNames []job.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) Refresh(ctx context.Context, projectName tenant.ProjectName, namespaceNames, jobNames []string, logWriter writer.LogWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) RefreshResourceDownstream(ctx context.Context, resourceURNs []job.ResourceURN, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) Validate(ctx context.Context, jobTenant tenant.Tenant, jobSpecs []*job.Spec, jobNamesWithInvalidSpec []job.Name, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: only check cyclic deps across internal upstreams (sources), need further discussion to check cyclic deps for external upstream
// assumption, all job specs from input are also the job within same project

func (j *JobService) validateDeleteJobs(ctx context.Context, jobTenant tenant.Tenant, toDelete []*job.Spec, logWriter writer.LogWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeleteJob(jobTenant tenant.Tenant, downstreams []*job.Downstream, toDeleteMap map[job.FullName]*job.Spec, jobToDelete *job.Spec, logWriter writer.LogWriter, me *errors.MultiError) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: refactor to put the log writer outside

func isJobSafeToDelete(toDeleteMap map[job.FullName]*job.Spec, downstreamFullNames []job.FullName) ([]job.FullName, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (j *JobService) getAllDownstreams(ctx context.Context, projectName tenant.ProjectName, jobName job.Name, visited map[job.FullName]bool) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAllJobsToValidateMap(incomingJobs, existingJobs []*job.Job, unmodifiedSpecs []*job.Spec) map[job.Name]*job.WithUpstream {
	_ = "STUB: not implemented"
	// TODO: check whether we need to accumulate encountered errors
	return nil
}

func getIdentifierToJobsMap(jobsToValidateMap map[job.Name]*job.WithUpstream) map[string][]*job.WithUpstream {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) resolveAndSaveUpstreams(ctx context.Context, jobTenant tenant.Tenant, logWriter writer.LogWriter, jobsToResolve ...[]*job.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) bulkAdd(ctx context.Context, tenantWithDetails *tenant.WithDetails, specsToAdd []*job.Spec, logWriter writer.LogWriter) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: consider do add inside parallel

func (j *JobService) bulkUpdate(ctx context.Context, tenantWithDetails *tenant.WithDetails, specsToUpdate []*job.Spec, logWriter writer.LogWriter) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) bulkDelete(ctx context.Context, jobTenant tenant.Tenant, toDelete []*job.Spec, logWriter writer.LogWriter) ([]job.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: reuse Delete method and pass forceFlag as false

func (*JobService) differentiateSpecs(existingJobs []*job.Job, specs []*job.Spec, jobNamesWithInvalidSpec []job.Name) (added, modified, deleted, unmodified []*job.Spec, err error) {
	_ = "STUB: not implemented"
	// TODO: consider checking multi-error if it is required here
	return nil, nil, nil, nil, nil
}

func (j *JobService) generateJobs(ctx context.Context, tenantWithDetails *tenant.WithDetails, specs []*job.Spec, logWriter writer.LogWriter) ([]*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) generateJob(ctx context.Context, tenantWithDetails *tenant.WithDetails, spec *job.Spec) (*job.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) validateCyclic(rootName job.Name, jobMap map[job.Name]*job.WithUpstream, identifierToJobMap map[string][]*job.WithUpstream) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*JobService) buildDAGTree(rootName job.Name, jobMap map[job.Name]*job.WithUpstream, identifierToJobMap map[string][]*job.WithUpstream) *tree.MultiRootTree {
	_ = "STUB: not implemented"
	return nil
}

// resource maybe from external optimus or outside project,
// as of now, we're not providing the capability to build tree from external optimus or outside project. skip

// sources: https://github.com/raystack/optimus/blob/a6dafbc1fbeb8e1f1eb8d4a6e9582ada4a7f639e/job/replay.go#L101
func findOrCreateDAGNode(dagTree *tree.MultiRootTree, dag tree.TreeData) *tree.TreeNode {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobService) GetJobBasicInfo(ctx context.Context, jobTenant tenant.Tenant, jobName job.Name, spec *job.Spec) (*job.Job, writer.BufferedLogger) {
	_ = "STUB: not implemented"
	return nil, *new(writer.BufferedLogger)
}

func (j *JobService) getJobNamesWithSameDestination(ctx context.Context, subjectJob *job.Job) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (j *JobService) GetUpstreamsToInspect(ctx context.Context, subjectJob *job.Job, localJob bool) ([]*job.Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) GetDownstream(ctx context.Context, subjectJob *job.Job, localJob bool) ([]*job.Downstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *JobService) raiseCreateEvent(job *job.Job) { _ = "STUB: not implemented"; return }

func (j *JobService) raiseUpdateEvent(job *job.Job) { _ = "STUB: not implemented"; return }

func (j *JobService) raiseStateChangeEvent(tnnt tenant.Tenant, jobName job.Name, state job.State) {
	_ = "STUB: not implemented"
	return
}

func (j *JobService) raiseDeleteEvent(tnnt tenant.Tenant, jobName job.Name) {
	_ = "STUB: not implemented"
	return
}

func (*JobService) groupDownstreamPerProject(downstreams []*job.Downstream) map[tenant.ProjectName][]*job.Downstream {
	_ = "STUB: not implemented"
	return nil
}

func raiseJobEventMetric(jobTenant tenant.Tenant, state string, metricValue int) {
	_ = "STUB: not implemented"
	return
}
