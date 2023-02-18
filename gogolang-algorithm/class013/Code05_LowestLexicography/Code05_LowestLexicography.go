package main

import (
	"fmt"
	"strings"
)

/*
贪心算法的解题套路

1，实现一个不依靠贪心策略的解法X，可以用最暴力的尝试

2，脑补出贪心策略A、贪心策略B、贪心策略C...

3，用解法X和对数器，用实验的方式得知哪个贪心策略正确

4，不要去纠结贪心策略的证明
*/

/*
给定一个由字符串组成的数组strs，
必须把所有的字符串拼接起来，
返回所有可能的拼接结果中，字典序最小的结果
*/

// "a" + "b"  Compare  "b" + "a"
func lowestString2(arr []string) string {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if strings.Compare(arr[j-1]+arr[j], arr[j]+arr[j-1]) > 0 {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	res := strings.Join(arr, "")
	return res
}

func main() {
	arr := []string{"b", "ba", "ff", "cd"}
	fmt.Println(lowestString2(arr))
	fmt.Println("a" == "a")
}
