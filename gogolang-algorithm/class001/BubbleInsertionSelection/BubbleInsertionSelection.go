package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 资源地址
// https://github.com/algorithmzuo

// 冒泡排序
func BubbleSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// 插入排序
func InsertSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

// 选择排序
func SelcetSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 0; i < len(arr)-1; i++ { // 比较的次数
		index := i
		for j := i + 1; j < len(arr); j++ { // 每次比较遍历的范围
			if arr[index] > arr[j] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	// fmt.Println(arr)
}

// 希尔排序
func ShellSort1(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for gap := 4; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for pre := i; pre > gap-1 && arr[pre] < arr[pre-gap]; pre -= gap { // 从小到大
				arr[pre], arr[pre-gap] = arr[pre-gap], arr[pre]
			}
		}
	}
}

// 改进间隔的希尔排序
func ShellSort2(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	// 关于 knuth 序列  GetMinStack=1  GetMinStack=3*GetMinStack+1
	h := 1
	for h <= len(arr) {
		h = h*3 + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len(arr); i++ {
			for pre := i; pre > gap-1 && arr[pre] < arr[pre-gap]; pre -= gap {
				arr[pre], arr[pre-gap] = arr[pre-gap], arr[pre]
			}
		}
	}
}

// 对数器雏形
func main() {
	/*
		对数器：
		随机生成的数组 测试随机的次数 （大样本随机测试）
		测试数组需要拷贝一份，不在同一地址空间
	*/
	testTime := 100000 // 测试次数
	maxSize := 1000    // 数组长度
	maxValue := 1000   // 数组值的范围
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		newArr := make([]int, maxSize)
		copy(newArr, arr)

		BubbleSort(newArr)
		// InsertSort(arr)
		// ShellSort2(arr)
		SelcetSort(arr)

		for j := 0; j < maxSize; j++ {
			if newArr[j] != arr[j] {
				succeed = false
				break
			}
		}
	}

	if succeed == false {
		fmt.Println("Oops")
	} else {
		fmt.Println("Success")
	}
}

// 生成随机数组
func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, maxSize)
	for index, _ := range arr {
		arr[index] = rand.Intn(maxValue)
	}
	return arr
}
