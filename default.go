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
		grDefault.Use(DefaultRouting()). // 驱动路由
							Use(DefaultOrm(options.Database...)). // 驱动orm
							Use(DefaultView()).                   // 驱动view模板
							Use(DefaultLogger()).                 // 驱动logger之logrus
							Use(DefaultCache(options.Cache)).     // 驱动cache之redis
							Use(DefaultView())                    // 驱动模板view
	})

	return grDefault
}
