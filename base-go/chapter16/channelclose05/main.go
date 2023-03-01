package main

import (
	"fmt"
)

func main() {
	intchan := make(chan int, 3)
	intchan <- 100
	intchan <- 200
	close(intchan)

	// 这时不能写入数据
	// intchan <- 300
	fmt.Println("ok!")
	n1 := <-intchan
	fmt.Println("n1=", n1)

	// 遍历管道
	intchan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intchan2 <- i * 2 // 放入100个数据
	}

	// 遍历
	// 不可以使用for循环 需要使用for range
	// 先关闭
	close(intchan2)
	for v := range intchan2 {
		fmt.Println("v=", v)
	}

}
