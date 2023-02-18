package main

import (
	"fmt"
)

// 二分查找
// 查找 条件
// 退出条件

// 使用指针(引用类型) 比值传递速度更快
func BinaryFind(arr *[6]int, leftIndex int, rightindex int, findVal int) {
	// 判断leftIndex 是否大于 rightindex
	if leftIndex > rightindex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间下标
	// middle := (leftIndex + rightindex) / 2
	middle := leftIndex + (rightindex-leftIndex)>>1

	if (*arr)[middle] > findVal {
		// 查找的范围
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightindex, findVal)
	} else {
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}

func main() {

	var arr [6]int = [6]int{1, 9, 80, 677, 1200, 4584}

	// 测试  需要传入地址
	BinaryFind(&arr, 0, len(arr)-1, 80)

}
