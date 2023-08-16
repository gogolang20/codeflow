package cache_server

import (
	"strconv"
	"time"

	"codeflow/template/localCache/cache"
)

// cache server
type cacheServer struct {
	memCache cache.Cache
}

func NewcacheServer() *cacheServer {
	return &cacheServer{
		memCache: cache.NewMemCache(),
	}
}

func (cs *cacheServer) SetMaxMemory(size int, unit string) bool {
	return cs.memCache.SetMaxMemory(strconv.Itoa(size) + " " + unit)
}

func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := time.Duration(0)
	if len(expire) > 0 {
		expireTs = expire[0]
	}
	return cs.memCache.Set(key, val, expireTs)
}

// 根据 key 获取 value
func (cs *cacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

func (cs *cacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

func (cs *cacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

func (cs *cacheServer) Flush() bool {
	return cs.memCache.Flush()
}

func (cs *cacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
