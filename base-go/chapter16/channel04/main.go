package main

import (
	"fmt"
)

func main() {
	// 声明 初始化
	var intchan chan int
	intchan = make(chan int, 3)

	// chan 的类型
	fmt.Printf("intchan=%v 地址=%p\n", intchan, &intchan)

	// 向管道写入数据
	intchan <- 10
	num := 211
	intchan <- num
	intchan <- 50

	// 查看管道的长度和 cap（容量：不可以动态增加）
	fmt.Printf("channel len=%v cap=%v\n", len(intchan), cap(intchan))

	// 管道读取数据
	var num2 int
	num2 = <-intchan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intchan), cap(intchan))

	num3 := <-intchan
	num4 := <-intchan
	fmt.Println("num3=", num3, "num4=", num4)
}
