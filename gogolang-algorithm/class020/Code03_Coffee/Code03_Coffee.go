package main

import (
	"container/heap"
	"fmt"
)

/*
给定一个数组arr，arr[i]代表第i号咖啡机泡一杯咖啡的时间
给定一个正数N，表示N个人等着咖啡机泡咖啡，每台咖啡机只能轮流泡咖啡
只有一台咖啡机，一次只能洗一个杯子，时间耗费a，洗完才能洗下一杯
每个咖啡杯也可以自己挥发干净，时间耗费b，咖啡杯可以并行挥发
假设所有人拿到咖啡之后立刻喝干净，
返回从开始等到所有咖啡机变干净的最短时间
三个参数：int[] arr、int N，int a、int b
*/

//timePoint + workTime 的小根堆
type Machine struct {
	timePoint int
	workTime  int
}

type Machines []*Machine

func (mc Machines) Len() int { return len(mc) }
func (mc Machines) Less(i, j int) bool {
	return mc[i].timePoint+mc[i].workTime > mc[j].timePoint+mc[j].workTime
}
func (mc Machines) Swap(i, j int) {
	mc[i], mc[j] = mc[j], mc[i]
}
func (mc *Machines) Push(x interface{}) {
	*mc = append(*mc, x.(*Machine))
}
func (mc *Machines) Pop() interface{} {
	old := *mc
	n := len(old)
	machine := old[n-1]
	*mc = old[0 : n-1]
	return machine
}

/*
方法一
*/

func minTime2(arr []int, n, a, b int) int {
	var mc = &Machines{}
	heap.Init(mc)
	for i := 0; i < len(arr); i++ {
		mach := &Machine{0, arr[i]}
		heap.Push(mc, mach)
	}
	drinks := make([]int, n)
	for i := 0; i < n; i++ {
		cur := heap.Pop(mc).(*Machine)
		cur.timePoint += cur.workTime
		drinks[i] = cur.timePoint
		heap.Push(mc, cur)
	}
	return bestTimeDp(drinks, a, b)
}

func bestTimeDp(drinks []int, wash, air int) int {
	N := len(drinks)
	maxFree := 0
	for i := 0; i < len(drinks); i++ {
		maxFree = max(maxFree, drinks[i]) + wash
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, maxFree+1)
	}
	for index := N - 1; index >= 0; index-- {
		for free := 0; free <= maxFree; free++ {
			selfClean1 := max(drinks[index], free) + wash
			if selfClean1 > maxFree {
				break // 因为后面的也都不用填了
			}
			// index号杯子 决定洗
			restClean1 := dp[index+1][selfClean1]
			p1 := max(selfClean1, restClean1)
			// index号杯子 决定挥发
			selfClean2 := drinks[index] + air
			restClean2 := dp[index+1][free]
			p2 := max(selfClean2, restClean2)
			dp[index][free] = min(p1, p2)
		}
	}
	return dp[0][0]
}

/*
方法二
*/
func minTime1(arr []int, n, a, b int) int {
	var mc = &Machines{}
	heap.Init(mc)
	for i := 0; i < len(arr); i++ {
		mach := &Machine{0, arr[i]}
		heap.Push(mc, mach)
	}
	drinks := make([]int, n)
	for i := 0; i < n; i++ {
		cur := heap.Pop(mc).(*Machine)
		cur.timePoint += cur.workTime
		drinks[i] = cur.timePoint
		heap.Push(mc, cur)
	}
	return bestTime(drinks, a, b, 0, 0)
}

// drinks 所有杯子可以开始洗的时间
// wash 单杯洗干净的时间（串行）
// air 挥发干净的时间(并行)
// free 洗的机器什么时候可用
// drinks[index.....]都变干净，最早的结束时间（返回）
func bestTime(drinks []int, wash, air, index, free int) int {
	if index == len(drinks) {
		return 0
	}
	// index号杯子 决定洗
	selfClean1 := max(drinks[index], free) + wash
	restClean1 := bestTime(drinks, wash, air, index+1, selfClean1)
	p1 := max(selfClean1, restClean1)
	// index号杯子 决定挥发
	selfClean2 := drinks[index] + air
	restClean2 := bestTime(drinks, wash, air, index+1, free)
	p2 := max(selfClean2, restClean2)
	return min(p1, p2)
}

func main() {
	fmt.Println(minTime2([]int{1, 3, 7}, 20, 3, 5))
	fmt.Println(minTime1([]int{1, 3, 7}, 20, 3, 5))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
