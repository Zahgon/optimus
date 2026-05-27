package dag

import (
	_ "embed"
	"text/template"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/sdk/plugin"
)

//go:embed dag.py.tmpl
var dagTemplate []byte

type PluginRepo interface {
	GetByName(name string) (*plugin.Plugin, error)
}

type Compiler struct {
	hostname string

	template   *template.Template
	pluginRepo PluginRepo
}

func (c *Compiler) Compile(jobDetails *scheduler.JobWithDetails) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDagCompiler(hostname string, repo PluginRepo) (*Compiler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
