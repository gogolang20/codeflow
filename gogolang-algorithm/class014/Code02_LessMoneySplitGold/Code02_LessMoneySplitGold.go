package main

import (
	"container/heap"
	"fmt"
)

//要实现优先级队列
//https://studygolang.com/pkgdoc

/*
一块金条切成两半，是需要花费和长度数值一样的铜板的。
比如长度为20的金条，不管怎么切，都要花费20个铜板。 一群人想整分整块金条，怎么分最省铜板?
例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为60，金条要分成10，20，30三个部分。

如果先把长度60的金条分成10和50，花费60; 再把长度50的金条分成20和30，花费50;一共花费110铜板。
但如果先把长度60的金条分成30和30，花费60;再把长度30金条分成10和20， 花费30;一共花费90铜板。
输入一个数组，返回分割的最小代价。
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

func lessMoney2(items IntHeap) int {
	heap.Init(&items)
	sum, cur := 0, 0
	for len(items) > 1 {
		cur = heap.Remove(&items, 0).(int) + heap.Remove(&items, 0).(int) //先弹出两个数
		sum += cur
		heap.Push(&items, cur)
	}
	return sum
}
func main() {
	items := IntHeap{10, 20, 30}
	fmt.Println(lessMoney2(items))
}
