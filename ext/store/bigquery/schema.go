package bigquery

const (
	ModeNullable = "nullable"
	ModeRequired = "required"
	ModeRepeated = "repeated"

	EntityResourceSchema = "bigquery_schema"
)

const (
	KindDataset       string = "dataset"
	KindTable         string = "table"
	KindView          string = "view"
	KindExternalTable string = "external_table"
)

type Schema []Field

func (s Schema) Validate() error { _ = "STUB: not implemented"; return nil }

type Field struct {
	Name        string `mapstructure:"name,omitempty"`
	Type        string `mapstructure:"type,omitempty"`
	Description string `mapstructure:"description,omitempty"`
	Mode        string `mapstructure:"mode,omitempty"`

	// optional sub-schema, when record type
	Schema Schema `mapstructure:"schema,omitempty"`
}

func (f Field) Validate() error {
	_ = "STUB: not implemented" // nolint:gocritic
	return nil
}
