package main

import "fmt"

/*
判断str1和str2是否是旋转字符串
*/

func isRotation(a, b string) bool {
	if a == "" || b == "" || len(a) != len(b) {
		return false
	}
	b2 := b + b
	return getIndexOf(b2, a) != -1
}

// KMP Algorithm
func getIndexOf(s, m string) int {
	if len(s) < len(m) {
		return -1
	}
	ss := []byte(s)
	ms := []byte(m)
	si := 0
	mi := 0
	next := getNextArray(ms)
	for si < len(ss) && mi < len(ms) {
		if ss[si] == ms[mi] {
			si++
			mi++
		} else if next[mi] == -1 {
			si++
		} else {
			mi = next[mi]
		}
	}
	if mi == len(ms) {
		return si - mi
	} else {
		return -1
	}
}

func getNextArray(ms []byte) []int {
	if len(ms) == 1 {
		return []int{-1}
	}
	next := make([]int, len(ms))
	next[0] = -1
	next[1] = 0
	pos := 2
	cn := 0
	for pos < len(next) {
		if ms[pos-1] == ms[cn] {
			next[pos] = cn + 1
			pos++
			cn++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[pos] = 0
			pos++
		}
	}
	return next
}

func main() {
	str1 := "yunzuocheng"
	str2 := "zuochengyun"
	fmt.Println(isRotation(str1,str2))
}
