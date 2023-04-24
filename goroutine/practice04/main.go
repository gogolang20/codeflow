package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

/*
使用 work pool 同时生成最多100个随机数，拒绝偶数值。
如果已经收集到100个奇书，就让协程停止运行
*/

func main() {
	values := make(chan int)
	shutdown := make(chan struct{})
	poolSize := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(id int) {
			for {
				n := rand.Intn(1000)
				select {
				case values <- n:
					fmt.Printf("worker %d sent %d\n", id, n)
				case <-shutdown:
					fmt.Printf("worker %d shutting down\n", id)
					wg.Done()
					return
				}
			}
		}(i)
	}

	var nums []int
	for i := range values {
		if i%2 == 0 {
			fmt.Println("Discarding", i)
			continue
		}
		fmt.Println("Keeping", i)
		nums = append(nums, i)
		if len(nums) == 100 {
			break
		}
	}

	fmt.Println("Receiver sending shutdown signal")
	close(shutdown)

	wg.Wait()
	fmt.Printf("Result count: %d\n", len(nums))

	fmt.Println(nums)
}
