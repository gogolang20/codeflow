package main

import (
	"fmt"
)

func main() {
	// map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	// 增加一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	// 如何动态增加 slice 的 append 动态增加
	// 先定义一个monsters信息
	newMonter := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonter)
	fmt.Println(newMonter)
	fmt.Println(monsters)
}
