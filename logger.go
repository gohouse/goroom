package goroom

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Logger struct {
	*logrus.Logger
}

var logger *Logger
var logger_once sync.Once

func NewLogger() *Logger {
	logger_once.Do(func() {
		logger = new(Logger)
	})
	return logger
}

func DefaultLogger() Handler {
	return func(gr *GoRoom) {
		logger = NewLogger()
		logger.Logger = logrus.New()
	}
}
