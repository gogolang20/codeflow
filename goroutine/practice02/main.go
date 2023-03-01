package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用扇出模式同时生成100个随机数。让每个goroutine生成一个随机数，并通过缓冲channel将
该数字返回给主goroutine。设置缓冲区channel的大小，以便永远不会发送阻塞。不要分配比您
需要的更多的缓冲区。让主goroutine显示它收到的每个随机数，然后终止程序。
*/
const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	values := make(chan int, goroutines)

	for gr := 0; gr < goroutines; gr++ {
		go func() {
			values <- rand.Intn(1000)
		}()
	}

	wait := goroutines
	var nums []int
	for wait > 0 {
		nums = append(nums, <-values)
		wait--
	}
	fmt.Println(nums)
}
