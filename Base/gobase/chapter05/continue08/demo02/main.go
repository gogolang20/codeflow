package main

import (
	"fmt"
)

func main() {
	// 某人有100000元 没经过一个路口 需要缴费
	// 现金>50000 每次交5%
	// 现金<=50000 每次交1000
	// 计算可以经过多少次路口 使用for break 完成

	// 计算路过路口次数
	var count int
	var money float64 = 100000.0
	for {
		if money > 50000 {
			money *= 0.95
			count++
		} else if money > 1000 {
			money -= 1000
			count++
		} else {
			break
		}
	}
	fmt.Printf("能通过路口的次数：%v\n", count)
}
