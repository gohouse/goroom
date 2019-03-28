package goroom

import (
	"github.com/gohouse/gorose"
	"sync"
)

type Orm struct {
	*gorose.Connection
}

var orm *Orm
var orm_once sync.Once

func GetOrmInstance() *Orm {
	orm_once.Do(func() {
		orm = new(Orm)
	})
	return orm
}

func (o *Orm) Boot(args ...interface{}) func(*GoRoom) {
	return func(srv *GoRoom) {
		// 初始化数据库链接
		connection, err := gorose.Open(args...)
		if err != nil {
			panic(err)
		}

		// 这一步是为了确保单例初始化
		o = GetOrmInstance()
		o.Connection = connection
		srv.Orm = o
	}
}
