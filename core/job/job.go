package job

import (
	"github.com/raystack/optimus/core/tenant"
)

const (
	EntityJob = "job"

	UpstreamTypeStatic   UpstreamType = "static"
	UpstreamTypeInferred UpstreamType = "inferred"

	UpstreamStateResolved   UpstreamState = "resolved"
	UpstreamStateUnresolved UpstreamState = "unresolved"

	MetricJobEvent                      = "job_events_total"
	MetricJobEventStateAdded            = "added"
	MetricJobEventStateUpdated          = "updated"
	MetricJobEventStateDeleted          = "deleted"
	MetricJobEventStateUpsertFailed     = "upsert_failed"
	MetricJobEventStateDeleteFailed     = "delete_failed"
	MetricJobEventStateValidationFailed = "validation_failed"
	MetricJobEventEnabled               = "enabled"
	MetricJobEventDisabled              = "disabled"

	MetricJobRefreshResourceDownstream = "refresh_resource_downstream_total"
)

type Job struct {
	tenant tenant.Tenant

	spec *Spec

	destination ResourceURN
	sources     []ResourceURN
}

func (j *Job) Tenant() tenant.Tenant { _ = "STUB: not implemented"; return *new(tenant.Tenant) }

func (j *Job) Spec() *Spec { _ = "STUB: not implemented"; return nil }

func (j *Job) GetName() string { _ = "STUB: not implemented"; return "" }

func (j *Job) FullName() string { _ = "STUB: not implemented"; return "" }

func (j *Job) GetJobWithUnresolvedUpstream() (*WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *Job) getInferredUpstreamsToResolve() []*Upstream { _ = "STUB: not implemented"; return nil }

func (j *Job) getStaticUpstreamsToResolve() ([]*Upstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResourceURN string

func (n ResourceURN) String() string { _ = "STUB: not implemented"; return "" }

func (j *Job) Destination() ResourceURN { _ = "STUB: not implemented"; return *new(ResourceURN) }

func (j *Job) Sources() []ResourceURN { _ = "STUB: not implemented"; return nil }

func (j *Job) StaticUpstreamNames() []SpecUpstreamName { _ = "STUB: not implemented"; return nil }

func (j *Job) ProjectName() tenant.ProjectName {
	_ = "STUB: not implemented"
	return *new(tenant.ProjectName)
}

func NewJob(tenant tenant.Tenant, spec *Spec, destination ResourceURN, sources []ResourceURN) *Job {
	_ = "STUB: not implemented"
	return nil
}

type Jobs []*Job

func (j Jobs) GetJobNames() []Name { _ = "STUB: not implemented"; return nil }

func (j Jobs) GetNameAndSpecMap() map[Name]*Spec { _ = "STUB: not implemented"; return nil }

func (j Jobs) GetNameAndJobMap() map[Name]*Job { _ = "STUB: not implemented"; return nil }

func (j Jobs) GetNamespaceNameAndJobsMap() map[tenant.NamespaceName][]*Job {
	_ = "STUB: not implemented"
	return nil
}

func (j Jobs) GetSpecs() []*Spec { _ = "STUB: not implemented"; return nil }

func (j Jobs) GetJobsWithUnresolvedUpstreams() ([]*WithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WithUpstream struct {
	job       *Job
	upstreams []*Upstream
}

func NewWithUpstream(job *Job, upstreams []*Upstream) *WithUpstream {
	_ = "STUB: not implemented"
	return nil
}

func (w WithUpstream) GetName() string {
	_ = "STUB: not implemented" // to support multiroot DataTree
	return ""
}

func (w WithUpstream) Job() *Job { _ = "STUB: not implemented"; return nil }

func (w WithUpstream) Upstreams() []*Upstream { _ = "STUB: not implemented"; return nil }

func (w WithUpstream) Name() Name { _ = "STUB: not implemented"; return *new(Name) }

func (w WithUpstream) GetUnresolvedUpstreams() []*Upstream { _ = "STUB: not implemented"; return nil }

func (w WithUpstream) GetResolvedUpstreams() []*Upstream { _ = "STUB: not implemented"; return nil }

type WithUpstreams []*WithUpstream

func (w WithUpstreams) GetSubjectJobNames() []Name { _ = "STUB: not implemented"; return nil }

func (w WithUpstreams) MergeWithResolvedUpstreams(resolvedUpstreamsBySubjectJobMap map[Name][]*Upstream) []*WithUpstream {
	_ = "STUB: not implemented"
	return nil
}

type Upstream struct {
	name     Name
	host     string
	resource ResourceURN
	taskName TaskName

	projectName   tenant.ProjectName
	namespaceName tenant.NamespaceName

	_type UpstreamType
	state UpstreamState

	external bool
}

func NewUpstreamResolved(name Name, host string, resource ResourceURN, jobTenant tenant.Tenant, upstreamType UpstreamType, taskName TaskName, external bool) *Upstream {
	_ = "STUB: not implemented"
	return nil
}

func NewUpstreamUnresolvedInferred(resource ResourceURN) *Upstream {
	_ = "STUB: not implemented"
	return nil
}

func NewUpstreamUnresolvedStatic(name Name, projectName tenant.ProjectName) *Upstream {
	_ = "STUB: not implemented"
	return nil
}

func (u *Upstream) Name() Name { _ = "STUB: not implemented"; return *new(Name) }

func (u *Upstream) Host() string { _ = "STUB: not implemented"; return "" }

func (u *Upstream) Resource() ResourceURN { _ = "STUB: not implemented"; return *new(ResourceURN) }

func (u *Upstream) Type() UpstreamType { _ = "STUB: not implemented"; return *new(UpstreamType) }

func (u *Upstream) State() UpstreamState { _ = "STUB: not implemented"; return *new(UpstreamState) }

func (u *Upstream) ProjectName() tenant.ProjectName {
	_ = "STUB: not implemented"
	return *new(tenant.ProjectName)
}

func (u *Upstream) NamespaceName() tenant.NamespaceName {
	_ = "STUB: not implemented"
	return *new(tenant.NamespaceName)
}

func (u *Upstream) External() bool { _ = "STUB: not implemented"; return false }

func (u *Upstream) TaskName() TaskName { _ = "STUB: not implemented"; return *new(TaskName) }

func (u *Upstream) FullName() string { _ = "STUB: not implemented"; return "" }

type UpstreamType string

func (d UpstreamType) String() string { _ = "STUB: not implemented"; return "" }

func UpstreamTypeFrom(str string) (UpstreamType, error) {
	_ = "STUB: not implemented"
	return *new(UpstreamType), nil
}

type UpstreamState string

func (d UpstreamState) String() string { _ = "STUB: not implemented"; return "" }

type Upstreams []*Upstream

func (u Upstreams) ToFullNameAndUpstreamMap() map[string]*Upstream {
	_ = "STUB: not implemented"
	return nil
}

func (u Upstreams) ToResourceDestinationAndUpstreamMap() map[string]*Upstream {
	_ = "STUB: not implemented"
	return nil
}

func (u Upstreams) Deduplicate() []*Upstream { _ = "STUB: not implemented"; return nil }

// keep static upstreams in the map if exists

func mapsToUpstreams(upstreamsMaps ...map[string]*Upstream) []*Upstream {
	_ = "STUB: not implemented"
	return nil
}

type FullName string

func FullNameFrom(projectName tenant.ProjectName, jobName Name) FullName {
	_ = "STUB: not implemented"
	return *new(FullName)
}

func (f FullName) String() string { _ = "STUB: not implemented"; return "" }

type FullNames []FullName

func (f FullNames) String() string { _ = "STUB: not implemented"; return "" }

type Downstream struct {
	name Name

	projectName   tenant.ProjectName
	namespaceName tenant.NamespaceName

	taskName TaskName
}

func NewDownstream(name Name, projectName tenant.ProjectName, namespaceName tenant.NamespaceName, taskName TaskName) *Downstream {
	_ = "STUB: not implemented"
	return nil
}

func (d Downstream) Name() Name { _ = "STUB: not implemented"; return *new(Name) }

func (d Downstream) ProjectName() tenant.ProjectName {
	_ = "STUB: not implemented"
	return *new(tenant.ProjectName)
}

func (d Downstream) NamespaceName() tenant.NamespaceName {
	_ = "STUB: not implemented"
	return *new(tenant.NamespaceName)
}

func (d Downstream) TaskName() TaskName { _ = "STUB: not implemented"; return *new(TaskName) }

func (d Downstream) FullName() FullName { _ = "STUB: not implemented"; return *new(FullName) }

type DownstreamList []*Downstream

func (d DownstreamList) GetDownstreamFullNames() FullNames {
	_ = "STUB: not implemented"
	return *new(FullNames)
}
