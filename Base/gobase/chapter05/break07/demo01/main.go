package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 为了生成一个随机数 还需要给rand设置一个种子
	// time.Now().Unix()  调用 Now 函数的一个 unix方法  返回一个从1970年01.01 0时0分0秒到现在的秒数
	// fmt.Println(time.Now().Unix())
	// rand.Seed(time.Now().Unix())

	// 随机生成一个1~100整数
	// n := rand.Intn(100) + 1 //[0~100)
	// fmt.Println(n)

	// 不停的随机 生成98时停止 记录生成次数
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1 // 1-100
		fmt.Println(" n= ", n)
		count++
		if n == 98 {
			break
		}
	}
	fmt.Println("生成98 一共使用了", count)

	fmt.Println("-----------------------")
	var counts [10]int
	for i := 0; i < 100000; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) + 1
		counts[n-1]++
	}
	for i, v := range counts {
		fmt.Printf("%v 出现的次数=%v\n", i+1, v)
	}

	fmt.Println("-----------------------")

	// break 演示
lable2:
	for i := 0; i <= 4; i++ {
		// lable1: //设置了一个标签
		for j := 0; j <= 10; j++ {
			if j == 2 {
				// break //break  默认会跳出最近的for循环
				break lable2 // j=0 j=1
			}
			fmt.Println("j=", j)
		}
	}
}
