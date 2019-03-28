package goroom

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Routing struct {
	*gin.Engine
}

var routing *Routing
var routing_once sync.Once

func NewRouting() *Routing {
	routing_once.Do(func() {
		routing = new(Routing)
	})
	return routing
}

func DefaultRouting() func(*GoRoom) {
	return func(srv *GoRoom) {
		// 这一步是为了确保单例初始化
		routing = NewRouting()
		routing.Engine = gin.Default()
	}
}
