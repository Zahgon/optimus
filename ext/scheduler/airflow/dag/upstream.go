package dag

import (
	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

type Upstreams struct {
	HTTP      []*scheduler.HTTPUpstreams
	Upstreams []Upstream
}

func (u Upstreams) Empty() bool { _ = "STUB: not implemented"; return false }

type Upstream struct {
	JobName  string
	Tenant   tenant.Tenant
	Host     string
	TaskName string
}

func SetupUpstreams(upstreams scheduler.Upstreams, host string) Upstreams {
	_ = "STUB: not implemented"
	return *new(Upstreams)
}
