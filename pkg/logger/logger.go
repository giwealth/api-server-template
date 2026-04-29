package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
	sl "log/syslog"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/syslog"
	"gitlab.haochang.tv/gopkg/logger"
)

var defaultLogger *logrus.Logger

// Init 初始化日志
func init() {
	al, err := logger.NewLogger(logger.APPLogsV1)
	if err != nil {
		panic("create appLogsV1 logger error")
	}

	if l, ok := al.Formatter.(*logger.APPLogsV1Formatter); ok {
		l.Service = os.Getenv("SERVICE")
		l.Environment = os.Getenv("ENV")
		l.TimeLayout = time.RFC3339
	}
	logrus.SetFormatter(al.Formatter)

	if s := os.Getenv("LOG_LEVEL"); s != "" {
		if lvl, err := logrus.ParseLevel(s); err == nil {
			logrus.SetLevel(lvl)
		}
	}

	// 环境变量LOG_TO_SYSLOG=1，开启日志输出到syslog
	if os.Getenv("LOG_TO_SYSLOG") == "1" {
		if err := HookSyslog(al, fmt.Sprintf("%s-app", os.Getenv("SERVICE"))); err != nil {
			panic(errors.WithStack(err))
		}
	}

	// 环境变量LOG_TO_STDOUT=0, 关闭日志输出到标准输出
	if os.Getenv("LOG_TO_STDOUT") == "0" {
		if err := HookStdOut(al); err != nil {
			panic(errors.WithStack(err))
		}
	}

	defaultLogger = al
}

// HookSyslog 日志输出到syslog
func HookSyslog(l *logrus.Logger, tag string) error {
	h, err := syslog.NewSyslogHook("", "", sl.LOG_LOCAL6, tag)
	if err != nil {
		return errors.WithStack(err)
	}

	l.Hooks.Add(h)

	return nil
}

// HookStdOut 关闭日志输出到标准输出
func HookStdOut(l *logrus.Logger) error {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errors.WithStack(err)
	}

	writer := bufio.NewWriter(src)
	l.SetOutput(writer)

	return nil
}

// DefaultEntry 默认日志对象
func DefaultEntry() *logrus.Entry {
	return logrus.NewEntry(defaultLogger).WithField("channel", "default")
}