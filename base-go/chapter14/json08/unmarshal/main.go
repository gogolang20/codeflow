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

func unmarshalStruct() {
	// 模拟演示 实际开发中会有程序或网络传入数据
	str := "{\"Age\":500,\"Birthday\":\"1011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("monster after unmarshal=%v monster.Age=%v\n", monster, monster.Age)
}

func testMap() string {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿!!!"
	a["age"] = 30
	a["address"] = "火云洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("json err=%v\n", err)
	}
	// fmt.Printf("a map after json =%v\n", string(data))
	return string(data)
}

func unmarshalMap() {
	// str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"
	str := testMap()

	var a map[string]interface{}
	// 反序列化不需要make

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("a after unmarshal=%v\n", a)

}

func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":3,\"name\":\"jack\"}," +
		"{\"address\":\"上海\",\"age\":5,\"name\":\"tom\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("slice after unmarshal=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
