package writer

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type deployResourceSpecificationResponseWriter struct {
	stream pb.ResourceService_DeployResourceSpecificationServer
}

func NewDeployResourceSpecificationResponseWriter(stream pb.ResourceService_DeployResourceSpecificationServer) LogWriter {
	_ = "STUB: not implemented"
	return *new(LogWriter)
}

func (l *deployResourceSpecificationResponseWriter) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}
