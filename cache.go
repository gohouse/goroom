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

func NewCache() *Cache {
	cache_once.Do(func() {
		cache = new(Cache)
	})
	return cache
}

func DefaultCache(args ...interface{}) Handler {
	argsLength := len(args)
	if argsLength == 0 {
		panic("cache's args need")
	}
	return func(srv *GoRoom) {
		// 这一步是为了确保单例初始化
		cache = NewCache()
		// redis链接参数设置

		c1, err := redis.DialURL(args[0].(string))
		cache.Conn = c1
		if err != nil {
			panic(err.Error())
		}
	}
}
