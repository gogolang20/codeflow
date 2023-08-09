package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//常规写法
func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	finishedCh := make(chan struct{})

	go func(ctx context.Context, finishedCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(ctx, finishedCh)

	<-sig
	cancel()
	<-finishedCh
	fmt.Println("finished")
}
