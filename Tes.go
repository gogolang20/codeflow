package main

import (
	"fmt"
)

func main() {
	Solution(100)
}

func Solution(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
		} else if enable_print == 1 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}

/*
提供由N个城市组成的基础设施 编号从1到N 以及他们之间的M条双向道路 道路从起点到终点之间不会相交
对于由道路直接连接到每队城市 让我们将其网络排名定义为与两个城市中到任何一个连接到道路总数
给定两个数组A，B 由每个M整数和一个整数N组成 其中 A[i]  和 B[i] 是第i条道路两端的城市
返回整个基础设施中的最大网络排名
*/
