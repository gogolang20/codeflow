package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//大神 Rob Pike
func main() {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				// 结束后 通过ch 通知主goroutine
				fmt.Println("stopped")
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	// ch 作为一个 channel，传递给子 goroutine， 待其结束后从中返回
	ch := make(chan struct{})
	stopCh <- ch
	<-ch
	fmt.Println("finished")
}
