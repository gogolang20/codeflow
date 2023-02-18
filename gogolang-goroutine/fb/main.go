package main

import (
	"fmt"
)

// goroutine 实现斐波那契

func Foo(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	var c = make(chan int)
	var quit = make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	Foo(c, quit)
}
