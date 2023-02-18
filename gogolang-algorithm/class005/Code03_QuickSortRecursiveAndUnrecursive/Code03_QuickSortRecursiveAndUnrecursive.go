package main

func netherlandsFlag(arr []int, L int, R int) []int {
	if L > R {
		return nil
	}
	if L == R {
		return []int{L, R}
	}
	less := L - 1      // < 区 右边界
	more := R          // > 区 左边界
	index := L         // 遍历到的数
	for index < more { // 当前位置，不能和 >区的左边界撞上
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			less++
			index++
		} else {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		}
	}
	arr[more], arr[R] = arr[R], arr[more] // <[R]   =[R]   >[R]
	return []int{less + 1, more}
}
