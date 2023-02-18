package main

import "fmt"

/*
给定一个正数n，求n的裂开方法数，
规定：后面的数不能比前面的数小
比如4的裂开方法有：
1+1+1+1、1+1+2、1+3、2+2、4
5种，所以返回5
*/

// n为正数
func wayss(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return process(1, n)
}

// 上一个拆出来的数是pre
// 还剩rest需要去拆
// 返回拆解的方法数
func process(pre, rest int) int {
	if rest == 0 { //到达了拆分的终结点
		return 1
	}
	if pre > rest {
		return 0
	}
	ways := 0
	for first := pre; first <= rest; first++ {
		ways += process(first, rest-first)
	}
	return ways
}

func dp1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
		dp[pre][pre] = 1 //对角线位置
	}
	for pre := n - 1; pre >= 1; pre-- {
		for rest := pre + 1; rest <= n; rest++ {
			ways := 0
			for first := pre; first <= rest; first++ {
				ways += dp[first][rest-first]
			}
			dp[pre][rest] = ways
		}
	}
	return dp[1][n]
}

func dp2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for pre := 1; pre <= n; pre++ {
		dp[pre][0] = 1
		dp[pre][pre] = 1 //对角线位置
	}
	for pre := n - 1; pre >= 1; pre-- { //第0行不用
		for rest := pre + 1; rest <= n; rest++ {
			dp[pre][rest] = dp[pre+1][rest]
			dp[pre][rest] += dp[pre][rest-pre]
		}
	}
	return dp[1][n]
}

func main() {
	fmt.Println(wayss(4))
	fmt.Println(dp1(4))
	fmt.Println(dp2(4))
}
