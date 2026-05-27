package errors

type MultiError struct {
	msg    string
	Errors []error
}

func NewMultiError(msg string) *MultiError { _ = "STUB: not implemented"; return nil }

func (m *MultiError) Append(err error) { _ = "STUB: not implemented"; return }

// Flatten the multi error

func (m *MultiError) Error() string { _ = "STUB: not implemented"; return "" }

func MultiToError(e error) error { _ = "STUB: not implemented"; return nil }

func (m *MultiError) ToErr() error { _ = "STUB: not implemented"; return nil }
