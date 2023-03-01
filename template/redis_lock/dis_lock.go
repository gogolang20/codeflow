package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

/*
缓存击穿 穿透 雪崩问题
一旦发生，大量请求直接到DB，压力太大
需要分布式锁 限制请求量，获取锁之后更新过期到缓存
然后删除key 释放锁
其他服务则回一直重试 获取锁

redis 实现分布式锁 From <go语言高级编程>

1 what 分布式锁
特征
- 独占排他使用: setnx
- 防止死锁发生:
	设置锁过期时间
	锁不可重入: 可重入性
- 原子性:
	获取锁和设置key过期时间: set key value ex 3 nx
	判断和释放锁之间 也需要原子性: lua脚本
- 防误删key:
	先判断 再删除
- 可重入性
- 自动续期:

2 why 分布式锁
	多个服务之间操作共享资源

3 how 分布式锁
	setnx 设置锁
	del	释放锁
	重试 获取锁
*/

func Version1(lockKey string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*2)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// 释放锁
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
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

// 防误删key and 判断和删除的原子性
func Version2(lockKey string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	counterKey := "counter"
	uid := uuid.NewV4().String() // 防误删 value

	// 不断尝试获取锁
	for {
		resp := client.SetNX(lockKey, uid, time.Second*2) // 设置过期时间
		lockSuccess, err := resp.Result()
		if err != nil {
			fmt.Println("set nx error: ", err)
			return
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
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// 释放锁
	// 实现判断和删除的原子性
	// eval "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end" 1 lock 500
	script := `if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end`
	eval, err := client.Eval(script, []string{lockKey}, uid).Result()
	if err == nil && eval.(int64) > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}

	// 没有实现判断和删除的原子性
	// checkResp := client.Get(lockKey)
	// if checkResp.String() != uid {
	// 	return
	// }
	// delResp := client.Del(lockKey)
	// unlockSuccess, err := delResp.Result()
	// if err == nil && unlockSuccess > 0 {
	// 	println("unlock success!")
	// } else {
	// 	println("unlock failed", err)
	// }
}

/*
可重入锁 and 锁自动续期
127.0.0.1:6379> EXISTS lock
127.0.0.1:6379> HSET lock uuid 1
127.0.0.1:6379> HGET lock uuid
127.0.0.1:6379> HEXISTS lock uid
127.0.0.1:6379> HINCRBY lock uuid 1
127.0.0.1:6379> HGET lock uuid
127.0.0.1:6379> EXPIRE lock 30
127.0.0.1:6379> TTL lock

加锁
1 判断锁是否存在， 否则直接获取锁: EXISTS key
2 如果锁存在则判断是否自己的锁, 自己的锁则重入: HINCRBY key uuid 1
3 否则重试
redis:
eval "if redis.call('exists', KEYS[1]) == 0 or redis.call('hexists', KEYS[1], ARGV[1]) == 1 then redis.call('hincrby', KEYS[1], ARGV[1], 1) redis.call('expire', KEYS[1], ARGV[2]) return 1 else return 0 end" 1 lock uuid 30

if redis.call('exists', KEYS[1]) == 0 or redis.call('hexists', KEYS[1], ARGV[1]) == 1
then
redis.call('hincrby', KEYS[1], ARGV[1], 1)
redis.call('expire', KEYS[1], ARGV[2])
return 1
else
return 0
end

解锁
1 判断自己的锁是否存在(hexists), 不存在返回nil
2 如果自己的锁存在, 则减1, 判断减1后的值是否为0, 为0则释放锁(del)则返回1
3 不为0, 返回0
redis:
eval "if redis.call('hexists', KEYS[1], ARGV[1]) == 0 then return nil elseif redis.call('hincrby', KEYS[1], ARGV[1], -1) == 0 then return redis.call('del', KEYS[1]) else return 0 end" 1 lock uid

if redis.call('hexists', KEYS[1], ARGV[1]) == 0
then
return nil
elseif redis.call('hincrby', KEYS[1], ARGV[1], -1) == 0
then
return redis.call('del', KEYS[1])
else
return 0
end
*/
func Version3(lockKey string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	counterKey := "counter"
	uid := uuid.NewV4().String()

	for {
		add_lock := `if redis.call('exists', KEYS[1]) == 0 or redis.call('hexists', KEYS[1], ARGV[1]) == 1 then redis.call('hincrby', KEYS[1], ARGV[1], 1) redis.call('expire', KEYS[1], ARGV[2]) return 1 else return 0 end`
		lockSuccess, err := client.Eval(add_lock, []string{lockKey}, uid, 3).Result()
		if err != nil {
			fmt.Println("set lock error: ", err)
			return
		}
		if lockSuccess == 1 {
			break
		}
	}

	// key 续期
	// Timer 间隔时间， and 重新设置的过期时间
	// 每睡 1/3 的锁过期时间 check 一下是否过期， 重新设置过期时间
	go func(uid string) {
		ti := time.NewTimer(1 * time.Second)
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
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)
	// ===== ===== end logic ===== =====

	// 释放锁
	release_lock := `if redis.call('hexists', KEYS[1], ARGV[1]) == 0 then return nil elseif redis.call('hincrby', KEYS[1], ARGV[1], -1) == 0 then return redis.call('del', KEYS[1]) else return 0 end`
	eval, err := client.Eval(release_lock, []string{lockKey}, uid).Result()
	if err == nil && eval.(int64) > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

/*
悲观锁（Pessimistic Lock）： 就是很悲观，每次去拿数据的时候都认为别人会修改。所以每次在拿数据的时候都会上锁。
这样别人想拿数据就被挡住，直到悲观锁被释放，悲观锁中的共享资源每次只给一个线程使用，其它线程阻塞，用完后再把资源转让给其它线程
但是在效率方面，处理加锁的机制会产生额外的开销，还有增加产生死锁的机会。另外还会降低并行性，如果已经锁定了一个线程A，
其他线程就必须等待该线程A处理完才可以处理

乐观锁（Optimistic Lock）： 就是很乐观，每次去拿数据的时候都认为别人不会修改。所以不会上锁，但是如果想要更新数据，
则会在更新前检查在读取至更新这段时间别人有没有修改过这个数据。如果修改过，则重新读取，再次尝试更新，
循环上述步骤直到更新成功（当然也允许更新失败的线程放弃操作）,乐观锁适用于多读的应用类型，这样可以提高吞吐量

说到乐观锁，就必须提到一个概念：CAS
什么是CAS呢？Compare-and-Swap，即比较并替换，也有叫做Compare-and-Set的，比较并设置。
1、比较：读取到了一个值A，在将其更新为B之前，检查原值是否仍为A（未被其他线程改动）。
2、设置：如果是，将A更新为B，结束。[1]如果不是，则什么都不做。
上面的两步操作是原子性的，可以简单地理解为瞬间完成，在CPU看来就是一步操作。
有了CAS，就可以实现一个乐观锁，允许多个线程同时读取（因为根本没有加锁操作），但是只有一个线程可以成功更新数据，
并导致其他要更新数据的线程回滚重试。 CAS利用CPU指令，从硬件层面保证了操作的原子性，以达到类似于锁的效果。

Java中真正的CAS操作调用的native方法
因为整个过程中并没有“加锁”和“解锁”操作，因此乐观锁策略也被称为无锁编程。
换句话说，乐观锁其实不是“锁”，它仅仅是一个循环重试CAS的算法而已，但是CAS有一个问题那就是会产生ABA问题，什么是ABA问题，以及如何解决呢？

ABA 问题：
如果一个变量V初次读取的时候是A值，并且在准备赋值的时候检查到它仍然是A值，那我们就能说明它的值没有被其他线程修改过了吗？
很明显是不能的，因为在这段时间它的值可能被改为其他值，然后又改回A，那CAS操作就会误认为它从来没有被修改过。这个问题被称为CAS操作的 "ABA"问题。

ABA 问题解决：
我们需要加上一个版本号（Version）,在每次提交的时候将版本号+1操作，那么下个线程去提交修改的时候，
会带上版本号去判断，如果版本修改了，那么线程重试或者提示错误信息~

原文链接：https://blog.csdn.net/qq_44625745/article/details/125364292
*/
