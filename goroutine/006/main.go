package main

import (
	"fmt"
	"sync"
)

/*
两个共routine来回传递一个整数10次。当每个goroutine接收到整数时打印。每次通过整数都增加。一旦整数等于10，立刻终止程序
*/

func main() {
	share := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		goroutine("Bill", share)
		wg.Done()
	}()
	go func() {
		goroutine("Joan", share)
		wg.Done()
	}()

	share <- 1
	wg.Wait()
}

func goroutine(name string, share chan int) {
	for {
		value, ok := <-share
		if !ok {
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		fmt.Printf("Goroutine %s Inc %d\n", name, value)
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		share <- (value + 1)
	}
}
