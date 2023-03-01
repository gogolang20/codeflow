package main

import (
	"fmt"
	"sync"
)

/*
goroutine 交替打印奇偶数
-----------------------------------------  方式一   -----------------------------------------
*/

// var (
//	wg = sync.WaitGroup{}
//	ch = make(chan struct{})
// )
//
// func main() {
//	wg.Add(2)
//
//	go worker01()
//	go worker02()
//
//	wg.Wait()
// }
//
// func worker01() {
//	for i := 1; i < 100; i++ {
//		ch <- struct{}{}
//		if i&1 == 0 {
//			fmt.Println("this is odd ", i)
//		}
//	}
//	wg.Done()
// }
//
// func worker02() {
//	for i := 1; i < 100; i++ {
//		<-ch
//		if i&1 != 0 {
//			fmt.Println("this is worker02 ", i)
//
//		}
//	}
//	wg.Done()
// }

/*
-----------------------------------------  方式二   -----------------------------------------
*/

var (
	MAX_NUM = 200

	ch       = make(chan struct{})
	flagChan = make(chan int, MAX_NUM)
	wg       = sync.WaitGroup{}
)

func main() {
	for i := 1; i <= MAX_NUM; i++ {
		flagChan <- i
	}
	wg.Add(2)
	go work01()
	go work02()
	wg.Wait()
	// time.Sleep(2 * time.Second)
}

func work01() {
	for {
		// block
		ch <- struct{}{}
		x := <-flagChan
		if x == MAX_NUM || x == MAX_NUM-1 {
			wg.Done()
		}
		fmt.Println(x)
	}
}

func work02() {
	for {
		// read
		<-ch
		y := <-flagChan
		if y == MAX_NUM || y == MAX_NUM-1 {
			wg.Done()
		}
		fmt.Println(y)
	}
}
