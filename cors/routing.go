package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/goroom"
	"sync"
)

type Routing struct {
	*gin.Engine
}

var routing *Routing
var routing_once sync.Once

func GetRoutingInstance() *Routing {
	routing_once.Do(func() {
		routing = new(Routing)
	})
	return routing
}

func(r *Routing) Boot(args ...interface{}) func(*goroom.Engin) {
	return func(srv *goroom.Engin) {
		// 这一步是为了确保单例初始化
		r = GetRoutingInstance()
		r.Engine = gin.Default()
		srv.Routing = r
	}
}
