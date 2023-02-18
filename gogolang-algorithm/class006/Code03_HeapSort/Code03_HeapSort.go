package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
堆排序 ！！！（40min左右开始）
完全二叉树
    每个节点下标为 i
    左子树的下标为 2*i +1
    右子树的下标为 2*i +2
    父节点的下标为 (i -1)/2

堆可以理解成完全二叉树
大根堆：每个子树，最大值是子树的头节点 //子树：头结点往下的节点都算
小根堆

大根堆如何插入一个元素
大根堆如何删除一个最大元素
    将最后一个节点置换到头节点，删除最大元素
    调整堆的结构，左右子树较大的往上heapify

一个有序大根堆，如果有个下标的值修改了
	调用一次heapify 再调用一次heapInsert ，堆再次有序

一个几乎有序的数组排序，移动不超过K的距离，使数组有序
    建立一个K长度的小根堆，不断弹出最小值，再从后面加入新元素，再弹出最小值
*/

// 堆排序
// 时间复杂度 O(N * logN)
func heapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	//第一种建堆方式 : 调整成大根堆
	for index := range arr {
		heapInsert(arr, index)
	}
	//第二种建堆方式 : 从下往上建大根堆    更快！！！
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr)) // O(logN)
	}
	//循环将最大的数放到数组最后一位
	for heapSize := len(arr) - 1; heapSize > 0; heapSize-- {
		arr[0], arr[heapSize] = arr[heapSize], arr[0]
		heapify(arr, 0, heapSize)
	}
}

// 新加进来的数，现在停在了index位置，请依次往上移动，
// 移动到0位置，或者干不掉自己的父亲了，停！
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] { //index == 0 也会终止循环
		arr[index], arr[(index-1)>>1] = arr[(index-1)>>1], arr[index]
		index = (index - 1) / 2
	}
}

// 从index位置，往下看，不断的下沉
// 停：较大的孩子都不再比index位置的数大；已经没孩子了
func heapify(arr []int, index int, heapSize int) {
	largest := 0
	for left := index*2 + 1; left < heapSize; left = index*2 + 1 {
		//1 选出子树中较大的一个 孩子的下标
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		//2 index 值大于等于 孩子值
		if arr[index] >= arr[largest] {
			break
		}
		//3 子树值大于节点值  互换
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
	}
}

func main() {
	testTime := 50000 //测试次数
	maxSize := 100    //数组长度
	maxValue := 1000  //数组值的范围
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := make([]int, maxSize)
		copy(arr1, arr)

		heapSort(arr)
		Selection(arr1)
		//Insertion(arr1)

		for j := 0; j < maxSize; j++ {
			if arr1[j] != arr[j] {
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

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func Selection(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[index], arr[i] = arr[i], arr[index]
	}
}

func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, maxSize)
	for index := range arr {
		arr[index] = rand.Intn(maxValue)
	}
	return arr
}
