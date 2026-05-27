package survey

import (
	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/config"
)

// InititalizeSurvey defines surveys related to init client config
type InititalizeSurvey struct {
	logger log.Logger
}

// NewInitializeSurvey initializes init survey
func NewInitializeSurvey(logger log.Logger) *InititalizeSurvey {
	_ = "STUB: not implemented"
	return nil
}

// AskToConfirm askes the user to confirm on a message
func (*InititalizeSurvey) AskToConfirm(message, help string, defaultValue bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AskInitClientConfig askes the user to init client config
func (i *InititalizeSurvey) AskInitClientConfig(dirPath string) (*config.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *InititalizeSurvey) askHost() (host string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *InititalizeSurvey) askInitProject() (project config.Project, err error) {
	_ = "STUB: not implemented"
	return *new(config.Project), nil
}

func (i *InititalizeSurvey) askInitNamespaces(dirPath string) ([]*config.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*InititalizeSurvey) askInitNamespaceDatastoreType() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *InititalizeSurvey) askInitNamespaceName(dirPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
