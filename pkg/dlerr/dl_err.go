package dlerr

import (
	"fmt"
	"net/http"
)

// DlError is a custom error type that wraps error code, message and error sources
type DlError struct {
	code    DlCode
	msg     string
	data    any
	sources []error
}

func New(code DlCode, msg string, source ...error) *DlError {
	return &DlError{
		code:    code,
		msg:     msg,
		sources: source,
	}
}

func (e *DlError) Error() string {
	if len(e.sources) > 0 {
		return fmt.Sprintf("dlCode: %v, msg:%s, sources: %v", e.code, e.msg, e.sources)
	}
	return fmt.Sprintf("dlCode: %v, msg: %s", e.code, e.msg)
}

func (e *DlError) HttpCode() int {

	if e == nil {
		return http.StatusInternalServerError
	}
	return e.code.HttpCode()
}
func (e *DlError) Code() DlCode {
	return e.code
}

func (e *DlError) Message() string {
	return e.msg
}

func (e *DlError) Unwrap() []error {
	return e.sources
}
