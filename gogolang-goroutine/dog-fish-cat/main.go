package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 资料地址：没有次数限制
// https://www.bilibili.com/video/BV1EF411h7Xq?p=15
// func cat(fishCh, catCh chan struct{}, wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func() {
// 		for {
// 			fmt.Println("cat")
// 			catCh <- struct{}{}
// 			<-fishCh //顺序不能修改
// 		}
// 		wg.Done()
// 	}()
// }
// func dog(catCh, dogCh chan struct{}, wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func() {
// 		for {
// 			<-catCh
// 			fmt.Println("dog")
// 			dogCh <- struct{}{}
// 		}
// 		wg.Done()
// 	}()
// }
// func fish(dogCh, fishCh chan struct{}, wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func() {
// 		for {
// 			<-dogCh
// 			fmt.Println("fish")
// 			time.Sleep(time.Second) //等待一秒钟
// 			fishCh <- struct{}{}
// 		}
// 		wg.Done()
// 	}()
// }
// func main() {
// 	catCh := make(chan struct{})
// 	dogCh := make(chan struct{})
// 	fishCh := make(chan struct{})
// 	wg := sync.WaitGroup{}

// 	cat(fishCh, catCh, &wg)
// 	dog(catCh, dogCh, &wg)
// 	fish(dogCh, fishCh, &wg)
// 	wg.Wait()
// }

// 控制打印次数 100次
func main() {
	var wg sync.WaitGroup

	var dogcounter uint64
	var fishcounter uint64
	var catcounter uint64
	var dogch = make(chan struct{}, 1)
	var fishch = make(chan struct{}, 1)
	var catch = make(chan struct{}, 1)

	wg.Add(3)
	dogch <- struct{}{}
	go dog(&wg, dogcounter, dogch, fishch)
	go fish(&wg, fishcounter, fishch, catch)
	go cat(&wg, catcounter, catch, dogch)

	wg.Wait()
}

func dog(wg *sync.WaitGroup, counter uint64, dogch, fishch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
		}
		<-dogch
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishch <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, counter uint64, fishch, catch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
		}
		<-fishch
		fmt.Println("Fish")
		atomic.AddUint64(&counter, 1)
		catch <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter uint64, catch, dogch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
		}
		<-catch
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogch <- struct{}{}
	}
}
