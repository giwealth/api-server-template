package httpapi

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	errBadRequest   = newAPIError(http.StatusBadRequest, 40000, "请求参数错误")
	errNotFound     = newAPIError(http.StatusNotFound, 40004, "请求数据不存在")
	errInternal     = newAPIError(http.StatusInternalServerError, 50000, "服务器内部错误")
	errTimeout      = newAPIError(http.StatusGatewayTimeout, 50004, "请求超时")
	errUnauthorized = newAPIError(http.StatusUnauthorized, 40001, "验证失败")
	errForbidden    = newAPIError(http.StatusForbidden, 40003, "无权操作")
)

// Error 错误响应
type apiError struct {
	status  int
	Code    int    `json:"code"`
	Message string `json:"message"`

	data any
}

// NewError 构造错误
func newAPIError(status, code int, message string) apiError {
	return apiError{
		status:  status,
		Code:    code,
		Message: message,
	}
}

func (e apiError) Data() (any, bool) {
	return e.data, e.data != nil
}

// Error 实现错误接口
func (e apiError) Error() string {
	return e.Message
}

// Status 状态码
func (e apiError) Status() int {
	return e.status
}

func (e *apiError) WithData(data any) *apiError {
	e.data = data
	return e
}

// Wrap 包裹一个底层错误
func (e apiError) Wrap(err error) error {
	return e.WrapMessage(err, "")
}

// WrapMessage 包裹一个底层错误并附带提示详情
func (e apiError) WrapMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	if v, ok := errors.Cause(err).(apiError); ok {
		return v
	}
	if message == "" {
		message = fmt.Sprintf("%v", err)
	} else {
		message = fmt.Sprintf("%v %s", err, message)
	}
	return errors.Wrap(e, message)
}
