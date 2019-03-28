package goroom

import (
	"sync"
)

type Options struct {
	Database []interface{}
	Cache    string
}

var grDefault *GoRoom
var gr_once sync.Once

func (gr *GoRoom) Default(options *Options) *GoRoom {
	gr_once.Do(func() {
		// 初始化框架
		grDefault = New()
		// 添加驱动
		grDefault.Use(GetRoutingInstance().Boot()). // 驱动路由
			Use(GetOrmInstance().Boot(options.Database...)). // 驱动orm
			Use(GetViewInstance().Boot()). // 驱动view模板
			Use(GetLoggerInstance().Boot()). // 驱动logger之logrus
			Use(GetCacheInstance().Boot(options.Cache)) // 驱动cache之redis
	})

	return grDefault
}

func (gr *GoRoom) DefaultRouting() *Routing {
	return gr.Routing.(*Routing)
}

func (gr *GoRoom) DefaultOrm() *Orm {
	return gr.Orm.(*Orm)
}

func (gr *GoRoom) DefaultLogger() *Logger {
	return gr.Logger.(*Logger)
}

func (gr *GoRoom) DefaultView() *View {
	return gr.View.(*View)
}

func (gr *GoRoom) DefaultCache() *Cache {
	return gr.Cache.(*Cache)
}
