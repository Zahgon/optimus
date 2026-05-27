package bigquery

import (
	"github.com/raystack/optimus/core/resource"
)

const (
	EntityTable = "resource_table"
)

type Table struct {
	Name resource.Name

	Description string     `mapstructure:"description,omitempty"`
	Schema      Schema     `mapstructure:"schema,omitempty"`
	Cluster     *Cluster   `mapstructure:"cluster,omitempty"`
	Partition   *Partition `mapstructure:"partition,omitempty"`

	ExtraConfig map[string]interface{} `mapstructure:",remain"`
}

func (t *Table) FullName() string { _ = "STUB: not implemented"; return "" }

func (t *Table) Validate() error { _ = "STUB: not implemented"; return nil }

type Cluster struct {
	Using []string `mapstructure:"using,omitempty"`
}

func (c Cluster) Validate() error { _ = "STUB: not implemented"; return nil }

type Partition struct {
	Field string `mapstructure:"field,omitempty"`

	Type       string `mapstructure:"type,omitempty"`
	Expiration int64  `mapstructure:"expiration,omitempty"`

	Range *Range `mapstructure:"range,omitempty"`
}

func (p Partition) Validate() error { _ = "STUB: not implemented"; return nil }

type Range struct {
	Start    int64 `mapstructure:"start,omitempty"`
	End      int64 `mapstructure:"end,omitempty"`
	Interval int64 `mapstructure:"interval,omitempty"`
}
