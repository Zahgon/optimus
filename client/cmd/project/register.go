package project

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

	dirPath        string
	withNamespaces bool
}

// NewRegisterCommand initializes command to create a project
func NewRegisterCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (r *registerCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterProject registers a project to the targeted server host
func RegisterProject(logger log.Logger, conn *grpc.ClientConn, project config.Project) error {
	_ = "STUB: not implemented"
	return nil
}
