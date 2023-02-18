package main

import "fmt"

/*
规定1和A对应、2和B对应、3和C对应...26和Z对应
那么一个数字字符串比如"111”就可以转化为:
"AAA"、"KA"和"AK"
给定一个只有数字字符组成的字符串str，返回有多少种转化结果
*/

// str只含有数字字符0~9
// 返回多少种转化方案
func number(str string) int {
	if str == "" || len(str) == 0 {
		return 0
	}
	return process([]byte(str), 0)
}

// str[0..i-1]转化无需过问
// str[i.....]去转化，返回有多少种转化方法
func process(str []byte, i int) int {
	if len(str) == i {
		return 1
	}
	if str[i] == '0' { // 之前的决定有问题
		return 0
	}
	ways := process(str, i+1) // 可能性一，i单转
	if i+1 < len(str) && (str[i]-'0')*10+str[i+1]-'0' < 27 {
		ways += process(str, i+2)
	}
	return ways
}

/*
动态规划
*/
func dp(str string) int {
	if str == "" || len(str) == 0 {
		return 0
	}
	s := []byte(str)
	dps := make([]int, len(s)+1)
	dps[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if str[i] != '0' {
			ways := dps[i+1]
			if i+1 < len(str) && (str[i]-'0')*10+str[i+1]-'0' < 27 {
				ways += dps[i+2]
			}
			dps[i] = ways
		}
	}
	return dps[0]
}
func main() {
	fmt.Println(number("11111"))
	fmt.Println(dp("11111"))
}
