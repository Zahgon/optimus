package bigquery

import (
	"github.com/raystack/optimus/core/resource"
)

const (
	EntityExternalTable = "resource_external_table"
)

type ExternalTable struct {
	Name resource.Name

	Description string          `mapstructure:"description,omitempty"`
	Schema      Schema          `mapstructure:"schema,omitempty"`
	Source      *ExternalSource `mapstructure:"source,omitempty"`

	ExtraConfig map[string]interface{} `mapstructure:",remain"`
}

func (e *ExternalTable) FullName() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalTable) Validate() error { _ = "STUB: not implemented"; return nil }

type ExternalSource struct {
	SourceType string   `mapstructure:"type,omitempty"`
	SourceURIs []string `mapstructure:"uris,omitempty"`

	// Additional configs for CSV, GoogleSheets, Bigtable, and Parquet formats.
	Config map[string]interface{} `mapstructure:"config"`
}

func (e ExternalSource) Validate() error { _ = "STUB: not implemented"; return nil }
