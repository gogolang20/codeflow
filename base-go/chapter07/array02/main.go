package main

import (
	"fmt"
)

func main() {
	// 从终端循环的输入5个成绩 保存输出
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值", i+1)
		fmt.Scanln(&score[i])
	}

	// 遍历数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}

	// 数组的4种初始化方法
	// 1
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	// 2
	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02=", numArr02)

	// 3 ...是固定的写法
	var numArr03 = [...]int{1, 2, 3}
	fmt.Println("numArr03=", numArr03)

	// 4
	var numArr04 = [...]int{1: 800, 2: 900, 0: 400}
	fmt.Println("numArr04=", numArr04)

	// 类型推导
	numArr05 := [...]int{1: 800, 2: 900, 0: 400}
	fmt.Println("numArr05=", numArr05)

}
