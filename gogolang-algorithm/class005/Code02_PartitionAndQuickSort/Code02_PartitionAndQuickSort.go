package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
荷兰国旗问题
用x 划分数组
时间复杂度O(N) 不可以使用辅助数组
<=x | >x

<x | =x= | >x
等于时直接跳下一个
大于时，当前数与大于区前一个交换，大于区向左扩，当前数下标不动
如果没有目标数，以数组最后一个数做目标

快速排序 1.0
由荷兰国旗 <=x | >x 推导，划分区域
时间复杂度O(N^2)

快速排序 2.0
<x | =x= | >x 划分区域
时间复杂度O(N^2)

快速排序 3.0 随机快排 ！！！
时间复杂度 O(N * logN)
*/

// 快速排序2.0
func quickSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L int, R int) {
	if L >= R {
		return
	}
	rand.Seed(time.Now().UnixNano())
	ex := rand.Intn(R - L)                // 必须用变量保存随机数
	arr[L+ex], arr[R] = arr[R], arr[L+ex] // 加上随机交换，变成快速排序3.0
	equalArea := netherlandsFlag(arr, L, R)
	process(arr, L, equalArea[0]-1)
	process(arr, equalArea[1]+1, R)
}

// https://www.cnblogs.com/nima/p/12724868.html
// 荷兰国旗问题
// <x | =x= | >x
// 返回等于数组最后一个数的小标集合
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

// 冒泡排序
func Bubble(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func main() {
	maxSize := 100
	maxValue := 1000
	arr := make([]int, maxSize)
	arr1 := make([]int, maxSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxSize; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	copy(arr1, arr)

	quickSort(arr)
	Bubble(arr1)
	// Insertion(arr1)

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr1[i] {
			fmt.Println("ERROR")
		}
	}
	fmt.Println("Finish")
}
