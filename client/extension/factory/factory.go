package factory

import (
	"github.com/raystack/optimus/client/extension/model"
)

// ParseRegistry is the registry for all parsers defined by each provider
var ParseRegistry []model.Parser

// ClientRegistry stores all clients defined by each provider
var ClientRegistry = &ClientFactory{}

// ClientFactory is a factory to store client
type ClientFactory struct {
	registry map[string]model.Client
}

// Add adds client based on provider
func (c *ClientFactory) Add(provider string, newClient model.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Get gets client for a specified provider
func (c *ClientFactory) Get(provider string) (model.Client, error) {
	_ = "STUB: not implemented"
	return *new(model.Client), nil
}
