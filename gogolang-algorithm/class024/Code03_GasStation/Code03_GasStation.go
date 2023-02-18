package main

import "container/list"

/*
加油站的良好出发点问题
*/

// 测试链接：https://leetcode.com/problems/gas-station
// https://leetcode.cn/problems/gas-station/submissions/  中文版测试通过
func canCompleteCircuit(gas, cost []int) int {
	good := goodArray(gas, cost)
	for i := 0; i < len(gas); i++ {
		if good[i] {
			return i
		}
	}
	return -1
}

func goodArray(g, c []int) []bool {
	N := len(g)
	M := N << 1
	arr := make([]int, M)

	for i := 0; i < N; i++ {
		arr[i] = g[i] - c[i]
		arr[i+N] = g[i] - c[i]
	}
	for i := 1; i < M; i++ {
		arr[i] += arr[i-1]
	}
	w := list.New()
	for i := 0; i < N; i++ {
		for w.Len() > 0 && arr[w.Back().Value.(int)] >= arr[i] {
			w.Remove(w.Back())
		}
		w.PushBack(i)
	}
	ans := make([]bool, N)
	offset := 0
	for i, j := 0, N; j < M; {
		if arr[w.Front().Value.(int)]-offset >= 0 {
			ans[i] = true
		}
		if w.Front().Value.(int) == i {
			w.Remove(w.Front())
		}
		for w.Len() > 0 && arr[w.Back().Value.(int)] >= arr[j] {
			w.Remove(w.Back())
		}
		w.PushBack(j)
		offset = arr[i]
		i++
		j++
	}
	return ans
}
