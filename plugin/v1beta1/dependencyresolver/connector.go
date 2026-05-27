package dependencyresolver

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	oplugin "github.com/raystack/optimus/sdk/plugin"
)

var _ plugin.GRPCPlugin = &Connector{}

type Connector struct {
	plugin.NetRPCUnsupportedPlugin
	plugin.GRPCPlugin

	impl oplugin.DependencyResolverMod

	logger hclog.Logger
}

func (p *Connector) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Connector) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPlugin(impl oplugin.DependencyResolverMod, logger hclog.Logger) *Connector {
	_ = "STUB: not implemented"
	return nil
}

func NewPluginClient(logger hclog.Logger) *Connector { _ = "STUB: not implemented"; return nil }

func Serve(t oplugin.DependencyResolverMod, logger hclog.Logger) { _ = "STUB: not implemented"; return }
