package main

import (
	"fmt"
	"strings"
)

// 资源地址
// https://www.bilibili.com/video/BV1EF411h7Xq?p=10
// 无重复最长字串 -- 滑动窗口思想
func lengthOflongestSubstring(s string) int {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 && (i+1) > end {
			end = i + 1
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

// 资源地址
// https://www.bilibili.com/video/BV1EF411h7Xq?p=14
// 反转字符串
func recerse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	num := lengthOflongestSubstring("sagdsagwogqb")
	fmt.Println(num)

	str := recerse("我爱我的祖国 123.000")
	fmt.Println(str)
}
