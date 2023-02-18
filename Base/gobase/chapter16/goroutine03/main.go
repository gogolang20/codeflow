package main

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个函数 计算各个数的阶乘 放入到一个 map 中
// 全局 map
var (
	myMap = make(map[int]int, 10)
	// lock 是一个全局的互斥锁
	// synchornized 同步
	// mutex 互斥的 是一个结构体
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 将结果放入到 map 中
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {

	// 开启多个协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 2)

	lock.Lock()
	// 输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%v\n", i, v)
	}
	lock.Unlock()
}
