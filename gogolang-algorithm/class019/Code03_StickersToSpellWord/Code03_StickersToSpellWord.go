package main

import "math"

/*
给定一个字符串str，给定一个字符串类型的数组arr，出现的字符都是小写英文
arr每一个字符串，代表一张贴纸，你可以把单个字符剪开使用，目的是拼出str来
返回需要至少多少张贴纸可以完成这个任务。
例子：str= "babac"，arr = {"ba","c","abcd"}
ba + ba + c  3    abcd + abcd 2    abcd+ba 2
所以返回2
*/

// 本题测试链接：https://leetcode.com/problems/stickers-to-spell-word
func minStickers1(stickers []string, target string) int {
	ans := process1(stickers, target)
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}
func process1(stickers []string, target string) int {
	if target == "" {
		return 0
	}
	min := math.MaxInt
	for _, first := range stickers {
		rest := minus(target, first)
		if len(rest) != len(target) {
			min = mins(min, process1(stickers, rest))
		}
	}
	if min == math.MaxInt {
		return math.MaxInt
	} else {
		return min + 1
	}
}
func minus(s1, s2 string) string {
	str1 := []byte(s1)
	str2 := []byte(s2)
	count := make([]int, 26)
	for _, cha := range str1 {
		count[cha-'a']++
	}
	for _, cha := range str2 {
		count[cha-'a']--
	}
	builder := make([]byte, 0)
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			for j := 0; j < count[i]; j++ {
				builder = append(builder, byte(i+'a'))
			}
		}
	}
	return string(builder)
}

/*
方法二
*/
func minStickers2(stickers []string, target string) int {
	counts := make([][]int, len(stickers))
	for i := range counts {
		counts[i] = make([]int, 26)
	}
	for i := 0; i < len(stickers); i++ {
		str := []byte(stickers[i])
		for _, v := range str {
			counts[i][v-'a']++
		}
	}
	ans := process2(counts, target)
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}

// stickers[i] 数组，当初i号贴纸的字符统计 int[][] stickers -> 所有的贴纸
func process2(stickers [][]int, t string) int {
	if len(t) == 0 {
		return 0
	}
	target := []byte(t)
	tcounts := make([]int, 26)
	for _, v := range target {
		tcounts[v-'a']++
	}
	min := math.MaxInt
	for i := 0; i < len(stickers); i++ {
		sticker := stickers[i]
		// 最关键的优化(重要的剪枝!这一步也是贪心!)
		if sticker[target[0]-'a'] > 0 {
			builder := make([]byte, 0)
			for j := 0; j < 26; j++ {
				if tcounts[j] > 0 {
					nums := tcounts[j] - sticker[j]
					for k := 0; k < nums; k++ {
						builder = append(builder, byte(j+'a'))
					}
				}
			}
			rest := string(builder)
			min = mins(min, process2(stickers, rest))
		}
	}
	if min == math.MaxInt {
		return min
	} else {
		return min + 1
	}
}

/*
方法三
时间通过测试
非常重要的一个题目！！！
*/
func minStickers3(stickers []string, target string) int {
	counts := make([][]int, len(stickers))
	for i := range counts {
		counts[i] = make([]int, 26)
	}
	for i := 0; i < len(stickers); i++ {
		str := []byte(stickers[i])
		for _, v := range str {
			counts[i][v-'a']++
		}
	}
	dp := map[string]int{"": 0}
	ans := process3(counts, target, dp)
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}
func process3(stickers [][]int, t string, dp map[string]int) int {
	if value, ok := dp[t]; ok {
		return value
	}
	target := []byte(t)
	tcounts := make([]int, 26)
	for _, v := range target {
		tcounts[v-'a']++
	}
	min := math.MaxInt
	for i := 0; i < len(stickers); i++ {
		sticker := stickers[i]
		// 最关键的优化(重要的剪枝!这一步也是贪心!)
		if sticker[target[0]-'a'] > 0 {
			builder := make([]byte, 0)
			for j := 0; j < 26; j++ {
				if tcounts[j] > 0 {
					nums := tcounts[j] - sticker[j]
					for k := 0; k < nums; k++ {
						builder = append(builder, byte(j+'a'))
					}
				}
			}
			rest := string(builder)
			min = mins(min, process3(stickers, rest, dp))
		}
	}
	ans := 0
	if min == math.MaxInt {
		ans = min
	} else {
		ans = min + 1
	}
	dp[t] = ans
	return ans
}

func mins(a, b int) int {
	if a < b {
		return a
	}
	return b
}
