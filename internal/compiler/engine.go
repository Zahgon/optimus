package compiler

import (
	"text/template"
	"time"
)

const (
	EntityCompiler = "compiler"

	// ISODateFormat https://en.wikipedia.org/wiki/ISO_8601
	ISODateFormat = "2006-01-02"

	ISOTimeFormat = time.RFC3339
)

// Engine compiles a set of defined macros using the provided context
type Engine struct {
	baseTemplate *template.Template
}

func NewEngine() *Engine { _ = "STUB: not implemented"; return nil }

func (e *Engine) Compile(templateMap map[string]string, context map[string]any) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) CompileString(input string, context map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
