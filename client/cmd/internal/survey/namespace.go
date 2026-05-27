package survey

import (
	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/config"
)

// NamespaceSurvey defines surveys related to namespace
type NamespaceSurvey struct {
	logger log.Logger
}

// NewNamespaceSurvey initializes namespace survey
func NewNamespaceSurvey(logger log.Logger) *NamespaceSurvey { _ = "STUB: not implemented"; return nil }

// AskToSelectNamespace askesk the user through CLI about namespace to be selected
func (n *NamespaceSurvey) AskToSelectNamespace(clientConfig *config.ClientConfig) (*config.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
