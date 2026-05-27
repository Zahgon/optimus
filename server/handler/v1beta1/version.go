package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type VersionHandler struct {
	version string
	l       log.Logger

	pb.UnimplementedRuntimeServiceServer
}

func (vh VersionHandler) Version(_ context.Context, version *pb.VersionRequest) (*pb.VersionResponse, error) {
	_ = "STUB: not implemented" // nolint: unparam
	return nil, nil
}

func NewVersionHandler(l log.Logger, version string) *VersionHandler {
	_ = "STUB: not implemented"
	return nil
}
