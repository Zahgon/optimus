package errors

type ErrorType string

func (s ErrorType) String() string { _ = "STUB: not implemented"; return "" }

const (
	ErrInternalError   ErrorType = "Internal Error"
	ErrNotFound        ErrorType = "Not Found"
	ErrAlreadyExists   ErrorType = "Resource Already Exists"
	ErrInvalidArgument ErrorType = "Invalid Argument"
	ErrFailedPrecond   ErrorType = "Failed Precondition"

	ErrInvalidState ErrorType = "Invalid State"
)

type DomainError struct {
	ErrorType  ErrorType
	Entity     string
	Message    string
	WrappedErr error
}

func (*DomainError) Is(tgt error) bool { _ = "STUB: not implemented"; return false }

// nolint

func AddErrContext(err error, entity, msg string) *DomainError {
	_ = "STUB: not implemented"
	return nil
}

func IsErrorType(err error, errType ErrorType) bool { _ = "STUB: not implemented"; return false }

func NewError(errType ErrorType, entity, msg string) *DomainError {
	_ = "STUB: not implemented"
	return nil
}

func InternalError(entity, msg string, err error) *DomainError {
	_ = "STUB: not implemented"
	return nil
}

func InvalidStateTransition(entity, msg string) *DomainError { _ = "STUB: not implemented"; return nil }

func InvalidArgument(entity, msg string) *DomainError { _ = "STUB: not implemented"; return nil }

func AlreadyExists(entity, msg string) *DomainError { _ = "STUB: not implemented"; return nil }

func NotFound(entity, msg string) *DomainError { _ = "STUB: not implemented"; return nil }

func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

func As(err error, target any) bool { _ = "STUB: not implemented"; return false }

func (e *DomainError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *DomainError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *DomainError) DebugString() string { _ = "STUB: not implemented"; return "" }

func Wrap(entity, msg string, err error) error { _ = "STUB: not implemented"; return nil }

func WrapIfErr(entity, msg string, err error) error { _ = "STUB: not implemented"; return nil }

func GRPCErr(err error, msg string) error { _ = "STUB: not implemented"; return nil }
