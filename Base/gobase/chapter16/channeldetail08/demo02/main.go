package main

import (
	"fmt"
)

func main() {
	// 关于不关闭管道 取出数据
	// select
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时 如果不关闭会阻塞导致 deadlock
	// lable :
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan 读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan 读取的数据%s\n", v)
		default:
			fmt.Printf("都取不到")
			return
		}
	}

}
