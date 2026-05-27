package v1beta1

import (
	"github.com/raystack/optimus/core/job"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

func ToJobProto(jobEntity *job.Job) *pb.JobSpecification { _ = "STUB: not implemented"; return nil }

func fromJobProtos(protoJobSpecs []*pb.JobSpecification) ([]*job.Spec, []job.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func fromJobProto(js *pb.JobSpecification) (*job.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromResourceURNs(resourceURNs []job.ResourceURN) []string {
	_ = "STUB: not implemented"
	return nil
}

func fromRetryAndAlerts(jobRetry *job.Retry, alerts []*job.AlertSpec) *pb.JobSpecification_Behavior {
	_ = "STUB: not implemented"
	return nil
}

func toRetry(protoRetry *pb.JobSpecification_Behavior_Retry) *job.Retry {
	_ = "STUB: not implemented"
	return nil
}

func fromRetry(jobRetry *job.Retry) *pb.JobSpecification_Behavior_Retry {
	_ = "STUB: not implemented"
	return nil
}

func toHooks(hooksProto []*pb.JobSpecHook) ([]*job.Hook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromHooks(hooks []*job.Hook) []*pb.JobSpecHook { _ = "STUB: not implemented"; return nil }

func fromAsset(jobAsset job.Asset) map[string]string { _ = "STUB: not implemented"; return nil }

func toAlerts(notifiers []*pb.JobSpecification_Behavior_Notifiers) ([]*job.AlertSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromAlerts(jobAlerts []*job.AlertSpec) []*pb.JobSpecification_Behavior_Notifiers {
	_ = "STUB: not implemented"
	return nil
}

func toSpecUpstreams(upstreamProtos []*pb.JobDependency) (*job.UpstreamSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromSpecUpstreams(upstreams *job.UpstreamSpec) []*pb.JobDependency {
	_ = "STUB: not implemented"
	return nil
}

// TODO: upstream type?

func toMetadata(jobMetadata *pb.JobMetadata) (*job.Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromMetadata(metadata *job.Metadata) *pb.JobMetadata { _ = "STUB: not implemented"; return nil }

func toConfig(configs []*pb.JobConfigItem) (job.Config, error) {
	_ = "STUB: not implemented"
	return *new(job.Config), nil
}

func fromConfig(jobConfig job.Config) []*pb.JobConfigItem { _ = "STUB: not implemented"; return nil }

func toBasicInfoSectionProto(jobDetail *job.Job, logMessages []*pb.Log) *pb.JobInspectResponse_BasicInfoSection {
	_ = "STUB: not implemented"
	return nil
}

func toUpstreamProtos(upstreams []*job.Upstream, upstreamSpec *job.UpstreamSpec, upstreamLogs []*pb.Log) *pb.JobInspectResponse_UpstreamSection {
	_ = "STUB: not implemented"
	return nil
}

func toHTTPUpstreamProtos(httpUpstreamSpecs []*job.SpecHTTPUpstream) []*pb.HttpDependency {
	_ = "STUB: not implemented"
	return nil
}

func toDownstreamProtos(downstreamJobs []*job.Downstream) []*pb.JobInspectResponse_JobDependency {
	_ = "STUB: not implemented"
	return nil
}
