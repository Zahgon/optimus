package setup

import (
	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/internal/models"
)

type DummyJobBuilder struct {
	version     int
	owner       string
	description string

	retry     *job.Retry
	startDate job.ScheduleDate

	window models.Window

	taskConfig job.Config
	taskName   job.TaskName

	labels map[string]string

	hookConfig job.Config
	hookName   string

	alertConfig       job.Config
	alertName         string
	alertChannelNames []string

	asset map[string]string

	resourceRequestConfig *job.MetadataResourceConfig
	resourceLimitConfig   *job.MetadataResourceConfig
	scheduler             map[string]string

	name job.Name

	destinationURN job.ResourceURN
	sourceURNs     []job.ResourceURN

	specUpstreamNames []job.SpecUpstreamName
	specHTTPUpstreams []*job.SpecHTTPUpstream
}

func NewDummyJobBuilder() *DummyJobBuilder { _ = "STUB: not implemented"; return nil }

func (d *DummyJobBuilder) OverrideVersion(version int) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideOwner(owner string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideDescription(description string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideRetry(retry *job.Retry) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideStartDate(startDate job.ScheduleDate) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideWindow(window models.Window) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideTaskConfig(taskConfig job.Config) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideTaskName(taskName job.TaskName) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideLabels(labels map[string]string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideHookConfig(hookConfig job.Config) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideHookName(hookName string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideAlertConfig(alertConfig job.Config) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideAlertName(alertName string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideAlertChannelNames(alertChannelNames []string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideAsset(asset map[string]string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideResourceRequestConfig(config *job.MetadataResourceConfig) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideResourceLimitConfig(config *job.MetadataResourceConfig) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideScheduler(scheduler map[string]string) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideName(name job.Name) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideDestinationURN(destinationURN job.ResourceURN) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideSourceURNs(sourceURNs []job.ResourceURN) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideSpecUpstreamNames(specUpstreamNames []job.SpecUpstreamName) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) OverrideSpecHTTPUpstreams(specHTTPUpstreams []*job.SpecHTTPUpstream) *DummyJobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DummyJobBuilder) Build(tnnt tenant.Tenant) *job.Job { _ = "STUB: not implemented"; return nil }
