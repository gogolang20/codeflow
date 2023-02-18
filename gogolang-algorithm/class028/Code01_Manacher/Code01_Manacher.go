package main

import (
	"fmt"
	"math"
)

/*
Manacher算法

假设字符串str长度为N，想返回最长回文子串的长度

时间复杂度O(N)
*/

/*
Manacher算法核心

1）理解回文半径数组

2）理解所有中心的回文最右边界R，和取得R时的中心点C

3）理解   L…(i`)…C…(i)…R  的结构，以及根据i’回文长度进行的状况划分

4）每一种情况划分，都可以加速求解i回文半径的过程
*/

// 一： i 在 R 外
// 二： i 在 R 内； 1 i'<R ; 2 i'不 ; 3 i'=L压线
func manacher(s string) int {
	if s == "" {
		return 0
	}
	// "12132" -> "#1#2#1#3#2#"
	str := manacherString(s)
	pArr := make([]int, len(str)) // 回文半径的大小
	C := -1
	// 讲述中：R代表最右的扩成功的位置
	// coding：最右的扩成功位置的，再下一个位置
	R := -1
	max := math.MinInt
	for i := 0; i < len(str); i++ { // 0 1 2
		// R第一个违规的位置，i>= R
		// i位置扩出来的答案，i位置扩的区域，至少是多大。
		if R > i { // i 在 R 内
			pArr[i] = Min(pArr[2*C-i], R-i)
		} else {
			pArr[i] = 1
		}
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		max = Max(max, pArr[i])
	}
	return max - 1
}

func manacherString(str string) []byte {
	res := make([]byte, len(str)*2+1)
	charArr := []byte(str)
	index := 0
	for i := 0; i != len(res); i++ {
		if (i & 1) == 0 { // 偶数下标
			res[i] = '#'
		} else {
			res[i] = charArr[index]
			index++
		}
	}
	return res
}

// for test
func right(s string) int {
	if s == "" {
		return 0
	}
	str := manacherString(s)
	max := 0
	for i := 0; i < len(str); i++ {
		L := i - 1
		R := i + 1
		for L >= 0 && R < len(str) && str[L] == str[R] {
			L--
			R++
		}
		max = Max(max, R-L-1)
	}
	return max / 2
}

func main() {
	str := "1221"
	fmt.Println(manacher(str))
	fmt.Println(right(str))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


