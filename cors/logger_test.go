package cors

import (
	"github.com/gohouse/goroom"
	"testing"
)

func GetEnginInstance() *goroom.Engin {
	return goroom.New()
}

func TestLogger(t *testing.T) {
	GetEnginInstance().Use(GetLoggerInstance().Boot())

	GetEnginInstance().Logger.(*Logger).Info("logger success !")
}
