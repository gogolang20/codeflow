package main

import (
	"fmt"
	"sync"
)

// goroutine 交替打印奇偶数

/*
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

const (
	MAX_NUM = 200
)

func main() {
	var (
		ch       = make(chan struct{})
		flagChan = make(chan int, MAX_NUM)
		wg       = sync.WaitGroup{}
	)

	for i := 1; i <= MAX_NUM; i++ {
		flagChan <- i
	}
	wg.Add(2)
	go work01(&wg, ch, flagChan)
	go work02(&wg, ch, flagChan)
	wg.Wait()
	// time.Sleep(2 * time.Second)
}

func work01(wg *sync.WaitGroup, ch chan struct{}, flagChan chan int) {
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

func work02(wg *sync.WaitGroup, ch chan struct{}, flagChan chan int) {
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
