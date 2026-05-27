package survey

import (
	"github.com/AlecAivazis/survey/v2"

	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
)

// ResourceSpecCreateSurvey defines surveys for resource spec creation
type ResourceSpecCreateSurvey struct {
	resourceSpecReader local.SpecReader[*model.ResourceSpec]
}

// NewResourceSpecCreateSurvey initializes survey for resource spec create
func NewResourceSpecCreateSurvey(resourceSpecReader local.SpecReader[*model.ResourceSpec]) *ResourceSpecCreateSurvey {
	_ = "STUB: not implemented"
	return nil
}

// AskResourceSpecName asks the user to input the required resource spec name
func (r ResourceSpecCreateSurvey) AskResourceSpecName(rootDirPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ResourceSpecCreateSurvey) AskResourceSpecType() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r ResourceSpecCreateSurvey) isResourceSpecNameUnique(rootDirPath string) survey.Validator {
	_ = "STUB: not implemented"
	return *new(survey.Validator)
}
