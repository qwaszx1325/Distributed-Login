package dlerr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internal "example.com/simple-login/pkg/dlerr/gen"
	"google.golang.org/grpc/status"
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

func (e *DlError) WithData(data any) *DlError {
	e.data = data
	return e
}

func (e *DlError) Is(target error) bool {
	t, ok := target.(*DlError)
	if !ok {
		return false
	}
	return e.code == t.code
}

func (e *DlError) Data() any {
	return e.data
}

func (e *DlError) WithSource(err error) *DlError {
	e.sources = append(e.sources, err)
	return e
}

func FromGrpcErr(err error) (dlErr *DlError, ok bool) {
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	// Check if the error is our custom DlError
	for _, detail := range st.Details() {
		if proto, ok := detail.(*internal.DlErrorProto); ok {
			dlErr, err := fromProto(proto)
			if err != nil {
				return nil, false
			}
			return dlErr, true
		}
	}

	return nil, false
}

func (e *DlError) toProto() (*internal.DlErrorProto, error) {
	dataBytes, err := json.Marshal(e.data)
	if err != nil {
		return nil, err
	}

	sources := make([]string, len(e.sources))
	for i, src := range e.sources {
		if src != nil {
			sources[i] = src.Error()
		}
	}

	return &internal.DlErrorProto{
		Code:    int32(e.code),
		Message: e.msg,
		Data:    dataBytes,
		Source:  sources,
	}, nil
}

func fromProto(proto *internal.DlErrorProto) (*DlError, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(proto.Data, &data); err != nil {
		return nil, err
	}

	sources := make([]error, len(proto.Source))
	for i, src := range proto.Source {
		sources[i] = errors.New(src)
	}

	return &DlError{
		code:    DlCode(proto.Code),
		msg:     proto.Message,
		data:    data,
		sources: sources,
	}, nil
}
