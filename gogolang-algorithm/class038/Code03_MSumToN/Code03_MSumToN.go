package main

import "fmt"

/*
定义一种数：可以表示成若干（数量>1）连续正数和的数
比如:
5 = 2+3，5就是这样的数
12 = 3+4+5，12就是这样的数
1不是这样的数，因为要求数量大于1个、连续正数和
2 = 1 + 1，2也不是，因为等号右边不是连续正数
给定一个参数N，返回是不是可以表示成若干连续正数和的数
*/

func isMSum1(num int) bool {
	for start := 1; start <= num; start++ {
		sum := start
		for j := start + 1; j <= num; j++ {
			if sum+j > num {
				break
			}
			if sum+j == num {
				return true
			}
			sum += j
		}
	}
	return false
}

func isMSum2(num int) bool {
	//return !(num == (num & -num))
	return (num & (num - 1)) != 0
}

func main() {
	test := 7
	fmt.Println(isMSum2(test))
}
