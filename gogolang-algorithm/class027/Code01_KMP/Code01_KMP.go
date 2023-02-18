package main

import "fmt"

/*
假设字符串str长度为N，字符串match长度为M，M <= N

想确定str中是否有某个子串是等于match的。

时间复杂度O(N)
*/

/*
1）如何理解next数组

2）如何利用next数组加速匹配过程，优化时的两个实质！（私货解释）
*/

/*
给定两棵二叉树的头节点head1和head2

想知道head1中是否有某个子树的结构和head2完全一样
*/

// 返回的是第一次能找到 字符串开始的下标
func getIndexOf(s1, s2 string) int {
	if s1 == "" || s2 == "" || len(s1) < len(s2) {
		return -1
	}
	str1 := []byte(s1)
	str2 := []byte(s2)
	x, y := 0, 0
	// O(M) m <= n
	next := getNextArray(str2)
	// O(N)
	for x < len(str1) && y < len(str2) {
		if str1[x] == str2[y] {
			x++
			y++
		} else if next[y] == -1 { // y == 0
			x++ // s1 与 s2 的第一个byte 不相等
		} else {
			y = next[y]
		}
	}
	if y == len(str2) {
		return x - y
	} else {
		return -1
	}
}

func getNextArray(str2 []byte) []int {
	if len(str2) == 1 {
		return []int{-1}
	}
	next := make([]int, len(str2))
	next[0] = -1 // 人为规定的
	next[1] = 0
	i := 2  // 目前在哪个位置上求next数组的值
	cn := 0 // 当前是哪个位置的值再和i-1位置的字符比较
	for i < len(next) {
		if str2[i-1] == str2[cn] { // 配成功的时候
			next[i] = cn + 1
			i++
			cn++
		} else if cn > 0 { // cn 还能往左跳
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

func main() {
	s1 := "abcdef"
	s2 := "def"
	fmt.Println(getIndexOf(s1, s2))
}
