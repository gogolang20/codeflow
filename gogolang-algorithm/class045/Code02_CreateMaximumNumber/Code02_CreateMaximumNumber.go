package main

/*
给两个长度分别为M和N的整型数组nums1和nums2，其中每个值都不大于9，再给定一个正数K。
你可以在nums1和nums2中挑选数字，要求一共挑选K个，并且要从左到右挑。返回所有可能的结果中，代表最大数字的结果。
*/

// 测试链接: https://leetcode.com/problems/create-maximum-number/
// https://leetcode.cn/problems/create-maximum-number/submissions/

// 测试通过
func maxNumber1(nums1, nums2 []int, k int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	if k < 0 || k > len1+len2 {
		return nil
	}
	res := make([]int, k)
	dp1 := getdp(nums1) // 生成dp1这个表，以后从nums1中，只要固定拿N个数，
	dp2 := getdp(nums2)
	// get1 从arr1里拿的数量
	// K - get1 从arr2里拿的数量
	for get1 := Max(0, k-len2); get1 <= Min(k, len1); get1++ {
		// arr1 挑 get1个，怎么得到一个最优结果
		pick1 := maxPick(nums1, dp1, get1)
		pick2 := maxPick(nums2, dp2, k-get1)
		merge := merge(pick1, pick2)
		if !preMoreThanLast(res, 0, merge, 0) {
			res = merge
		}
	}
	return res
}

func merge(nums1, nums2 []int) []int {
	k := len(nums1) + len(nums2)
	ans := make([]int, k)
	for i, j, r := 0, 0, 0; r < k; r++ {
		if preMoreThanLast(nums1, i, nums2, j) {
			ans[r] = nums1[i]
			i++
		} else {
			ans[r] = nums2[j]
			j++
		}
	}
	return ans
}

func preMoreThanLast(nums1 []int, i int, nums2 []int, j int) bool {
	for i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
		i++
		j++
	}
	return j == len(nums2) || (i < len(nums1) && nums1[i] > nums2[j])
}

func getdp(arr []int) [][]int {
	size := len(arr)     // 0~N-1
	pick := len(arr) + 1 // 1 ~ N
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, pick)
	}
	// get 不从0开始，因为拿0个无意义
	for get := 1; get < pick; get++ { // 1 ~ N
		maxIndex := size - get
		// i~N-1
		for i := size - get; i >= 0; i-- {
			if arr[i] >= arr[maxIndex] {
				maxIndex = i
			}
			dp[i][get] = maxIndex
		}
	}
	return dp
}

func maxPick(arr []int, dp [][]int, pick int) []int {
	res := make([]int, pick)
	for resIndex, dpRow := 0, 0; pick > 0; pick-- {
		res[resIndex] = arr[dp[dpRow][pick]]
		dpRow = dp[dpRow][pick] + 1
		resIndex++
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
