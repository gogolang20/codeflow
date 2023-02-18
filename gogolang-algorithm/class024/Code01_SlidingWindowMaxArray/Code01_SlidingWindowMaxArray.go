package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

/*
滑动内最大值和最小值的更新结构

窗口不管L还是R滑动之后，都会让窗口呈现新状况，
如何能够更快的得到窗口当前状况下的最大值和最小值？
最好平均下来复杂度能做到O(1)
利用单调双端队列！

双端队列的含义：
	如果此时我依次让窗口缩小的话，哪些位置的数会依次成为窗口内的最大值 ！！！
*/

/*
假设一个固定大小为W的窗口，依次划过arr，
返回每一次滑出状况的最大值
例如，arr = [4,3,5,4,3,3,6,7], W = 3
返回：[5,5,5,4,6,7]
*/

// 暴力的对数器方法
func right(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	res := make([]int, len(arr)-w+1)
	index := 0
	L := 0
	R := w - 1
	for R < len(arr) {
		max := arr[L]
		for i := L + 1; i <= R; i++ {
			max = Max(max, arr[i])
		}
		res[index] = max
		index++
		L++
		R++
	}
	return res
}

// 双端队列实现 ！！！
func getMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	// qmax 窗口最大值的更新结构    放下标
	qmax := list.New() // 窗口最大值的更新结构
	res := make([]int, 0)
	for R := 0; R < len(arr); R++ {
		for qmax.Len() > 0 && arr[qmax.Back().Value.(int)] <= arr[R] {
			qmax.Remove(qmax.Back())
		}
		qmax.PushBack(R)                     // 从尾部加入
		if qmax.Front().Value.(int) == R-w { // 窗口过期位置下标
			qmax.Remove(qmax.Front()) // 从头弹出
		}
		if R >= w-1 { // 是否形成一个正常的窗口
			res = append(res, arr[qmax.Front().Value.(int)])
		}
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTime := 10000
	maxLen := 100
	maxValue := 200
	arr1 := generateRandomArray(maxLen, maxValue)
	arr2 := make([]int, maxLen)
	copy(arr2, arr1)
	w := rand.Intn(20)

	for i := 0; i < testTime; i++ {
		res1 := right(arr1, w)
		res2 := getMaxWindow(arr2, w)
		for index := range res1 {
			for res1[index] != res2[index] {
				fmt.Println("Oops!!!")
				break
			}
		}
	}
	fmt.Println("Finish")
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateRandomArray(maxLen, maxValue int) []int {

	res := make([]int, 0)
	for i := 0; i < maxLen; i++ {
		res = append(res, rand.Intn(maxValue))
	}
	return res
}
