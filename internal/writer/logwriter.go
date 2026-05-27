package writer

import (
	"github.com/raystack/salt/log"

	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type saltLogger struct {
	l log.Logger
}

func NewLogWriter(l log.Logger) LogWriter { _ = "STUB: not implemented"; return *new(LogWriter) }

func (l *saltLogger) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}

type BufferedLogger struct {
	Messages []*pb.Log
}

// nolint: unparam
func (b *BufferedLogger) Write(level LogLevel, message string) error {
	_ = "STUB: not implemented"
	return nil
}
