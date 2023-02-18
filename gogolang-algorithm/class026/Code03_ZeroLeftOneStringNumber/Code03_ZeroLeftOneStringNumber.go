package main

import "fmt"

/*
给定一个数N，想象只由0和1两种字符，组成的所有长度为N的字符串

如果某个字符串,任何0字符的左边都有1紧挨着,认为这个字符串达标

返回有多少达标的字符串
*/

func getNum1(n int) int {
	if n < 1 {
		return 0
	}
	return process(1, n)
}

func process(i, n int) int {
	if i == n-1 {
		return 2
	}
	if i == n {
		return 1
	}
	return process(i+1, n) + process(i+2, n)
}

func getNum3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-2)
	return 2*res[0][0] + res[1][0]
}

func matrixPower(m [][]int, p int) [][]int {
	res := make([][]int, len(m))
	for i := range res {
		res[i] = make([]int, len(m[0]))
	}
	for i := 0; i < len(res); i++ {
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = muliMatrix(res, tmp)
		}
		tmp = muliMatrix(tmp, tmp)
	}
	return res
}

func muliMatrix(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := range res {
		res[i] = make([]int, len(m2[0]))
	}

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

func main() {
	// 要求：递推式无条件转移
	N := 6
	fmt.Println(getNum1(N))
	fmt.Println(getNum3(N))

	f := 3
	fmt.Println(fi(f))
}

/*
用1*2的瓷砖，把N*2的区域填满

返回铺瓷砖的方法数
*/

func fi(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{{1, 1}, {1, 0}}
	res := matrixPower(base, n-2)
	return res[0][0] + res[1][0]
}
