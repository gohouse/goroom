package goroom

import (
	"sync"
)

type GoRoom struct {
}

type Handler func(*GoRoom)

var gr *GoRoom
var once sync.Once

func New() *GoRoom {
	once.Do(func() {
		gr = new(GoRoom)
	})
	return gr
}

// Use : 驱动中间件
func (e *GoRoom) Use(options ...Handler) *GoRoom {
	for _, option := range options {
		option(e)
	}
	return e
}
