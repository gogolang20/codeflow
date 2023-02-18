package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
蓄水池算法

解决的问题：

假设有一个源源吐出不同球的机器，

只有装下10个球的袋子，每一个吐出的球，要么放入袋子，要么永远扔掉

如何做到机器吐出每一个球之后，所有吐出的球都等概率被放进袋子里
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	test := 1000000
	poolSize := 17
	count := make([]int, poolSize+1)
	for i := 0; i < test; i++ {
		bag := make([]int, 10)
		bagi := 0
		for num := 1; num <= poolSize; num++ {
			if num <= 10 {
				bag[bagi] = num
				bagi++
			} else {                        // num > 10
				if rand.Intn(num+1) <= 10 { // 一定要把num球入袋子
					bagi = rand.Intn(10) // 随机选择bag中要丢弃的位置
					bag[bagi] = num
				}
			}
		}
		for _, num := range bag { // 统计每次进入bag的num
			count[num]++
		}
	}
	for i := 1; i <= poolSize; i++ {
		fmt.Println(count[i])
	}
}
