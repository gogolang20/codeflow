package main

import (
	"fmt"
	"math"
	"time"
)

/*
N皇后问题是指在N*N的棋盘上要摆N个皇后，
要求任何两个皇后不同行、不同列， 也不在同一条斜线上
给定一个整数n，返回n皇后的摆法有多少种。n=1，返回1
n=2或3，2皇后和3皇后问题无论怎么摆都不行，返回0
n=8，返回92
*/

func num1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n) // index表示行 value表示列
	return process1(0, record, n)
}

// 当前来到i行，一共是0~(N-1)行
// 在i行上放皇后，所有列都尝试
// 必须要保证跟之前所有的皇后不打架
// int[] record record[x] = y 之前的第x行的皇后，放在了y列上
// 返回：不关心i以上发生了什么，i.... 后续有多少合法的方法数
func process1(i int, record []int, n int) int {
	if i == n { // 之前找到了一种有效的解
		return 1
	}
	res := 0
	// i行的皇后，放哪一列呢？j列，
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process1(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	// 0..i-1
	for k := 0; k < i; k++ {
		//同列 || 共斜线
		if j == record[k] || math.Abs(float64(record[k]-j)) == math.Abs(float64((i-k))) {
			return false
		}
	}
	return true
}

/*
方法二
*/
// 请不要超过32皇后问题
func num2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	// 如果你是13皇后问题，limit 最右13个1，其他都是0
	limit := 0 //使用 位 表示 n 的数量
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1
	}
	return process2(limit, 0, 0, 0)
}

// 7皇后问题
// limit : 0....0 1 1 1 1 1 1 1
// 之前皇后的列影响：colLim
// 之前皇后的左下对角线影响：leftDiaLim
// 之前皇后的右下对角线影响：rightDiaLim
func process2(limit, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == limit {
		return 1
	}
	// pos中所有是1的位置，是你可以去尝试皇后的位置
	pos := limit & (^(colLim | leftDiaLim | rightDiaLim))
	mostRightOne := 0
	res := 0
	for pos != 0 {
		mostRightOne = pos & (^pos + 1) // 提取最右侧的 1
		pos = pos - mostRightOne
		res += process2(limit, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1, (rightDiaLim|mostRightOne)>>1)
	}
	return res
}

func main() {
	test := 14
	t1 := time.Now()
	fmt.Println(num1(test))
	fmt.Println(time.Since(t1))

	t2 := time.Now()
	fmt.Println(num2(test))
	fmt.Println(time.Since(t2))
}
