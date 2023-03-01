package main

import (
	"fmt"
	"time"
)

// 第三个综合应用
func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 判断素数函数
func pimeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for 循环
	// var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		// 判断 num 是否是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// 将数放入
			primeChan <- num
		}
	}

	fmt.Println("有一个primeChan 取不到数据退出了")
	// 向退出的管道写入 true
	exitChan <- true

}

func main() {
	// 声明三个管道
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) // 放入结果
	exitChan := make(chan bool, 8)

	start := time.Now().UnixNano()
	fmt.Println("start time=", start)
	// 开启一个协程 向intChan 放入数据
	go putNum(intChan)
	// 开启4个协程 输出数据  判断是否为素数
	for i := 0; i < 8; i++ {
		go pimeNum(intChan, primeChan, exitChan)
	}

	// 主线程除了
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		end := time.Now().UnixNano()
		fmt.Println("end time=", end)
		fmt.Println("使用 time=", end-start)
		// 取出4个结果
		close(primeChan)

	}()

	// 遍历 pimeNum 吧结果去除
	for {
		// res, ok := <-primeChan
		_, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		// fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main 线程退出")
}
