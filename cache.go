package goroom

import (
	"github.com/gomodule/redigo/redis"
	"sync"
)

type Cache struct {
	redis.Conn
}

var cache *Cache
var cache_once sync.Once

func GetCacheInstance() *Cache {
	cache_once.Do(func() {
		cache = new(Cache)
	})
	return cache
}

func (c *Cache) Boot(args ...interface{}) func(*GoRoom) {
	argsLength := len(args)
	if argsLength==0 {
		panic("cache's args need")
	}
	return func(srv *GoRoom) {
		// 这一步是为了确保单例初始化
		c = GetCacheInstance()
		// redis链接参数设置

		c1, err := redis.DialURL(args[0].(string))
		c.Conn = c1
		if err != nil {
			panic(err.Error())
		}

		//defer c.Conn.Close()

		srv.Cache = c
	}
}