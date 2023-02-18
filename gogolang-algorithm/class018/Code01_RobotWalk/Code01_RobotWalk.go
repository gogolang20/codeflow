package main

import "fmt"

/*
机器人寻路问题
假设有排成一行的N个位置，记为1~N，N 一定大于或等于 2
开始时机器人在其中的M位置上(M 一定是 1~N 中的一个)
如果机器人来到1位置，那么下一步只能往右来到2位置；
如果机器人来到N位置，那么下一步只能往左来到 N-1 位置；
如果机器人来到中间位置，那么下一步可以往左走或者往右走；
规定机器人必须走 K 步，最终能来到P位置(P也是1~N中的一个)的方法有多少种
给定四个参数 N、M、K、P，返回方法数。
*/

/*
方法一 暴力递归
*/
/*
N      总的路长
start  起始的index
aim    要到达目标的index
K      剩余的步数

返回方法的数量
*/
func ways1(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return process1(start, K, aim, N)
}

// 机器人当前来到的位置是cur，
// 机器人还有rest步需要去走，
// 最终的目标是aim，
// 有哪些位置？ 1~N
// 返回：机器人从cur出发，走过rest步之后，最终停在aim的方法数，是多少？
func process1(cur, rest, aim, N int) int {
	if rest == 0 { // 已经不需要走了
		if cur == aim {
			return 1
		} else {
			return 0
		}
	}
	if cur == 1 { // 1 -> 2
		return process1(2, rest-1, aim, N)
	}
	if cur == N { // N-1 <- N
		return process1(N-1, rest-1, aim, N)
	}
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

/*
方法二 傻缓存
记忆化搜索
*/
func ways2(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	// cur 范围： 1-N
	// rest 范围： 0-K
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := 0; j < K+1; j++ {
			dp[i][j] = -1
		}
	}
	// dp就是缓存表
	// dp[cur][rest] == -1 -> process2(cur, rest)之前没算过！
	// dp[cur][rest] != -1 -> process2(cur, rest)之前算过！返回值，dp[cur][rest]
	// N+1 * K+1
	return process2(start, K, aim, N, dp)
}

func process2(cur, rest, aim, N int, dp [][]int) int {
	if dp[cur][rest] != -1 { // 之前算过
		return dp[cur][rest] // 直接返回
	}
	ans := 0
	if rest == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = process2(2, rest-1, aim, N, dp)
	} else if cur == N {
		ans = process2(N-1, rest-1, aim, N, dp)
	} else {
		ans = process2(cur-1, rest-1, aim, N, dp) + process2(cur+1, rest-1, aim, N, dp)
	}
	dp[cur][rest] = ans // 加入到缓存里
	return ans
}

/*
方法三 动态规划
有边界的杨辉三角
*/
func ways3(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	dp := make([][]int, N+1)
	for index := range dp {
		dp[index] = make([]int, K+1)
	}
	dp[aim][0] = 1                     // 其余 dp[...][0] = 0
	for rest := 1; rest <= K; rest++ { // 一行一行往右走
		dp[1][rest] = dp[2][rest-1]    // 最上单独赋值
		for cur := 2; cur < N; cur++ { // 每一列从上往下赋值
			dp[cur][rest] = dp[cur-1][rest-1] + dp[cur+1][rest-1]
		}
		dp[N][rest] = dp[N-1][rest-1] // 最下单独赋值
	}
	return dp[start][K]
}

func main() {
	fmt.Println(ways1(5, 2, 4, 6))
	fmt.Println(ways2(5, 2, 4, 6))
	fmt.Println(ways3(5, 2, 4, 6))
}
