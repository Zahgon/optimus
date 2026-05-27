package job

import (
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/models"
)

const (
	DateLayout       = "2006-01-02"
	maxJobNameLength = 125
)

type Spec struct {
	version  int
	name     Name
	owner    string
	schedule *Schedule
	window   models.Window
	task     Task

	description  string
	labels       map[string]string
	metadata     *Metadata
	hooks        []*Hook
	asset        Asset
	alertSpecs   []*AlertSpec
	upstreamSpec *UpstreamSpec
}

func (s *Spec) Version() int { _ = "STUB: not implemented"; return 0 }

func (s *Spec) Name() Name { _ = "STUB: not implemented"; return *new(Name) }

func (s *Spec) Owner() string { _ = "STUB: not implemented"; return "" }

func (s *Spec) Schedule() *Schedule { _ = "STUB: not implemented"; return nil }

func (s *Spec) Window() models.Window { _ = "STUB: not implemented"; return *new(models.Window) }

func (s *Spec) Task() Task { _ = "STUB: not implemented"; return *new(Task) }

func (s *Spec) Description() string { _ = "STUB: not implemented"; return "" }

func (s *Spec) Labels() map[string]string { _ = "STUB: not implemented"; return nil }

func (s *Spec) Hooks() []*Hook { _ = "STUB: not implemented"; return nil }

func (s *Spec) AlertSpecs() []*AlertSpec { _ = "STUB: not implemented"; return nil }

func (s *Spec) UpstreamSpec() *UpstreamSpec { _ = "STUB: not implemented"; return nil }

func (s *Spec) Asset() Asset { _ = "STUB: not implemented"; return *new(Asset) }

func (s *Spec) Metadata() *Metadata { _ = "STUB: not implemented"; return nil }

type SpecBuilder struct {
	spec *Spec
}

func NewSpecBuilder(version int, name Name, owner string, schedule *Schedule, window models.Window, task Task) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecBuilder) Build() (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SpecBuilder) WithHooks(hooks []*Hook) *SpecBuilder { _ = "STUB: not implemented"; return nil }

func (s *SpecBuilder) WithAlerts(alerts []*AlertSpec) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecBuilder) WithSpecUpstream(specUpstream *UpstreamSpec) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecBuilder) WithAsset(asset Asset) *SpecBuilder { _ = "STUB: not implemented"; return nil }

func (s *SpecBuilder) WithMetadata(metadata *Metadata) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecBuilder) WithLabels(labels map[string]string) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecBuilder) WithDescription(description string) *SpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

type Specs []*Spec

func (s Specs) ToNameAndSpecMap() map[Name]*Spec { _ = "STUB: not implemented"; return nil }

func (s Specs) ToFullNameAndSpecMap(projectName tenant.ProjectName) map[FullName]*Spec {
	_ = "STUB: not implemented"
	return nil
}

func (s Specs) Validate() error { _ = "STUB: not implemented"; return nil }

func (s Specs) GetValid() []*Spec { _ = "STUB: not implemented"; return nil }

func (s Specs) getJobNameCount() map[Name]int { _ = "STUB: not implemented"; return nil }

type Name string

func NameFrom(name string) (Name, error) { _ = "STUB: not implemented"; return *new(Name), nil }

func (n Name) String() string { _ = "STUB: not implemented"; return "" }

type State string

const (
	ENABLED  State = "enabled"
	DISABLED State = "disabled"
)

func StateFrom(name string) (State, error) { _ = "STUB: not implemented"; return *new(State), nil }

func (n State) String() string { _ = "STUB: not implemented"; return "" }

type ScheduleDate string

func ScheduleDateFrom(date string) (ScheduleDate, error) {
	_ = "STUB: not implemented"
	return *new(ScheduleDate), nil
}

func (s ScheduleDate) String() string { _ = "STUB: not implemented"; return "" }

type Retry struct {
	count              int
	delay              int32
	exponentialBackoff bool
}

func (r Retry) Count() int { _ = "STUB: not implemented"; return 0 }

func (r Retry) Delay() int32 { _ = "STUB: not implemented"; return 0 }

func (r Retry) ExponentialBackoff() bool { _ = "STUB: not implemented"; return false }

func NewRetry(count int, delay int32, exponentialBackoff bool) *Retry {
	_ = "STUB: not implemented"
	return nil
}

type Schedule struct {
	startDate     ScheduleDate
	endDate       ScheduleDate
	interval      string
	dependsOnPast bool
	retry         *Retry
}

func (s Schedule) StartDate() ScheduleDate { _ = "STUB: not implemented"; return *new(ScheduleDate) }

func (s Schedule) EndDate() ScheduleDate { _ = "STUB: not implemented"; return *new(ScheduleDate) }

func (s Schedule) Interval() string { _ = "STUB: not implemented"; return "" }

func (s Schedule) DependsOnPast() bool { _ = "STUB: not implemented"; return false }

func (s Schedule) Retry() *Retry { _ = "STUB: not implemented"; return nil }

type ScheduleBuilder struct {
	schedule *Schedule
}

// TODO: move interval to optional
func NewScheduleBuilder(startDate ScheduleDate) *ScheduleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScheduleBuilder) Build() (*Schedule, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *ScheduleBuilder) WithInterval(interval string) *ScheduleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScheduleBuilder) WithEndDate(endDate ScheduleDate) *ScheduleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScheduleBuilder) WithDependsOnPast(dependsOnPast bool) *ScheduleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScheduleBuilder) WithRetry(retry *Retry) *ScheduleBuilder {
	_ = "STUB: not implemented"
	return nil
}

type Config map[string]string

func ConfigFrom(configs map[string]string) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

func (c Config) Map() map[string]string { _ = "STUB: not implemented"; return nil }

type TaskName string

func TaskNameFrom(name string) (TaskName, error) {
	_ = "STUB: not implemented"
	return *new(TaskName), nil
}

func (t TaskName) String() string { _ = "STUB: not implemented"; return "" }

type Task struct {
	name   TaskName
	config Config
}

func NewTask(name TaskName, config Config) Task { _ = "STUB: not implemented"; return *new(Task) }

func (t Task) Name() TaskName { _ = "STUB: not implemented"; return *new(TaskName) }

func (t Task) Config() Config { _ = "STUB: not implemented"; return *new(Config) }

type MetadataResourceConfig struct {
	cpu    string
	memory string
}

func (m MetadataResourceConfig) CPU() string { _ = "STUB: not implemented"; return "" }

func (m MetadataResourceConfig) Memory() string { _ = "STUB: not implemented"; return "" }

func NewMetadataResourceConfig(cpu, memory string) *MetadataResourceConfig {
	_ = "STUB: not implemented"
	return nil
}

type MetadataResource struct {
	request *MetadataResourceConfig
	limit   *MetadataResourceConfig
}

func (m MetadataResource) Request() *MetadataResourceConfig { _ = "STUB: not implemented"; return nil }

func (m MetadataResource) Limit() *MetadataResourceConfig { _ = "STUB: not implemented"; return nil }

func NewResourceMetadata(request, limit *MetadataResourceConfig) *MetadataResource {
	_ = "STUB: not implemented"
	return nil
}

type Metadata struct {
	resource  *MetadataResource
	scheduler map[string]string
}

func (m Metadata) Resource() *MetadataResource { _ = "STUB: not implemented"; return nil }

func (m Metadata) Scheduler() map[string]string { _ = "STUB: not implemented"; return nil }

func (m Metadata) validate() error { _ = "STUB: not implemented"; return nil }

type MetadataBuilder struct {
	metadata *Metadata
}

func NewMetadataBuilder() *MetadataBuilder { _ = "STUB: not implemented"; return nil }

func (m *MetadataBuilder) Build() (*Metadata, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MetadataBuilder) WithResource(resource *MetadataResource) *MetadataBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MetadataBuilder) WithScheduler(scheduler map[string]string) *MetadataBuilder {
	_ = "STUB: not implemented"
	return nil
}

type Hook struct {
	name   string
	config Config
}

func NewHook(name string, config Config) (*Hook, error) { _ = "STUB: not implemented"; return nil, nil }

func (h Hook) Name() string { _ = "STUB: not implemented"; return "" }

func (h Hook) Config() Config { _ = "STUB: not implemented"; return *new(Config) }

type Asset map[string]string

func AssetFrom(fileNameToContent map[string]string) (Asset, error) {
	_ = "STUB: not implemented"
	return *new(Asset), nil
}

func (a Asset) Map() map[string]string { _ = "STUB: not implemented"; return nil }

func (a Asset) validate() error { _ = "STUB: not implemented"; return nil }

type AlertSpec struct {
	on string

	channels []string
	config   Config
}

func NewAlertSpec(on string, channels []string, config Config) (*AlertSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a AlertSpec) On() string { _ = "STUB: not implemented"; return "" }

func (a AlertSpec) Channels() []string { _ = "STUB: not implemented"; return nil }

func (a AlertSpec) Config() Config {
	_ = "STUB: not implemented"

	// TODO: reconsider whether we still need it or not
	return *new(Config)
}

type SpecHTTPUpstream struct {
	name    string
	url     string
	headers map[string]string
	params  map[string]string
}

func (s SpecHTTPUpstream) Name() string { _ = "STUB: not implemented"; return "" }

func (s SpecHTTPUpstream) URL() string { _ = "STUB: not implemented"; return "" }

func (s SpecHTTPUpstream) Headers() map[string]string { _ = "STUB: not implemented"; return nil }

func (s SpecHTTPUpstream) Params() map[string]string { _ = "STUB: not implemented"; return nil }

func (s SpecHTTPUpstream) validate() error { _ = "STUB: not implemented"; return nil }

type SpecHTTPUpstreamBuilder struct {
	upstream *SpecHTTPUpstream
}

func NewSpecHTTPUpstreamBuilder(name, url string) *SpecHTTPUpstreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecHTTPUpstreamBuilder) Build() (*SpecHTTPUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SpecHTTPUpstreamBuilder) WithHeaders(headers map[string]string) *SpecHTTPUpstreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecHTTPUpstreamBuilder) WithParams(params map[string]string) *SpecHTTPUpstreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

type SpecUpstreamName string

func (s SpecUpstreamName) String() string { _ = "STUB: not implemented"; return "" }

func SpecUpstreamNameFrom(specUpstreamName string) SpecUpstreamName {
	_ = "STUB: not implemented"
	return *new(SpecUpstreamName)
}

func (s SpecUpstreamName) IsWithProjectName() bool { _ = "STUB: not implemented"; return false }

func (s SpecUpstreamName) GetProjectName() (tenant.ProjectName, error) {
	_ = "STUB: not implemented"
	return *new(tenant.ProjectName), nil
}

func (s SpecUpstreamName) GetJobName() (Name, error) {
	_ = "STUB: not implemented"
	return *new(Name), nil
}

type UpstreamSpec struct {
	upstreamNames []SpecUpstreamName
	httpUpstreams []*SpecHTTPUpstream
}

func (s UpstreamSpec) UpstreamNames() []SpecUpstreamName { _ = "STUB: not implemented"; return nil }

func (s UpstreamSpec) HTTPUpstreams() []*SpecHTTPUpstream { _ = "STUB: not implemented"; return nil }

func (s UpstreamSpec) validate() error { _ = "STUB: not implemented"; return nil }

type SpecUpstreamBuilder struct {
	upstream *UpstreamSpec
}

func NewSpecUpstreamBuilder() *SpecUpstreamBuilder { _ = "STUB: not implemented"; return nil }

func (s *SpecUpstreamBuilder) Build() (*UpstreamSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SpecUpstreamBuilder) WithUpstreamNames(names []SpecUpstreamName) *SpecUpstreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpecUpstreamBuilder) WithSpecHTTPUpstream(httpUpstreams []*SpecHTTPUpstream) *SpecUpstreamBuilder {
	_ = "STUB: not implemented"
	return nil
}

func NewLabels(labels map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check whether this is supposed to be here or in utils
func validateMap(input map[string]string) error { _ = "STUB: not implemented"; return nil }
