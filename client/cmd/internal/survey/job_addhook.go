package survey

import (
	"context"

	"github.com/raystack/optimus/client/local/model"
	"github.com/raystack/optimus/internal/models"
	"github.com/raystack/optimus/sdk/plugin"
)

// JobAddHookSurvey defines survey for job add hook
type JobAddHookSurvey struct {
	jobSurvey *JobSurvey
}

// NewJobAddHookSurvey initializes job add hook survey
func NewJobAddHookSurvey() *JobAddHookSurvey { _ = "STUB: not implemented"; return nil }

// AskToAddHook asks questions to add hook to a job
func (j *JobAddHookSurvey) AskToAddHook(pluginRepo *models.PluginRepository, jobSpec *model.JobSpec) (*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: remove the golint exception below
//nolint:gocritic

func (*JobAddHookSurvey) getHookConfig(cliMod plugin.CommandLineMod, answers plugin.Answers) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*JobAddHookSurvey) getAvailableHookNames(pluginRepo *models.PluginRepository) []string {
	_ = "STUB: not implemented"
	return nil
}

func (*JobAddHookSurvey) askToSelectHook(options []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*JobAddHookSurvey) isSelectedHookAlreadyInJob(jobSpec *model.JobSpec, selectedHookName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (j *JobAddHookSurvey) askHookQuestions(ctx context.Context, cliMod plugin.CommandLineMod, jobName string) (plugin.Answers, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Answers), nil
}

//nolint: gocritic
