package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type memCache struct {
	// 最大内存
	maxMemorySize int64
	// 最大内存字符串表示
	maxMemorySizeStr string
	// 当前已使用内存
	currMemorySize int64
	// 缓存 key value
	values map[string]*memCacheValue
	// lock
	locker sync.RWMutex
	// 过期时间间隔
	clearExpireItemTimeInterval time.Duration
}

type memCacheValue struct {
	// value
	val interface{}
	// 过期时间
	expireTime time.Time
	// 有效时长
	expire time.Duration
	// value 大小
	size int64
}

func NewMemCache() Cache {
	mc := &memCache{
		values:                      make(map[string]*memCacheValue, 0),
		clearExpireItemTimeInterval: 1 * time.Second,
	}

	go mc.clearExpireItem()

	return mc
}

// size: 1KB 100KB 1MB 2MB 1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println("[SetMaxMemory] : ", mc.maxMemorySize, mc.maxMemorySizeStr)

	return true
}

// 将 value 写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValSize(val),
	}
	mc.del(key)
	mc.add(key, v)
	if mc.currMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println("over max size: ", mc.maxMemorySize)
	}
	return true
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currMemorySize -= tmp.size
		delete(mc.values, key)
	}

}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemorySize += val.size
}

// 根据 key 获取 value
func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	mcv, ok := mc.get(key)
	if ok {
		if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}
		return mcv.val, ok
	}
	return 0, false
}

// 删除 key
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)

	return true
}

// 判断 key 是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	_, ok := mc.values[key]
	return ok
}

// 清空所有 key
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()

	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemorySize = 0
	return true
}

// 获取缓存中所有 key 的数量
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	return int64(len(mc.values))
}

func (mc *memCache) clearExpireItem() {
	ti := time.NewTicker(mc.clearExpireItemTimeInterval)
	defer ti.Stop()

	for {
		select {
		case <-ti.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
