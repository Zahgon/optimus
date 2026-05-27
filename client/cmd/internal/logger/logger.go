package logger

import (
	"io"

	"github.com/fatih/color"
	"github.com/raystack/salt/log"
)

type defaultLogger struct {
	writer   io.Writer
	exitFunc func(int)
}

func (d defaultLogger) Debug(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Info(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Warn(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Error(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (d defaultLogger) Fatal(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (defaultLogger) Level() string {
	_ = "STUB: not implemented"
	// this is to adhere to the logger interface
	return ""
}

func (d defaultLogger) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (d defaultLogger) write(c *color.Color, msg string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// NewClientLogger initializes client logger
func NewClientLogger() log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }
