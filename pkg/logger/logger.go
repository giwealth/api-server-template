package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
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

	defaultLogger = al
}

// DefaultEntry 默认日志对象
func DefaultEntry() *logrus.Entry {
	return logrus.NewEntry(defaultLogger).WithField("channel", "default")
}
