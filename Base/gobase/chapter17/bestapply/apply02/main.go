package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int
	Score float64
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expext struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白象精")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	val.Elem().Method(0).Call(nil)
}

func main() {
	var a Monster = Monster{
		Name:  "黄狮子",
		Age:   200,
		Score: 92.8,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json result=", string(result))

	TestStruct(&a)
	fmt.Println(a)
}
