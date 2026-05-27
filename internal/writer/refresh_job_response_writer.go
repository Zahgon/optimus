package writer

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type RefreshJobResponseWriter interface {
	LogWriter
}

type refreshJobResponseWriter struct {
	stream pb.JobSpecificationService_RefreshJobsServer
}

func NewRefreshJobResponseWriter(stream pb.JobSpecificationService_RefreshJobsServer) RefreshJobResponseWriter {
	_ = "STUB: not implemented"
	return *new(RefreshJobResponseWriter)
}

func (s *refreshJobResponseWriter) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}
