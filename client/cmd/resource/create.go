package resource

import (
	"github.com/raystack/salt/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/client/cmd/internal/survey"
	"github.com/raystack/optimus/config"
)

type createCommand struct {
	logger         log.Logger
	configFilePath string

	namespaceSurvey *survey.NamespaceSurvey
}

// NewCreateCommand initializes resource create command
func NewCreateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (c createCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: re-check if datastore needs to be in slice, currently assuming

// we are using the first datastore since we want to support only one datastore for a single namespace

// CreateDataStoreSpecFs creates specFS for data store
func CreateDataStoreSpecFs(namespace *config.Namespace) map[string]afero.Fs {
	_ = "STUB: not implemented"
	return nil
}
