package goroom

import (
	"sync"
)

type IGoRoom interface {
	Boot(args ...interface{}) func(room *GoRoom)
}

type GoRoom struct {
	Routing IGoRoom
	Orm     IGoRoom
	View    IGoRoom
	Logger  IGoRoom
	Cache   IGoRoom
}

var gr *GoRoom
var once sync.Once

func New() *GoRoom {
	once.Do(func() {
		gr = new(GoRoom)
	})
	return gr
}

// Use : 驱动中间件
func (e *GoRoom) Use(options ...func(*GoRoom)) *GoRoom {
	for _, option := range options {
		option(e)
	}
	return e
}
