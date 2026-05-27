package server

import (
	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/ext/scheduler/airflow"
	"github.com/raystack/optimus/ext/scheduler/airflow/dag"
)

func NewScheduler(l log.Logger, conf *config.ServerConfig, pluginRepo dag.PluginRepo, projecGetter airflow.ProjectGetter,
	secretGetter airflow.SecretGetter,
) (*airflow.Scheduler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
