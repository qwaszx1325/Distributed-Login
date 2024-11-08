package dlerr

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// DlCode is a custom error code type.
// It consists of 7 digits: the first 3 represent the HTTP standard status code,
// and the last 4 are our custom error codes.
type DlCode int

const (
	// status code from here
	OK = 200_2000 // 成功

	// Invalid Argument

	// 400 status code
	AccountError         = 400_0000 // 帳號錯誤
	AccountPasswordError = 400_0001 // 密碼錯誤
	// 403 status code from here
	Forbidden = 403_0000 // 禁止訪問
	// 404 status code
	ResponseNotFound = 404_0000 // 沒有Response
	ResourceNotFound = 404_0001 // 找不到資源
	// 409 status code from here
	Conflict        = 409_0000 // 衝突
	ResourceIsExist = 409_0001 // 資源已存在
	// 429 status code from here
	TooManyRequests = 429_0000 // 請求過多
	// 500 status code
	InternalServerError = 500_0000 // 內部錯誤
	// 501 status code from here
	NotImplemented = 501_0000 // 功能未實現
)

func (c DlCode) HttpCode() int {

	// get the http code from the DlCode
	httpCode := int(c) / 10000

	// check if the code is valid
	if http.StatusText(httpCode) == "" {
		return http.StatusInternalServerError
	}
	return httpCode
}

// GrpcCode returns the corresponding gRPC status code.
func (c DlCode) GrpcCode() codes.Code {
	switch c.HttpCode() {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusNotImplemented:
		return codes.Unimplemented
	default:
		return codes.Unknown
	}
}

// Int returns the integer value of the KgsCode.
func (c DlCode) Int() int {
	return int(c)
}
