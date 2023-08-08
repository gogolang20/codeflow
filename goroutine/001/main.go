package main

import (
	"net/http"
	"sync"
)

// 题目：请将函数foo补充完整，以n的并发访问完一个url列表，并发实现方式不限。（对url的访问直接以http.Get(url) 来表示即可）

func main() {
	n := 5
	urls := []string{"https://...", "..."} // 一个有很多URL的slice， 10000w
	foo(urls, n)
}

func foo(urls []string, n int) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 5)

	wg.Add(len(urls))

	for _, url := range urls {
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			http.Get(url)
			<-ch
		}()
	}
	wg.Wait()
}

// var wg = sync.WaitGroup{}

/*
有缓冲的方式控制goroutine并发数
*/
// func main() {
// 	// user request
// 	requestCount := 10
// 	fmt.Println("goroutine_num", runtime.NumGoroutine())

// 	// length of ch is the largest conurrency
// 	ch := make(chan bool, 3)
// 	for i := 0; i < requestCount; i++ {
// 		wg.Add(1)
// 		ch <- true
// 		go Read(ch, i)
// 	}

// 	wg.Wait()
// }

// func Read(ch chan bool, i int) {
// 	fmt.Printf("goroutine_num: %d, go func: %d\n", runtime.NumGoroutine(), i)
// 	<-ch
// 	wg.Done()
// }

/*
无缓冲的方式控制goroutine并发数
*/
// func main() {
// 	// user request
// 	requestCount := 10
// 	fmt.Println("goroutine_num", runtime.NumGoroutine())
//
// 	// length of ch is the largest conurrency
// 	ch := make(chan bool)
// 	for i := 0; i < 3; i++ {
// 		go Read(ch, i)
// 	}
//
// 	for i := 0; i < requestCount; i++ {
// 		wg.Add(1)
// 		ch <- true
// 	}
//
// 	wg.Wait()
// }
//
// func Read(ch chan bool, i int) {
// 	for _ = range ch {
// 		fmt.Printf("goroutine_num: %d, go func: %d\n", runtime.NumGoroutine(), i)
// 		wg.Done()
// 	}
// }
