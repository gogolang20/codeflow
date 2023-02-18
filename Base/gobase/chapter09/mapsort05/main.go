package main

import (
	"fmt"
	"sort"
)

func main() {
	// map排序
	map1 := make(map[int]int, 5)
	map1[10] = 100
	map1[1] = 53
	map1[9] = 143
	map1[3] = 79
	map1[8] = 18

	fmt.Println(map1)
	// 如何按照map的key的顺序进行排序输出
	// 先将map的 key放到切片中
	// 对切片排序
	// 遍历切片 按照key来输出map的值

	// 遍历切片
	var keys []int
	for k, v := range map1 {
		keys = append(keys, k)
		fmt.Println(v)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	// 遍历切片 按照key来输出map的值
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
		// fmt.Printf("map1[%v]=%v\n", i, k)
	}
	// map 只能对 key 排序！！！

}
