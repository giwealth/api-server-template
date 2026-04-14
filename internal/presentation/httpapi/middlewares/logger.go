package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"
	log "api-service-template/pkg/logger"
	"api-service-template/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	commonLogger "gitlab.haochang.tv/gopkg/logger"
)

type currentUserIDKeyType string

const (
	currentUserIDKey currentUserIDKeyType = "UserID"
)

var logger *logrus.Logger

// 初始化日志
func init() {
	reqLogger, err := commonLogger.NewLogger(commonLogger.HTTPRequestV1)
	if err != nil {
		panic(errors.WithStack(err))
	}
	if f, ok := reqLogger.Formatter.(*commonLogger.HTTPRequestV1Formatter); ok {
		f.Service = os.Getenv("SERVICE")
		f.Environment = os.Getenv("ENV")
		f.TimeLayout = "2006-01-02T15:04:05.000-07:00"
	}

	// 环境变量LOG_TO_SYSLOG=1，开启日志输出到syslog
	if os.Getenv("LOG_TO_SYSLOG") == "1" {
		if err := log.HookSyslog(reqLogger, fmt.Sprintf("%s-request", os.Getenv("SERVICE"))); err != nil {
			panic(errors.WithStack(err))
		}
	}

	// 环境变量LOG_TO_STDOUT=0, 关闭日志输出到标准输出
	if os.Getenv("LOG_TO_STDOUT") == "0" {
		if err := log.HookStdOut(reqLogger); err != nil {
			panic(errors.WithStack(err))
		}
	}

	logger = reqLogger
}

// RequestLogger 记录请求日志
func RequestLogger(ctx *gin.Context) {
	t := time.Now()

	reqForLog := makeBodyReadableReq(ctx)

	ctx.Next()
	LogRequest(ctx, reqForLog, t)
}

// LogRequest 请求日志
func LogRequest(c *gin.Context, req *http.Request, start time.Time) {
	respWriter := c.Writer

	fields := logrus.Fields{
		"request":       req,
		"user":          getRequestUser(c),
		"errno":         0,
		"executionTime": time.Since(start).Milliseconds(),
		"runtimeInfo": logrus.Fields{
			"lang": runtime.Version(),
			"pid":  os.Getpid(),
		},
	}
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		fields["hostname"] = hostname
	}
	reqErr := getRequestError(c)
	if reqErr != nil {
		fields["error"] = reqErr
	}

	resp := logrus.Fields{
		"status": respWriter.Status(),
	}
	statusCode := respWriter.Status()
	fields["statusCode"] = statusCode
	fields["response"] = resp

	HTTPRequestV1Logger().WithFields(fields).Info("api request")

	isAlertError := statusCode == http.StatusInternalServerError &&
		errors.Cause(reqErr) != context.Canceled
	if isAlertError {
		log.DefaultEntry().WithFields(logrus.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"statusCode": statusCode,
			"error":      reqErr,
		}).Error("api request error")
	}
}

func getRequestUser(ctx *gin.Context) string {
	// 读取保存的用户ID
	return CurrentUserID(ctx)
}

// CurrentUserID 当前的登录的用户ID
func CurrentUserID(ctx context.Context) string {
	user := ctx.Value(currentUserIDKey)
	if user == nil {
		return ""
	}
	
	return fmt.Sprintf("%d", user.(*domain.User).ID)
}

const requestErrorKey = "requestError"

// SaveRequestError 保存请求错误
func SaveRequestError(c *gin.Context, err error) {
	errs := getRequestError(c)
	if errs == nil {
		errs = err
	} else {
		errs = multierror.Append(errs, err)
	}
	c.Set(requestErrorKey, errs)
}

func getRequestError(c *gin.Context) error {
	value, ok := c.Get(requestErrorKey)
	if !ok {
		return nil
	}
	return value.(error)
}

func makeBodyReadableReq(c *gin.Context) *http.Request {
	r := c.Request
	r2 := r
	if r.Body != nil {
		r2 = r.Clone(c)
		var b bytes.Buffer
		_, err := b.ReadFrom(r.Body)
		if err == nil {
			_ = r.Body.Close()
			r.Body = ioutil.NopCloser(&b)
			r2.Body = ioutil.NopCloser(bytes.NewReader(b.Bytes()))
		}
	}

	return r2
}

// HTTPRequestV1Logger 请求日志
func HTTPRequestV1Logger() *logrus.Logger {
	return logger
}
