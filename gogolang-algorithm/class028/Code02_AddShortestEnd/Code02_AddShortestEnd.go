package main

import "fmt"

/*
最少添加多少byte，使string变成回文
*/

func shortestEnd(s string) string {
	if s == "" {
		return ""
	}
	str := manacherString(s)
	pArr := make([]int, len(str))
	C := -1
	R := -1
	maxContainsEnd := -1
	for i := 0; i != len(str); i++ {
		if R > i {
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
		if R == len(str) {
			maxContainsEnd = pArr[i]
			break
		}
	}
	res := make([]byte, len(s)-maxContainsEnd+1)
	for i := 0; i < len(res); i++ {
		res[len(res)-1-i] = str[i*2+1]
	}
	return string(res)
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

func main() {
	str := "abcd123321"
	fmt.Println(shortestEnd(str))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
