package survey

import (
	"context"

	"github.com/AlecAivazis/survey/v2"

	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/sdk/plugin"
)

// JobSurvey defines survey for job specification in general
type JobSurvey struct{}

// NewJobSurvey initializes job survey
func NewJobSurvey() *JobSurvey { _ = "STUB: not implemented"; return nil }

// AskToSelectJobName asks to select job name
func (*JobSurvey) AskToSelectJobName(jobSpecReader local.SpecReader[*model.JobSpec], jobDirPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (j *JobSurvey) askCliModSurveyQuestion(ctx context.Context, cliMod plugin.CommandLineMod, question plugin.Question) (plugin.Answers, error) {
	_ = "STUB: not implemented" //nolint: gocritic
	return *new(plugin.Answers), nil
}

// check if sub questions are attached on this question

//nolint: gocritic

func (*JobSurvey) getSurveyPromptFromPluginQuestion(question plugin.Question) survey.Prompt {
	_ = "STUB: not implemented" //nolint: gocritic
	return *new(survey.Prompt)
}

func (j *JobSurvey) getValidatePluginQuestion(ctx context.Context, cliMod plugin.CommandLineMod, question plugin.Question) survey.Validator {
	_ = "STUB: not implemented" //nolint: gocritic
	return *new(survey.Validator)
}

//nolint: gocritic

func (*JobSurvey) convertUserInputPluginToString(val interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
