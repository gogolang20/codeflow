package main

func getMaxLength(arr []int, K int) int {
	if arr == nil || len(arr) == 0 || K <= 0 {
		return 0
	}
	left := 0
	right := 0
	sum := arr[0]
	length := 0
	for right < len(arr) {
		if sum == K {
			length = Max(length, right-left+1)
			sum -= arr[left]
			left++
		} else if sum < K {
			right++
			if right == len(arr) {
				break
			}
			sum += arr[right]
		} else {
			sum -= arr[left]
			left++
		}
	}
	return length
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
