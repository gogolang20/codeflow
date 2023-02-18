package main

/*
给定一个数组 arr，代表一排有分数的气球。每打爆一个气球都能获得分数，假设打爆气 球 的分数为 X，获得分数的规则如下:
	1)如果被打爆气球的左边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为 L;如果被打爆气球的右边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为 R。 获得分数为 L*X*R。
	2)如果被打爆气球的左边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为 L;如果被打爆气球的右边所有气球都已经被打爆。获得分数为 L*X。
	3)如果被打爆气球的左边所有的气球都已经被打爆;如果被打爆气球的右边有没被打爆的 气球，找到离被打爆气球最近的气球，假设分数为 R;如果被打爆气球的右边所有气球都 已经 被打爆。获得分数为 X*R。
	4)如果被打爆气球的左边和右边所有的气球都已经被打爆。获得分数为 X。
目标是打爆所有气球，获得每次打爆的分数。通过选择打爆气球的顺序，可以得到不同的总分，请返回能获得的最大分数。
【举例】
arr = {3,2,5} 如果先打爆3，获得3*2;再打爆2，获得2*5;最后打爆5，获得5;
	最后总分21 如果先打爆3，获得3*2;再打爆5，获得2*5;最后打爆2，获得2;
	最后总分18 如果先打爆2，获得3*2*5;再打爆3，获得3*5;最后打爆5，获得5;
	最后总分50 如果先打爆2，获得3*2*5;再打爆5，获得3*5;最后打爆3，获得3;
	最后总分48 如果先打爆5，获得2*5;再打爆3，获得3*2;最后打爆2，获得2;
	最后总分18 如果先打爆5，获得2*5;再打爆2，获得3*2;最后打爆3，获得3;最后总分19 返回能获得的最大分数为50
*/

// 本题测试链接 : https://leetcode.cn/problems/burst-balloons/

// 超出时间限制
func maxCoins0(arr []int) int {
	// [3,2,1,3]
	// [1,3,2,1,3,1]
	N := len(arr)
	help := make([]int, N+2)
	for i := 0; i < N; i++ {
		help[i+1] = arr[i]
	}
	help[0] = 1
	help[N+1] = 1
	return funcIn(help, 1, N)
}

// L-1位置，和R+1位置，永远不越界，并且，[L-1] 和 [R+1] 一定没爆呢！
// 返回，arr[L...R]打爆所有气球，最大得分是什么
func funcIn(arr []int, L, R int) int {
	if L == R {
		return arr[L-1] * arr[L] * arr[R+1]
	}
	// 尝试每一种情况，最后打爆的气球，是什么位置
	// L...R
	// L位置的气球，最后打爆
	max := funcIn(arr, L+1, R) + arr[L-1]*arr[L]*arr[R+1]
	// R位置的气球，最后打爆
	max = Max(max, funcIn(arr, L, R-1)+arr[L-1]*arr[R]*arr[R+1])
	// 尝试所有L...R，中间的位置，(L,R)
	for i := L + 1; i < R; i++ {
		// i位置的气球，最后打爆
		left := funcIn(arr, L, i-1)
		right := funcIn(arr, i+1, R)
		last := arr[L-1] * arr[i] * arr[R+1]
		cur := left + right + last
		max = Max(max, cur)
	}
	return max
}

// 超出时间限制
func maxCoins1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	N := len(arr)
	help := make([]int, N+2)
	help[0] = 1
	help[N+1] = 1
	for i := 0; i < N; i++ {
		help[i+1] = arr[i]
	}
	return process(help, 1, N)
}

// 打爆arr[L..R]范围上的所有气球，返回最大的分数
// 假设arr[L-1]和arr[R+1]一定没有被打爆
func process(arr []int, L, R int) int {
	if L == R { // 如果arr[L..R]范围上只有一个气球，直接打爆即可
		return arr[L-1] * arr[L] * arr[R+1]
	}
	// 最后打爆arr[L]的方案，和最后打爆arr[R]的方案，先比较一下
	max := Max(arr[L-1]*arr[L]*arr[R+1]+process(arr, L+1, R), arr[L-1]*arr[R]*arr[R+1]+process(arr, L, R-1))
	// 尝试中间位置的气球最后被打爆的每一种方案
	for i := L + 1; i < R; i++ {
		max = Max(max, arr[L-1]*arr[i]*arr[R+1]+process(arr, L, i-1)+process(arr, i+1, R))
	}
	return max
}

// 测试通过
func maxCoins2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	N := len(arr)
	help := make([]int, N+2)
	help[0] = 1
	help[N+1] = 1
	for i := 0; i < N; i++ {
		help[i+1] = arr[i]
	}
	dp := make([][]int, N+2)
	for i := range dp {
		dp[i] = make([]int, N+2)
	}
	for i := 1; i <= N; i++ {
		dp[i][i] = help[i-1] * help[i] * help[i+1]
	}
	for L := N; L >= 1; L-- {
		for R := L + 1; R <= N; R++ {
			ans := help[L-1]*help[L]*help[R+1] + dp[L+1][R]
			ans = Max(ans, help[L-1]*help[R]*help[R+1]+dp[L][R-1])
			for i := L + 1; i < R; i++ {
				ans = Max(ans, help[L-1]*help[i]*help[R+1]+dp[L][i-1]+dp[i+1][R])
			}
			dp[L][R] = ans
		}
	}
	return dp[1][N]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
