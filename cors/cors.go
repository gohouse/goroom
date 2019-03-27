package cors

import (
	"github.com/gohouse/goroom"
	"sync"
)

type Options struct {
	Database []interface{}
	Cache string
}

var grDefault *goroom.Engin
var gr_once sync.Once

func Default(conf ...interface{}) *goroom.Engin {
	gr_once.Do(func() {
		// 初始化框架
		grDefault = goroom.New()
		// 添加驱动
		grDefault.Use(GetRoutingInstance().Boot()). // 驱动路由
			Use(GetOrmInstance().Boot(conf[0],conf[1])). // 驱动orm
			Use(GetViewInstance().Boot()). // 驱动view模板
			Use(GetLoggerInstance().Boot()). // 驱动logger之logrus
			Use(GetCacheInstance().Boot(conf[2])) // 驱动cache之redis
		//&redis.Options{
		//	Addr:     "localhost:6379",
		//	Password: "", // no password set
		//	DB:       0,  // use default DB
		//})
	})

	return grDefault
}

func DefaultRouting(gr *goroom.Engin) *Routing {
	return gr.Routing.(*Routing)
}

func DefaultOrm(gr *goroom.Engin) *Orm {
	return gr.Orm.(*Orm)
}

func DefaultLogger(gr *goroom.Engin) *Logger {
	return gr.Logger.(*Logger)
}

func DefaultView(gr *goroom.Engin) *View {
	return gr.View.(*View)
}

func DefaultCache(gr *goroom.Engin) *Cache {
	return gr.Cache.(*Cache)
}
