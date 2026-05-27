package survey

import (
	"github.com/AlecAivazis/survey/v2"

	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/internal/models"
	"github.com/raystack/optimus/internal/utils"
	"github.com/raystack/optimus/sdk/plugin"
)

const (
	ISODateLayout         = "2006-01-02"
	jobSpecDefaultVersion = 2
)

var (
	validateDate         = utils.ValidatorFactory.NewFromRegex(`\d{4}-\d{2}-\d{2}`, "date must be in YYYY-MM-DD format")
	validateNoSlash      = utils.ValidatorFactory.NewFromRegex(`^[^/]+$`, "`/` is disallowed")
	validateResourceName = utils.ValidatorFactory.NewFromRegex(`^[a-zA-Z0-9][a-zA-Z0-9_\-\.]+$`,
		`invalid name (can only contain characters A-Z (in either case), 0-9, "-", "_" or "." and must start with an alphanumeric character)`)
	validateJobName = survey.ComposeValidators(validateNoSlash, validateResourceName, survey.MinLength(3),
		survey.MaxLength(220))
)

// JobCreateSurvey defines survey for job creation operation
type JobCreateSurvey struct {
	jobSurvey *JobSurvey
}

// NewJobCreateSurvey initializes job create survey
func NewJobCreateSurvey() *JobCreateSurvey { _ = "STUB: not implemented"; return nil }

// AskToCreateJob asks questions to create job
func (j *JobCreateSurvey) AskToCreateJob(pluginRepo *models.PluginRepository, jobSpecReader local.SpecReader[*model.JobSpec], jobDir, defaultJobName string) (model.JobSpec, error) {
	_ = "STUB: not implemented"
	return *new(model.JobSpec), nil
}

func (*JobCreateSurvey) getJobAsset(cliMod plugin.CommandLineMod, answers plugin.Answers) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*JobCreateSurvey) getTaskConfig(cliMod plugin.CommandLineMod, answers plugin.Answers) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*JobCreateSurvey) getAvailableTaskNames(pluginRepo *models.PluginRepository) []string {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobCreateSurvey) getCreateQuestions(jobSpecReader local.SpecReader[*model.JobSpec], jobDir, defaultJobName string, availableTaskNames []string) []*survey.Question {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobCreateSurvey) askCreateQuestions(questions []*survey.Question) (model.JobSpec, error) {
	_ = "STUB: not implemented"
	return *new(model.JobSpec), nil
}

func (*JobCreateSurvey) getPluginCliMod(pluginRepo *models.PluginRepository, taskName string) (plugin.CommandLineMod, error) {
	_ = "STUB: not implemented"
	return *new(plugin.CommandLineMod), nil
}

func (j *JobCreateSurvey) askPluginQuestions(cliMod plugin.CommandLineMod, jobName string) (plugin.Answers, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Answers), nil
}

//nolint: gocritic
//nolint: gocritic

// getValidateJobUniqueness return a validator that checks if the job already exists with the same name
func (*JobCreateSurvey) getValidateJobUniqueness(jobSpecReader local.SpecReader[*model.JobSpec], jobDir string) survey.Validator {
	_ = "STUB: not implemented"
	return *new(survey.Validator)
}

func (*JobCreateSurvey) getWindowParameters(winName string) model.JobSpecTaskWindow {
	_ = "STUB: not implemented"
	return *new(model.JobSpecTaskWindow)
}

// default
