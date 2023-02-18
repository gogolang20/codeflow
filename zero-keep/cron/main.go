package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	// 创建任务调度器实例
	c := cron.New()

	// 注册任务到调度器，注册的任务都是异步执行的。
	// c.AddFunc("0 30 * * * *", func() {
	// 	fmt.Println("every hour on the half hour run...")
	// })
	addFunc, err := c.AddFunc("@every 2m", func() {
		fmt.Println("every hour run...")
	})
	if err != nil {
		log.Fatal("add func err: ", err)
	}

	fmt.Println("from addfunc: ", addFunc)

	// 启动计划任务
	c.Start()
}
