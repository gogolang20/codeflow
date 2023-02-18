package main

import "fmt"

/*
给定两个长度都为N的数组weights和values，
weights[i]和values[i]分别代表 i号物品的重量和价值。
给定一个正数bag，表示一个载重bag的袋子，
你装的物品不能超过这个重量。
返回你能装下最多的价值是多少?
*/

// 所有的货，重量和价值，都在w和v数组里
// 为了方便，其中没有负数
// bag背包容量，不能超过这个载重
// 返回：不超重的情况下，能够得到的最大价值
func maxValue(w, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	return process(w, v, 0, bag)
}

// index 0~N
// rest 负~bag
func process(w, v []int, index, rest int) int {
	if rest < 0 {
		return -1
	}
	if index == len(w) {
		return 0
	}
	p1 := process(w, v, index+1, rest) //不要当前位置的货
	p2 := 0
	next := process(w, v, index+1, rest-w[index]) //要当前位置的货
	if next != -1 {                               //装不下上次的货
		p2 = v[index] + next
	}
	return max(p1, p2)
}

/*
动态规划
*/
func dp(w, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	dps := make([][]int, len(w)+1)
	for i := range dps {
		dps[i] = make([]int, bag+1)
	}
	for index := len(w) - 1; index >= 0; index-- {
		for rest := 0; rest <= bag; rest++ {
			p1 := dps[index+1][rest]
			p2, next := 0, 0
			if rest-w[index] < 0 {
				next = -1
			} else {
				next = dps[index+1][rest-w[index]]
			}
			if next != -1 {
				p2 = v[index] + next
			}
			dps[index][rest] = max(p1, p2)
		}
	}
	return dps[0][bag]
}

func main() {
	weights := []int{3, 2, 4, 7, 3, 1, 7}
	values := []int{5, 6, 3, 19, 12, 4, 2}
	bag := 15
	fmt.Println(maxValue(weights, values, bag))
	fmt.Println(dp(weights, values, bag))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
