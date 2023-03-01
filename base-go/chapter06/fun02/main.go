package main

import (
	"fmt"
)

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test n1=", n1) // n1=11
}

// 一个函数 getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum=", sum) // 30
	// 当函数有return时 就是将结果返回给调用者
	// 即谁调用我 就返回给谁
	return sum
	// 或者写成
	// return n1 + n2
}

// 返回数的差和和
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {

	n1 := 10
	// 调用test函数
	test(n1)
	fmt.Println("main n1=", n1) // n1=10

	sum := getSum(10, 20)
	fmt.Println("main sum=", sum) // 30

	// 调用getSumAndSub函数
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("main res1=%v res2=%v\n", res1, res2)

	// 使用_忽略返回值
	_, res3 := getSumAndSub(5, 2)
	fmt.Println("main res3=", res3)
}
