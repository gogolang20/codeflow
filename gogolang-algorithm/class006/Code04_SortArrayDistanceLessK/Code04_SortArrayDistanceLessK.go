package main

import (
	"container/heap"
	"fmt"
)

/*
已知一个几乎有序的数组。几乎有序是指，如果把数组排好顺序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。

请选择一个合适的排序策略，对这个数组进行排序。

解题思路：
使用小根堆
设置 K+1 的 heapSize
不断弹出最小值，在加入一个值 --> 循环
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortedArrDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}
	// 默认小根堆
	h := IntHeap(arr)
	heap.Init(&h)
	index := 0
	// 0...K-1
	for ; index <= Min(len(arr)-1, k-1); index++ {
		heap.Push(&h, arr[index])
	}
	i := 0
	for ; index < len(arr); i++ {
		heap.Push(&h, arr[index])
		arr[i] = heap.Pop(&h).(int)
		index++
	}
	for len(h) != 0 {
		arr[i] = heap.Pop(&h).(int)
		i++
	}
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
