package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
最多同时生成100个随机数。不要发送所有的100个值，因为发送/接收的数量是未知的
*/

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	values := make(chan int)
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for gr := 0; gr < goroutines; gr++ {
		go func() {
			defer wg.Done()
			n := rand.Intn(1000)
			if n%2 == 0 {
				return
			}
			values <- n
		}()
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	var nums []int
	for n := range values {
		nums = append(nums, n)
	}

	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
