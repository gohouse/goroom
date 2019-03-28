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

func NewOrm() *Orm {
	orm_once.Do(func() {
		orm = new(Orm)
	})
	return orm
}

func DefaultOrm(args ...interface{}) Handler {
	return func(srv *GoRoom) {
		orm = NewOrm()
		// 初始化数据库链接
		connection, err := gorose.Open(args...)
		if err != nil {
			panic(err)
		}
		orm.Connection = connection
	}
}
