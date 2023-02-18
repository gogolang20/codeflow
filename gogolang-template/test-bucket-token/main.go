package main

import (
	"fmt"
	"sync"
	"time"
)

// 桶令牌限流器
type TokenLimiter struct {
	limit float64 // 速率
	burst int     // 桶大小

	mu     sync.Mutex //
	tokens float64    // 桶里token 数量
	last   time.Time  // 上一次消耗令牌的时间
}

func NewTokenLimiter(limit float64, burst int) *TokenLimiter {
	return &TokenLimiter{limit: limit, burst: burst}
}

func (lim *TokenLimiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

func (lim *TokenLimiter) AllowN(now time.Time, n int) bool {
	lim.mu.Lock()
	defer lim.mu.Unlock()
	// 计算上次请求补充了多少 token
	delta := now.Sub(lim.last).Seconds() * lim.limit
	lim.tokens += delta

	if lim.tokens > float64(lim.burst) {
		lim.tokens = float64(lim.burst)
	}

	if lim.tokens < float64(n) {
		return false
	}

	lim.tokens -= float64(n)
	lim.last = now
	return true

}

func main() {
	limiter := NewTokenLimiter(3, 5)
	for {
		n := 4
		for i := 0; i < n; i++ {
			go func(i int) {
				if !limiter.Allow() {
					fmt.Printf("forbid [%d]\n", i)
				} else {
					fmt.Printf("allow [%d]\n", i)
				}
			}(i)
		}
		time.Sleep(time.Second)
		fmt.Println("===========================")
	}
}
