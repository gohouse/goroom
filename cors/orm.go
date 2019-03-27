package cors

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/goroom"
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

func (o *Orm) Boot(args ...interface{}) func(*goroom.Engin) {
	return func(srv *goroom.Engin) {
		// 初始化数据库链接
		connection,err := gorose.Open([]interface{}{"sqlite3","./db.sqlite"}...)
		if err!=nil{
			panic(err)
		}

		// 这一步是为了确保单例初始化
		o = GetOrmInstance()
		o.Connection = connection
		srv.Orm = o
	}
}
