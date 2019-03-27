package cors

import (
	"github.com/gohouse/goroom"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestCache(t *testing.T) {
	gr := goroom.New().Use(GetCacheInstance().Boot("redis://localhost:6379"))

	cr := DefaultCache(gr).Conn

	//cr.Set("stra", "b", 1000000000)
	cr.Do("HSet","hkey", "field","val")

	res,err := redis.StringMap(cr.Do("HGetAll","hkey"))
	if err!=nil {
		t.Fatal("fail: ",err.Error())
	} else {
		t.Log("success: ",res)
	}
}
