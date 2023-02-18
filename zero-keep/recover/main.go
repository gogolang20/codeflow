package main

import (
	"fmt"
	"net/http"
	"sync"
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

func S(urls []string, n int) {
	var wg sync.WaitGroup
	ch := make(chan string, n)

	wg.Add(len(urls))
	for _, url := range urls {
		ch <- url
		go func() {
			http.Get(url)
			<-ch
		}()
	}

	wg.Wait()
}
