package main

func maxLength(arr []int, k int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	// key : 前缀和
	// value : 0~value这个前缀和是最早出现key这个值的
	set := make(map[int]int)
	set[0] = -1 // important
	length := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if _, ok := set[sum-k]; ok {
			length = Max(i-set[sum-k], length)
		}
		if _, ok := set[sum]; !ok {
			set[sum] = i
		}
	}
	return length
}

//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
