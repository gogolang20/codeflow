package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"` // 反射机制
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1011-11-11",
		Sal:      8000.00,
		Skill:    "牛魔拳",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("json err=%v\n", err)
	}
	fmt.Printf("monster after json =%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("json err=%v\n", err)
	}
	fmt.Printf("a map after json =%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 3
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 5
	m2["address"] = "上海"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("json err=%v\n", err)
	}
	fmt.Printf("slice after json =%v\n", string(data))
}

func main() {

	// 序列化 切片 结构体 map
	testStruct()
	fmt.Println()
	testMap()
	fmt.Println()
	testSlice()
}
