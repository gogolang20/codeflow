package main

import "fmt"

/*
给定一个正整数N，表示有N份青草统一堆放在仓库里
有一只牛和一只羊，牛先吃，羊后吃，它俩轮流吃草
不管是牛还是羊，每一轮能吃的草量必须是：
1，4，16，64…(4的某次方)
谁最先把草吃完，谁获胜
假设牛和羊都绝顶聪明，都想赢，都会做出理性的决定
根据唯一的参数N，返回谁会赢
*/

// 如果n份草，最终先手赢，返回"先手"
// 如果n份草，最终后手赢，返回"后手"
func winner1(n int) string {
	if n < 5 {
		if n == 0 || n == 2 {
			return "后手"
		} else {
			return "先手"
		}
	}
	base := 1
	for base <= n {
		if winner1(n-base) == "后手" {
			return "先手"
		}
		if base > n/4 { // 防止base*4之后溢出
			break
		}
		base *= 4
	}
	return "后手"
}

func winner2(n int) string {
	if n%5 == 0 || n%5 == 2 {
		return "后手"
	} else {
		return "先手"
	}
}

func main() {
	num := 10
	fmt.Println("1: ",winner1(num))
	fmt.Println("2: ",winner2(num))
	for i := 0; i <= num; i++ {
		fmt.Println(winner2(i))
	}
}
