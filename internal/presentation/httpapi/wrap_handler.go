package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	mw "api-service-template/internal/presentation/httpapi/middlewares"
	// "api-service-template/internal/presentation/httpapi/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// WrapHandler 自定义的Handler处理结构
func wrapHandler(fn interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		r, err := newHandlerRequest(fn)
		if err != nil {
			responseError(c, err, errInternal)
			return
		}

		if err := c.ShouldBind(r); err != nil {
			e := getError(err.(validator.ValidationErrors), r)
			responseError(c, e, errBadRequest)
			return
		}
		if err := c.ShouldBindUri(r); err != nil {
			e := getError(err.(validator.ValidationErrors), r)
			responseError(c, e, errBadRequest)
			return
		}
		// 请求参数的复杂验证，需在请求结构体实现Parse接口
		if parser, ok := r.(RequestParser); ok {
			if err := parser.Parse(); err != nil {
				responseError(c, err, errBadRequest)
				return
			}
		}

		// 接口超时时间
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		resp, err := call2(fn, ctx, r)
		if err != nil {
			responseError(c, err, errInternal)
			return
		}

		if w, ok := resp.(ResponseWriter); ok {
			w.Write(c.Writer)
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

func responseError(c *gin.Context, err error, defaultErr apiError) {
	mw.SaveRequestError(c, err)

	var finalErr apiError
	if e, ok := errors.Cause(err).(apiError); ok {
		finalErr = e
	} else if errors.Is(err, context.DeadlineExceeded) {
		finalErr = errTimeout
	} else if errors.Is(defaultErr, errBadRequest) {
		finalErr = newAPIError(http.StatusBadRequest, 40000, fmt.Sprintf("请求参数错误: %v", err))
	} else {
		finalErr = defaultErr
	}
	c.JSON(finalErr.Status(), finalErr)
}

// RequestParser 请求解析和验证
type RequestParser interface {
	Parse() error
}

// ResponseWriter 自定义的响应结构, 类似下发文件等，content-type不是application/json的数据等
type ResponseWriter interface {
	Write(w http.ResponseWriter)
}

// Call 调用函数
func call(fn interface{}, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("parameters of function are error")
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)
	return result, nil
}

// Call2 调用函数，该函数有2个返回值并且第二个是err
func call2(fn interface{}, params ...interface{}) (interface{}, error) {
	result, err := call(fn, params...)
	if err != nil {
		return nil, err
	}
	if len(result) != 2 {
		return nil, errors.New("returns of function are error")
	}
	if result[1].IsNil() {
		return result[0].Interface(), nil
	}
	err, ok := result[1].Interface().(error)
	if !ok {
		return nil, errors.New("should return a error for function")
	}
	return nil, err
}

// newHandlerRequest 构造Handler的请求实例
func newHandlerRequest(fn interface{}) (interface{}, error) {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}
	if f.Type().NumIn() != 2 {
		return nil, errors.New("parameters number should be two")
	}
	reqPtr := f.Type().In(1)
	if reqPtr.Kind() != reflect.Ptr {
		return nil, errors.New("request parameter must be a pointer")
	}

	req := reflect.New(reqPtr.Elem()).Interface()
	return req, nil
}
