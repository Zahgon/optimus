package model

import (
	"time"

	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type JobSpec struct {
	Version      int                 `yaml:"version,omitempty"`
	Name         string              `yaml:"name"`
	Owner        string              `yaml:"owner"`
	Description  string              `yaml:"description,omitempty"`
	Schedule     JobSpecSchedule     `yaml:"schedule"`
	Behavior     JobSpecBehavior     `yaml:"behavior"`
	Task         JobSpecTask         `yaml:"task"`
	Asset        map[string]string   `yaml:"-"`
	Labels       map[string]string   `yaml:"labels,omitempty"`
	Hooks        []JobSpecHook       `yaml:"hooks"`
	Dependencies []JobSpecDependency `yaml:"dependencies"`
	Metadata     *JobSpecMetadata    `yaml:"metadata,omitempty"`
	Path         string              `yaml:"-"`
}

type JobSpecSchedule struct {
	StartDate string `yaml:"start_date"`
	EndDate   string `yaml:"end_date,omitempty"`
	Interval  string `yaml:"interval"`
}

type JobSpecBehavior struct {
	DependsOnPast bool                      `yaml:"depends_on_past"`
	Retry         *JobSpecBehaviorRetry     `yaml:"retry,omitempty"`
	Notify        []JobSpecBehaviorNotifier `yaml:"notify,omitempty"`
}

type JobSpecBehaviorRetry struct {
	Count              int           `yaml:"count,omitempty"`
	Delay              time.Duration `yaml:"delay,omitempty"`
	ExponentialBackoff bool          `yaml:"exponential_backoff,omitempty"`
}

type JobSpecBehaviorNotifier struct {
	On       string            `yaml:"on"`
	Config   map[string]string `yaml:"config"`
	Channels []string          `yaml:"channels"`
}

type JobSpecTask struct {
	Name   string            `yaml:"name"`
	Config map[string]string `yaml:"config,omitempty"`
	Window JobSpecTaskWindow `yaml:"window"`
}

type JobSpecTaskWindow struct {
	Size       string `yaml:"size"`
	Offset     string `yaml:"offset"`
	TruncateTo string `yaml:"truncate_to"`
}

type JobSpecHook struct {
	Name   string            `yaml:"name"`
	Config map[string]string `yaml:"config,omitempty"`
}

type JobSpecDependency struct {
	JobName string                 `yaml:"job,omitempty"`
	Type    string                 `yaml:"type,omitempty"`
	HTTP    *JobSpecDependencyHTTP `yaml:"http,omitempty"`
}

type JobSpecDependencyHTTP struct {
	Name          string            `yaml:"name"`
	RequestParams map[string]string `yaml:"params,omitempty"`
	URL           string            `yaml:"url"`
	Headers       map[string]string `yaml:"headers,omitempty"`
}

type JobSpecMetadata struct {
	Resource *JobSpecMetadataResource `yaml:"resource,omitempty"`
	Airflow  *JobSpecMetadataAirflow  `yaml:"airflow,omitempty"`
}

type JobSpecMetadataResource struct {
	Request *JobSpecMetadataResourceConfig `yaml:"request,omitempty"`
	Limit   *JobSpecMetadataResourceConfig `yaml:"limit,omitempty"`
}

type JobSpecMetadataResourceConfig struct {
	Memory string `yaml:"memory,omitempty"`
	CPU    string `yaml:"cpu,omitempty"`
}

type JobSpecMetadataAirflow struct {
	Pool  string `yaml:"pool" json:"pool"`
	Queue string `yaml:"queue" json:"queue"`
}

func (j *JobSpec) ToProto() *pb.JobSpecification { _ = "STUB: not implemented"; return nil }

func (j *JobSpec) getProtoJobMetadata() *pb.JobMetadata { _ = "STUB: not implemented"; return nil }

func (*JobSpec) getProtoJobSpecMetadataResourceConfig(jobSpecMetadataResourceConfig *JobSpecMetadataResourceConfig) *pb.JobSpecMetadataResourceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSpec) getProtoJobSpecBehavior() *pb.JobSpecification_Behavior {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSpec) getProtoJobSpecHooks() []*pb.JobSpecHook { _ = "STUB: not implemented"; return nil }

func (j *JobSpec) getProtoJobDependencies() []*pb.JobDependency {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobSpec) getProtoJobConfigItems() []*pb.JobConfigItem {
	_ = "STUB: not implemented"
	return nil
}

// TODO: on server, convert name to upper case

// TODO: there are some refactors required, however it will be addressed once we relook at the job spec inheritance
func (j *JobSpec) MergeFrom(anotherJobSpec *JobSpec) { _ = "STUB: not implemented"; return }

// already exists just inherit

// configs

// channels

// check if hook already present in child

// try to copy configs

// copy non existing hooks

// TODO: refactor this function since this is used only within a single method in job spec.
// the intent is also not clear, especially when reading the calling method above.
func getValue[V int | string | bool | time.Duration](reference, other V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

func ToJobSpec(protoSpec *pb.JobSpecification) *JobSpec { _ = "STUB: not implemented"; return nil }

func toJobSpecMetadata(protoMetadata *pb.JobMetadata) *JobSpecMetadata {
	_ = "STUB: not implemented"
	return nil
}

func toJobSpecDependencies(protoDependencies []*pb.JobDependency) []JobSpecDependency {
	_ = "STUB: not implemented"
	return nil
}

func toJobSpecHooks(protoHooks []*pb.JobSpecHook) []JobSpecHook {
	_ = "STUB: not implemented"
	return nil
}

func toJobSpecBehavior(protoBehavior *pb.JobSpecification_Behavior, dependsOnPast bool) JobSpecBehavior {
	_ = "STUB: not implemented"
	return *new(JobSpecBehavior)
}

func configProtoToMap(configProtoItems []*pb.JobConfigItem) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
