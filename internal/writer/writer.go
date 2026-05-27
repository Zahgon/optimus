package writer

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type LogLevel int

const (
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

type LogWriter interface {
	Write(LogLevel, string) error
}

func newLogStatusProto(lvl LogLevel, msg string) *pb.Log { _ = "STUB: not implemented"; return nil }
