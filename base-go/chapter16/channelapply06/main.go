package main

import (
	"fmt"
)

// write data
func writeData(intchan chan int) {
	for i := 1; i <= 50; i++ {
		intchan <- i
		fmt.Println("write=", i)
	}
	close(intchan)
}

// read data
func readData(intchan chan int, exitchan chan bool) {
	for {
		// builtin func close 的写法
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// readData 读取完成后  即任务完成
	exitchan <- true
	close(exitchan)
}

func main() {
	// 创建两个管道
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)

	go writeData(intchan)
	go readData(intchan, exitchan)

	for {
		_, ok := <-exitchan
		// 写 ok 或者 ！ok 都可以运行
		if ok {
			break
		}
	}
}
