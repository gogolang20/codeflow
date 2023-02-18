package main

import (
	"math"
)

/*
TSP问题 有N个城市，任何两个城市之间的都有距离，任何一座城市到自己的距离都为0。所有点到点的距 离都存在一个N*N的二维数组matrix里，
也就是整张图由邻接矩阵表示。现要求一旅行商从k城市 出发必须经过每一个城市且只在一个城市逗留一次，最后回到出发的k城，返回总距离最短的路的距离。
参数给定一个matrix，给定k。
*/

func t1(matrix [][]int) int {
	N := len(matrix) // 0...N-1
	// set
	// set.get(i) != null i这座城市在集合里
	// set.get(i) == null i这座城市不在集合里
	set := make(map[int]int)
	for i := 0; i < N; i++ {
		set[i] = 1
	}
	return func1(matrix, set, 0)
}

func func1(matrix [][]int, set map[int]int, start int) int {
	cityNum := 0
	for i := 0; i < len(set); i++ {
		if res, ok := set[i]; res != -1 && ok {
			cityNum++
		}
	}
	if cityNum == 1 {
		return matrix[start][0]
	}
	// cityNum > 1  不只start这一座城
	set[start] = -1
	min := math.MaxInt
	for i := 0; i < len(set); i++ {
		if res, ok := set[i]; res != -1 && ok {
			// start -> i i... -> 0
			cur := matrix[start][i] + func1(matrix, set, i)
			min = Min(min, cur)
		}
	}
	set[start] = 1
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
