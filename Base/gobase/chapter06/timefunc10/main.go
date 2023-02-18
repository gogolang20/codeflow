package main

import (
	"fmt"
	"time"
)

func main() {
	// 日期和时间相关 函数和方法的使用
	// 1: 获取当前时间
	now := time.Now()
	// time的值是多少 类型是什么
	// now 的类型是 time.Time
	fmt.Printf("now=%v now type=%T\n", now, now)

	// 2:通过now可以获取到年月日 时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期和时间 两种方式
	// 第一种 传统方式 SPrintf 或 Printf
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())

	dataStr := fmt.Sprintf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dataStr=%v\n", dataStr)

	// 第二种 now.Format()
	// 中间间隔可以改变 时间的数字不可以改变
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// //每隔1s中打印一个数字 到100停止
	// //休眠函数
	// //每隔0.1秒打印一个数字
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	//休眠
	// 	//time.Sleep(time.Second)i
	// 	time.Sleep(100 * time.Millisecond)
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	// Uxin 和 UxixNano 方法使用案例
	fmt.Printf("Unix时间戳=%v UnixNano时间戳=%v\n", now.Unix(), now.UnixNano())

}
