package service

import (
	"context"
	"time"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/core/tenant"
)

const (
	// taskConfigPrefix will be used to prefix all the config variables of
	// transformation instance, i.e. task
	taskConfigPrefix = "TASK__"

	// projectConfigPrefix will be used to prefix all the config variables of
	// a project, i.e. registered entities
	projectConfigPrefix = "GLOBAL__"

	contextProject       = "proj"
	contextSecret        = "secret"
	contextSystemDefined = "inst"
	contextTask          = "task"

	SecretsStringToMatch = ".secret."

	TimeISOFormat = time.RFC3339

	// Configuration for system defined variables
	configDstart        = "DSTART"
	configDend          = "DEND"
	configExecutionTime = "EXECUTION_TIME"
	configDestination   = "JOB_DESTINATION"
)

type TenantService interface {
	GetDetails(ctx context.Context, tnnt tenant.Tenant) (*tenant.WithDetails, error)
	GetSecrets(ctx context.Context, tnnt tenant.Tenant) ([]*tenant.PlainTextSecret, error)
}

type TemplateCompiler interface {
	Compile(templateMap map[string]string, context map[string]any) (map[string]string, error)
}

type AssetCompiler interface {
	CompileJobRunAssets(ctx context.Context, job *scheduler.Job, systemEnvVars map[string]string, scheduledAt time.Time, contextForTask map[string]interface{}) (map[string]string, error)
}

type InputCompiler struct {
	tenantService TenantService
	compiler      TemplateCompiler
	assetCompiler AssetCompiler

	logger log.Logger
}

func (i InputCompiler) Compile(ctx context.Context, job *scheduler.Job, config scheduler.RunConfig, executedAt time.Time) (*scheduler.ExecutorInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare template context and compile task config

// Compile asset files

// If request for hook, add task configs to templateContext

func (i InputCompiler) compileConfigs(configs map[string]string, templateCtx map[string]any) (map[string]string, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getSystemDefinedConfigs(job *scheduler.Job, runConfig scheduler.RunConfig, executedAt time.Time) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitConfigWithSecrets(conf map[string]string) (map[string]string, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJobInputCompiler(tenantService TenantService, compiler TemplateCompiler, assetCompiler AssetCompiler, logger log.Logger) *InputCompiler {
	_ = "STUB: not implemented"
	return nil
}
