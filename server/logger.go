package server

import (
	"io"

	"github.com/raystack/salt/log"
)

type defaultLogger struct {
	logger *log.Logrus
}

func (d defaultLogger) Debug(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Info(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Warn(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Error(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Fatal(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Level() string { _ = "STUB: not implemented"; return "" }

func (d defaultLogger) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func NewLogger(level string) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }
