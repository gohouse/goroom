package goroom

import (
	"testing"
)

func GetGoRoomInstance() *GoRoom {
	return New()
}

func TestLogger(t *testing.T) {
	GetGoRoomInstance().Use(GetLoggerInstance().Boot())

	GetGoRoomInstance().Logger.(*Logger).Info("logger success !")
}
