package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Version1()
			Version2("lock")
			// expireIncr()
		}()
	}
	wg.Wait()
}

func Version1(lockKey string) {
	client := NewRedis()
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*2)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		log.Fatalln("lock result: ", err, lockSuccess)
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			log.Fatalln("set value error: ", err)
		}
	}
	println("current counter is ", cntValue)

	// 释放锁
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		fmt.Println("unlock success!")
	} else {
		fmt.Println("unlock failed", err)
	}
}

// 防误删key and 判断和删除的原子性
func Version2(lockKey string) {
	client := NewRedis()
	counterKey := "counter"
	uid := uuid.NewV4().String() // 防误删 value

	// 不断尝试获取锁
	for {
		resp := client.SetNX(lockKey, uid, time.Second*2) // 设置过期时间
		lockSuccess, err := resp.Result()
		if err != nil {
			log.Fatalln("[Version2] set nx error: ", err)
		}
		if lockSuccess {
			break
		}
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			log.Fatalln("set value error: ", err)
		}
	}
	fmt.Println("current counter is ", cntValue)

	// 释放锁
	// 实现判断和删除的原子性
	// eval "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end" 1 lock 500
	script := `if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end`
	eval, err := client.Eval(script, []string{lockKey}, uid).Result()
	if err == nil && eval.(int64) > 0 {
		fmt.Println("unlock success!")
	} else {
		fmt.Println("unlock fail: ", err)
	}
}

func Version3(lockKey string) {
	client := NewRedis()
	counterKey := "counter"
	uid := uuid.NewV4().String()

	for {
		add_lock := `if redis.call('exists', KEYS[1]) == 0 or redis.call('hexists', KEYS[1], ARGV[1]) == 1 then redis.call('hincrby', KEYS[1], ARGV[1], 1) redis.call('expire', KEYS[1], ARGV[2]) return 1 else return 0 end`
		lockSuccess, err := client.Eval(add_lock, []string{lockKey}, uid, 3).Result()
		if err != nil {
			log.Fatalln("set lock error: ", err)
		}
		if lockSuccess == 1 {
			break
		}
	}

	// key 续期
	// Timer 间隔时间， and 重新设置的过期时间
	// 每睡 1/3 的锁过期时间 check 一下是否过期， 重新设置过期时间
	go func(uid string) {
		ti := time.NewTicker(1 * time.Second)
		defer ti.Stop()
		for {
			select {
			case <-ti.C:
				// 重新设置 key 过期时间
				// eval "if redis.call('hexists', KEYS[1], ARGV[1]) == 1 then return redis.call('expire', KEYS[1], ARGV[2]) else return 0 end" 1 lock uid 2
				reSetExpireStr := "if redis.call('hexists', KEYS[1], ARGV[1]) == 1 then return redis.call('expire', KEYS[1], ARGV[2]) else return 0 end"
				expireEval, err := client.Eval(reSetExpireStr, []string{lockKey}, uid, 2).Result()
				if err != nil {
					break
				}
				if expireEval.(int64) > 0 {
					println("set success!")
				}
			}
		}
	}(uid)

	// ===== ===== start logic ===== =====
	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			log.Fatalln("set value error: ", err)
		}
	}
	println("current counter is ", cntValue)
	// ===== ===== end logic ===== =====

	// 释放锁
	release_lock := `if redis.call('hexists', KEYS[1], ARGV[1]) == 0 then return nil elseif redis.call('hincrby', KEYS[1], ARGV[1], -1) == 0 then return redis.call('del', KEYS[1]) else return 0 end`
	eval, err := client.Eval(release_lock, []string{lockKey}, uid).Result()
	if err != nil {
		println("unlock failed", err)
	}
	if eval == nil {
		println("unlock success!")
	}
	if eval.(int64) > 0 {
		println("unlock success!")
	}
	if eval.(int64) == 0 {
		println("unlock success!")
	}
}
