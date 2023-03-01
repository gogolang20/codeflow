package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("ok")
}

func main() {
	go func() {
		// 每秒调用一次 proc() 函数
		// 程序不能退出
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {}
}
