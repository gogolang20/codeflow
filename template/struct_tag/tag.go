package main

import (
	"fmt"
	"reflect"
)

// 资源地址
// https://www.bilibili.com/video/BV1EF411h7Xq?p=17&spm_id_from=pageDriver

type UserInfo struct {
	Name     string `bilibili:"BILIBILI_NAME" abc:"name"`
	PublicWX string `bilibili:"BILIBILI_PUBLICWX"`
}

func PrintTag(ptr interface{}) {
	reType := reflect.TypeOf(ptr)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		panic("传入的参数不是结构体指针")
	}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		labtag := tag.Get("bilibili")

		fmt.Println("tag:", tag)
		fmt.Println("labtag:", labtag)
	}
}

func main() {
	userInfo := &UserInfo{
		Name:     "cloud study",
		PublicWX: "cncf_kubernetes",
	}
	PrintTag(userInfo)
}
