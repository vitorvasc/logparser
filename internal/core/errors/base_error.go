package errors

type BaseError interface {
	error
	GetCode() string
	GetMessage() string
}

type baseError struct {
	Code    string
	Message string
}

func (e *baseError) Error() string {
	return e.Message
}

func (e *baseError) GetCode() string {
	return e.Code
}

func (e *baseError) GetMessage() string {
	return e.Message
}

func NewError(code string, message string) BaseError {
	return &baseError{
		Code:    code,
		Message: message,
	}
}
