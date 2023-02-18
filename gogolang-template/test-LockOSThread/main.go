package main

import (
	"fmt"
	"net"
	"runtime/pprof"
	"sync"
)

// 资料地址
// https://www.bilibili.com/video/BV1EF411h7Xq?p=18&spm_id_from=333.788.top_right_bar_window_history.content.click

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	fmt.Printf("Before 线程数：%d\n", threadProfile.Count())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			// runtime.LockOSThread()
			for j := 0; j < 100; j++ {
				_, err := net.LookupHost("www.baidu.com")
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Done()
	fmt.Printf("After 线程数：%d\n", threadProfile.Count())
}

// GODEBUG=netdns=go go run main.go
// GODEBUG=netdns=cgo go run main.go
