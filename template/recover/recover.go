package main

import (
	"fmt"
	"time"
)

func main() {
	GOOOOO(func() {
		fmt.Println("hello goroutine!")
		panic("this is panic!")
	})

	time.Sleep(2 * time.Second)
}

func GOOOOO(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f()
	}()
}
