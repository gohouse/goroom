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
- 会话: session.go  
- 缓存: cache.go  

### goroom提供对应的默认组件介绍
- 路由: 默认以 gin 为路由组件, 同时, 可以使用 gin 的所有中间件  
- 数据库: 默认以gorose为orm组件  
- 模板: 默认以官方的 `html/template` 为模板组件  
- 日志: 默认以 logrus 为日志库, 同时, 可以使用 logrus 的所有中间件  
- 会话: 默认以 gin-contrib/session 为会话库  
- 缓存: 默认以 redigo 为缓存服务  


## 自定义中间件开发规范

