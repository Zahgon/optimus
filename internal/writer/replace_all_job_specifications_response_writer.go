package writer

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type ReplaceAllJobSpecificationsResponseWriter interface {
	LogWriter
}

type replaceAllJobSpecificationsResponseWriter struct {
	stream pb.JobSpecificationService_ReplaceAllJobSpecificationsServer
}

func NewReplaceAllJobSpecificationsResponseWriter(stream pb.JobSpecificationService_ReplaceAllJobSpecificationsServer) ReplaceAllJobSpecificationsResponseWriter {
	_ = "STUB: not implemented"
	return *new(ReplaceAllJobSpecificationsResponseWriter)
}

func (s *replaceAllJobSpecificationsResponseWriter) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}
