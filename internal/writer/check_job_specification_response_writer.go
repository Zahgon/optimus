package writer

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type checkJobSpecificationResponseWriter struct {
	stream pb.JobSpecificationService_CheckJobSpecificationsServer
}

func NewCheckJobSpecificationResponseWriter(stream pb.JobSpecificationService_CheckJobSpecificationsServer) LogWriter {
	_ = "STUB: not implemented"
	return *new(LogWriter)
}

func (s *checkJobSpecificationResponseWriter) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}
