package main

// 最优解 时间复杂度：O(N)
func maxLengthAwesome(arr []int, k int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	minSums := make([]int, len(arr))
	minSumEnds := make([]int, len(arr))
	minSums[len(arr)-1] = arr[len(arr)-1]
	minSumEnds[len(arr)-1] = len(arr) - 1
	for i := len(arr) - 2; i >= 0; i-- {
		if minSums[i+1] < 0 {
			minSums[i] = arr[i] + minSums[i+1]
			minSumEnds[i] = minSumEnds[i+1]
		} else {
			minSums[i] = arr[i]
			minSumEnds[i] = i
		}
	}
	end := 0 // 迟迟扩不进来那一块儿的开头位置
	sum := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		// while循环结束之后：
		// 1) 如果以i开头的情况下，累加和<=k的最长子数组是arr[i..end-1]，看看这个子数组长度能不能更新res；
		// 2) 如果以i开头的情况下，累加和<=k的最长子数组比arr[i..end-1]短，更新还是不更新res都不会影响最终结果；
		for end < len(arr) && sum+minSums[end] <= k {
			sum += minSums[end]
			end = minSumEnds[end] + 1
		}
		ans = Max(ans, end-i)
		if end > i { // 还有窗口，哪怕窗口没有数字 [i~end) [4,4)
			sum -= arr[i]
		} else { // i == end,  即将 i++, i > end, 此时窗口概念维持不住了，所以end跟着i一起走
			end = i + 1
		}
	}
	return ans
}
