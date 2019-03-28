## GoRoom简介
GoRoom web framework 是一个完全组件化的 go web 框架,
所有模块组件化, 全部以中间件的形式提供, 支持按需加载任何组件, 
同时你可以把任何组件替代为自定义组件.  

## 中间件介绍  

### 目前 GoRoom 提供5大常用组件:  
- 路由: Routing.go  
- 数据库: orm.go  
- 模板: view.go  
- 日志: logger.go  
- 缓存: cache.go  

### goroom提供对应的默认组件介绍
- 路由: 默认以 gin 为路由组件, 同时, 可以使用 gin 的所有中间件  
- 数据库: 默认以gorose为orm组件  
- 模板: 默认以官方的 `html/template` 为模板组件  
- 日志: 默认以 logrus 为日志库, 同时, 可以使用 logrus 的所有中间件   
- 缓存: 默认以 redigo 为缓存服务  

## 示例
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/goroom"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化 goroom 结构体
	gr := goroom.New()
	// 按需加载模块
	gr.Use(goroom.GetRoutingInstance().Boot()).	// 加载路由
		Use(goroom.GetOrmInstance().Boot("sqlite3","db.sqlite")).	// 加载数据库
		Use(goroom.GetLoggerInstance().Boot())	// 加载log

	// 路由示例
	gr.DefaultRouting().GET("/", getUserInfo)

	// log示例
	gr.DefaultLogger().Info("view http://localhost:8999 in browser")
	// 监听端口
	gr.DefaultRouting().Run(":8999")
}

func getUserInfo(c *gin.Context) {
	gr := goroom.New()
	// orm数据示例
	dbDataInit(gr)
	res, err := gr.DefaultOrm().Table("users").First()
	if err!=nil {
		// log示例
		gr.DefaultLogger().Error(err.Error())
		c.JSON(200, map[string]interface{}{"msg":err.Error()})
	} else {
		c.JSON(200, res)
	}
}

func dbDataInit(gr *goroom.GoRoom)  {
	var aff int64
	var err error
	// 建表
	aff, err = gr.DefaultOrm().Execute(`CREATE TABLE IF NOT EXISTS "users" (
	 "uid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 "name" TEXT NOT NULL,
	 "age" integer NOT NULL
)`)
	if err != nil {
		gr.DefaultLogger().Error("建表失败啊:", err.Error())
		return
	}
	// 插入数据
	aff, err = gr.DefaultOrm().Table("users").Data(map[string]interface{}{
		"name": "goroom",
		"age":  18,
	}).Insert()
	if err != nil || aff == 0 {
		gr.DefaultLogger().Error("插入数据失败啊:", err.Error())
		return
	}
}

```


## 自定义中间件开发规范

