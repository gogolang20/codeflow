package main

import "fmt"

/*
打印一个字符串的全部排列
*/

//优化版本
func permutation2(s string) []string {
	ans := make([]string, 0)
	str := []byte(s)
	process2(str, 0, &ans)
	return ans
}

func process2(str []byte, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
	} else {
		for i := index; i < len(str); i++ {
			str[i], str[index] = str[index], str[i]
			process2(str, index+1, ans)
			str[i], str[index] = str[index], str[i] //恢复现场
		}
	}
}

func main() {
	str1 := "acc"
	fmt.Println(permutation2(str1))

	str2 := "acc"
	fmt.Println(permutation3(str2))
}

/*
打印一个字符串的全部排列，要求不要出现重复的排列
*/
func permutation3(s string) []string {
	ans := make([]string, 0)
	str := []byte(s)
	process3(str, 0, &ans)
	return ans
}

func process3(str []byte, index int, ans *[]string) {
	if index == len(str) {
		*ans = append(*ans, string(str))
	} else {
		var visited [256]bool
		for i := index; i < len(str); i++ {
			if !visited[str[i]] { //如果有过 byte 来过这个位置，跳过。比用set去重好
				visited[str[i]] = true
				str[i], str[index] = str[index], str[i]
				process3(str, index+1, ans)
				str[i], str[index] = str[index], str[i] //恢复现场
			}
		}
	}
}
