package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
内存缓存系统
1 支持设定过期时间 精度到秒
2 支持设定最大内存 内存超出时做出合适的处理
3 支持并发安全
4 按照以下接口要求实现
*/
type Cache interface {
	// size: 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size string) bool
	// 将 value 写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	// 根据 key 获取 value
	Get(key string) (interface{}, bool)
	// 删除 key
	Del(key string) bool
	// 判断 key 是否存在
	Exists(key string) bool
	// 清空所有 key
	Flush() bool
	// 获取缓存中所有 key 的数量
	keys() int64
}

type memCache struct {
	// 最大内存
	maxMemorySize    int64
	maxMemorySizeStr string
	// 当前已使用内存
	currMemorySize int64
	//
	//
	//
}

func NewMemCache() Cache {
	return &memCache{
		maxMemorySize: 0,
	}
}

// size: 1KB 100KB 1MB 2MB 1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)

	return true
}

// 将 value 写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {

	return false
}

// 根据 key 获取 value
func (mc *memCache) Get(key string) (interface{}, bool) {
	return 0, false
}

// 删除 key
func (mc *memCache) Del(key string) bool {
	return false
}

// 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
	return false
}

// 清空所有 key
func (mc *memCache) Flush() bool {
	return false
}

// 获取缓存中所有 key 的数量
func (mc *memCache) keys() int64 {
	return 0
}

func main() {
	cache := NewMemCache()
	cache.SetMaxMemory("100MB")

	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)

	// cache.Set("int", 1)
	// cache.Set("bool", false)
	// cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	cache.keys()
}

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
	case "GB":
	case "TB":
	case "PB":
	default:
		num = 0
	}
	return num, ""
}
