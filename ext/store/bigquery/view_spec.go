package bigquery

import (
	"github.com/raystack/optimus/core/resource"
)

const (
	EntityView = "resource_view"
)

type View struct {
	Name resource.Name

	Description string `mapstructure:"description,omitempty"`
	ViewQuery   string `mapstructure:"view_query,omitempty"`

	ExtraConfig map[string]interface{} `mapstructure:",remain"`
}

func (v *View) Validate() error { _ = "STUB: not implemented"; return nil }
