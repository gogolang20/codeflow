package main

import (
	"fmt"
)

func main() {
	// map的增 删 改 查
	// 第二种 (建议)
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "南京"
	fmt.Println(cities)

	// 因为no3 这个key已经存在 相当与修改
	cities["no3"] = "南京！！"
	fmt.Println(cities)

	// 演示删除 内置函数delete
	delete(cities, "no1")
	fmt.Println(cities)
	delete(cities, "no4")
	fmt.Println(cities)

	// 演示map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}
	// 一次性  map 删除
	// 方式一：遍历逐一删除
	// 方式二：直接make一个新的空间  cities = make(map[string]string)
	cities = make(map[string]string) // 重新make
	fmt.Println(cities)
}
