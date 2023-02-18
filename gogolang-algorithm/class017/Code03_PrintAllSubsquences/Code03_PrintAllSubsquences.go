package main

import (
	"fmt"
)

/*
打印一个字符串的全部子序列
*/

func subs(s string) []string {
	str := []byte(s)
	path := ""
	ans := make([]string, 0)
	process1(str, 0, &ans, path)
	return ans
}

// str 固定参数
// 来到了str[index]字符，index是位置
// str[0..index-1]已经走过了！之前的决定，都在path上
// 之前的决定已经不能改变了，就是path
// str[index....]还能决定，之前已经确定，而后面还能自由选择的话，
// 把所有生成的子序列，放入到ans里去
func process1(str []byte, index int, ans *[]string, path string) {
	if index == len(str) {
		*ans = append(*ans, path)
		return
	}
	// 没有要index位置的字符
	process1(str, index+1, ans, path)
	// 要了index位置的字符
	process1(str, index+1, ans, path+string(str[index]))
}

func main() {
	str1 := "abc"
	fmt.Println(subs(str1))

	str2 := "accccc"
	mp := subsNoRepeat(str2)
	for key := range mp {
		fmt.Printf("%v ", key)
	}
	fmt.Println()
}

/*
打印一个字符串的全部子序列，要求不要出现重复字面值的子序列
*/
func subsNoRepeat(s string) map[string]struct{} {
	str := []byte(s)
	path := ""
	ans := make(map[string]struct{})
	process2(str, 0, ans, path)
	return ans
}

func process2(str []byte, index int, ans map[string]struct{}, path string) {
	if index == len(str) {
		ans[path] = struct{}{}
		return
	}
	process2(str, index+1, ans, path)
	process2(str, index+1, ans, path+string(str[index]))
}
