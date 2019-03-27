package cors

import (
	"github.com/gohouse/goroom"
	"github.com/sirupsen/logrus"
	"sync"
)

type Logger struct {
	*logrus.Logger
}

var logger *Logger
var logger_once sync.Once

func GetLoggerInstance() *Logger {
	logger_once.Do(func() {
		logger = new(Logger)
	})
	return logger
}

func (l *Logger) Boot(args ...interface{}) func(*goroom.Engin) {
	return func(srv *goroom.Engin) {
		// 初始化数据库链接
		// 这一步是为了确保单例初始化
		l = GetLoggerInstance()
		l.Logger = logrus.New()
		srv.Logger = l
	}
}
