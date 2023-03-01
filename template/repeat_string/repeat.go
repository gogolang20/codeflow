package main

import (
	"fmt"
	"strings"
)

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

func main() {
	num := lengthOflongestSubstring("sagdsagwogqb")
	fmt.Println(num)
}
