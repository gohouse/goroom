package goroom

import (
	"fmt"
	"sync"
)

var VersionNo = "0.1"
var VersionText = "GoRoom is a freely plugged Web framework with component types"
var Version = fmt.Sprintf("\n%s\n%s",VersionNo, VersionText)

type Engin struct {
	Routing IGoRoom
	Orm IGoRoom
	View IGoRoom
	Logger IGoRoom
	Cache IGoRoom
}

var gr *Engin
var once sync.Once

func New() *Engin {
	once.Do(func() {
		gr = new(Engin)
	})
	return gr
}

// Use : 驱动中间件
func (e *Engin) Use(options ...func(*Engin)) *Engin {
	for _, option := range options {
		option(e)
	}
	return e
}
