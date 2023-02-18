package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	for i := 0; i <= 10; i++ {
		go func(i int) {
			for {
				// 希望程序能等当前这个周期休眠完 再优雅退出
				time.Sleep(time.Duration(i) * time.Second)
			}
		}(i)
	}
	fmt.Println(<-sig)
}
