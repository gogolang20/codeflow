package main

import (
	codeflow "codeflow/template/local_cache/cache-server"
	"fmt"
)

/*
内存缓存系统
1 支持设定过期时间 精度到秒
2 支持设定最大内存 内存超出时做出合适的处理
3 支持并发安全
4 按照以下接口要求实现
*/

func main() {
	cache := codeflow.NewcacheServer()
	cache.SetMaxMemory("100MB")

	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")

	fmt.Println()
	cache.Flush()
	cache.Keys()
}
