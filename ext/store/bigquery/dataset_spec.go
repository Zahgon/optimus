package bigquery

import (
	"regexp"

	"github.com/raystack/optimus/core/resource"
)

const (
	EntityDataset = "dataset"

	DatesetNameSections = 2
	TableNameSections   = 3
)

var (
	validProjectName = regexp.MustCompile(`^[a-z][a-z0-9-]{4,28}[a-z0-9]$`)
	validDatasetName = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	validTableName   = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

type DatasetDetails struct {
	Name resource.Name

	Description string                 `mapstructure:"description,omitempty"`
	ExtraConfig map[string]interface{} `mapstructure:",remain"`
}

func (d DatasetDetails) FullName() string { _ = "STUB: not implemented"; return "" }

func (DatasetDetails) Validate() error { _ = "STUB: not implemented"; return nil }

func ConvertSpecTo[T DatasetDetails | Table | View | ExternalTable](res *resource.Resource) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Dataset struct {
	Project     string
	DatasetName string
}

func DataSetFrom(project, datasetName string) (Dataset, error) {
	_ = "STUB: not implemented"
	return *new(Dataset), nil
}

func (d Dataset) FullName() string { _ = "STUB: not implemented"; return "" }

func DataSetFor(res *resource.Resource) (Dataset, error) {
	_ = "STUB: not implemented"
	return *new(Dataset), nil
}

func ResourceNameFor(res *resource.Resource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ValidateName(res *resource.Resource) error { _ = "STUB: not implemented"; return nil }

func URNFor(res *resource.Resource) (string, error) { _ = "STUB: not implemented"; return "", nil }
