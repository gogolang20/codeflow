package main

// leetcode 464题
// https://leetcode.cn/problems/can-i-win/

// 测试超时
// 1~choose 拥有的数字
// total 一开始的剩余
// 返回先手会不会赢
func canIWin0(choose, total int) bool {
	if total == 0 {
		return true
	}
	if (choose * (choose + 1) >> 1) < total {
		return false
	}
	arr := make([]int, choose)
	for i := 0; i < choose; i++ {
		arr[i] = i + 1
	}
	// arr[i] != -1 表示arr[i]这个数字还没被拿走
	// arr[i] == -1 表示arr[i]这个数字已经被拿走
	// 集合，arr，1~choose
	return process(arr, total)
}

// 当前轮到先手拿，
// 先手只能选择在arr中还存在的数字，
// 还剩rest这么值，
// 返回先手会不会赢
func process(arr []int, rest int) bool {
	if rest <= 0 {
		return false
	}
	// 先手去尝试所有的情况
	for i := 0; i < len(arr); i++ {
		if arr[i] != -1 {
			cur := arr[i]
			arr[i] = -1
			next := process(arr, rest-cur)
			arr[i] = cur
			if !next {
				return true
			}
		}
	}
	return false
}

// 测试超时
func canIWin1(choose, total int) bool {
	if total == 0 {
		return true
	}
	if (choose * (choose + 1) >> 1) < total {
		return false
	}
	return process1(choose, 0, total)
}

// 当前轮到先手拿，
// 先手可以拿1~choose中的任何一个数字
// status   i位如果为0，代表没拿，当前可以拿
//          i位为1，代表已经拿过了，当前不能拿
// 还剩rest这么值，
// 返回先手会不会赢
func process1(choose, status, rest int) bool {
	if rest <= 0 {
		return false
	}
	for i := 1; i <= choose; i++ {
		if ((1 << i) & status) == 0 { // i 这个数字，是此时先手的决定！
			if !process1(choose, (status | (1 << i)), rest-i) {
				return true
			}
		}
	}
	return false
}

// 测试通过
func canIWin2(choose, total int) bool {
	if total == 0 {
		return true
	}
	if (choose * (choose + 1) >> 1) < total {
		return false
	}
	dp := make([]int, 1<<(choose+1))
	// dp[status] == 1  true
	// dp[status] == -1  false
	// dp[status] == 0  process(status) 没算过！去算！
	return process2(choose, 0, total, dp)
}

func process2(choose, status, rest int, dp []int) bool {
	if dp[status] != 0 {
		if dp[status] == 1 {
			return true
		} else {
			return false
		}
	}
	ans := false
	if rest > 0 {
		for i := 1; i <= choose; i++ {
			if ((1 << i) & status) == 0 {
				if !process2(choose, (status | (1 << i)), rest-i, dp) {
					ans = true
					break
				}
			}
		}
	}
	if ans {
		dp[status] = 1
	} else {
		dp[status] = -1
	}
	return ans
}
