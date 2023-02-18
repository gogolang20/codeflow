package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
给定3个参数，N，M，K
怪兽有N滴血，等着英雄来砍自己
英雄每一次打击，都会让怪兽流失[0~M]的血量
到底流失多少？每一次在[0~M]上等概率的获得一个值
求K次打击之后，英雄把怪兽砍死的概率
*/

func right(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := math.Pow(float64(M+1), float64(K))
	kill := process(K, M, N)
	return kill / all
}

// 怪兽还剩hp点血
// 每次的伤害在[0~M]范围上
// 还有times次可以砍
// 返回砍死的情况数！
func process(times, M, hp int) float64 {
	if times == 0 {
		if hp <= 0 {
			return 1
		} else {
			return 0
		}
	}
	if hp <= 0 {
		return math.Pow(float64(M+1), float64(times))
	}
	ways := 0.0
	for i := 0; i <= M; i++ {
		ways += process(times-1, M, hp-i)
	}
	return ways
}

func dp1(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := math.Pow(float64(M+1), float64(K))
	dp := make([][]float64, K+1)
	for i := range dp {
		dp[i] = make([]float64, N+1)
	}
	dp[0][0] = 1
	for times := 1; times <= K; times++ {
		dp[times][0] = math.Pow(float64(M+1), float64(times))
		for hp := 1; hp <= N; hp++ {
			ways := 0.0
			for i := 0; i <= M; i++ {
				if hp-i >= 0 {
					ways += dp[times-1][hp-i]
				} else {
					ways += math.Pow(float64(M+1), float64(times-1))
				}
			}
			dp[times][hp] = ways
		}
	}
	kill := dp[K][N]
	return kill / all
}

func dp2(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := math.Pow(float64(M+1), float64(K))
	dp := make([][]float64, K+1)
	for i := range dp {
		dp[i] = make([]float64, N+1)
	}
	dp[0][0] = 1
	for times := 1; times <= K; times++ {
		dp[times][0] = math.Pow(float64(M+1), float64(times))
		for hp := 1; hp <= N; hp++ {
			dp[times][hp] = dp[times][hp-1] + dp[times-1][hp]
			if hp-1-M >= 0 {
				dp[times][hp] -= dp[times-1][hp-1-M]
			} else {
				dp[times][hp] -= math.Pow(float64(M+1), float64(times-1))
			}
		}
	}
	kill := dp[K][N]
	return kill / all
}

func main() {
	rand.Seed(time.Now().UnixNano())
	NMax := 10
	MMax := 10
	KMax := 10
	testTime := 200
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		N := rand.Intn(NMax)
		M := rand.Intn(MMax)
		K := rand.Intn(KMax)
		ans1 := right(N, M, K)
		ans2 := dp1(N, M, K)
		ans3 := dp2(N, M, K)
		if ans1 != ans2 || ans1 != ans3 {
			fmt.Println("Oops!")
			break
		}
	}
	fmt.Println("测试结束")
}
