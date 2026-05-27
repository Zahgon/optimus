package namespace

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/config"
)

const registerTimeout = time.Minute * 15

type registerCommand struct {
	logger log.Logger

	dirPath       string
	namespaceName string
}

// NewRegisterCommand initializes command for registering namespace
func NewRegisterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *registerCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterSelectedNamespaces registers all selected namespaces
func RegisterSelectedNamespaces(l log.Logger, conn *grpc.ClientConn, projectName string, selectedNamespaces ...*config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterNamespace registers one namespace to the targeted server
func RegisterNamespace(l log.Logger, conn *grpc.ClientConn, projectName string, namespace *config.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}
