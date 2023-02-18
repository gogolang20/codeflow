package main

import (
	"fmt"
)

func main() {
	// string的底层是一个byte数组
	str := "hello@atguigu中国"

	// 获取切片
	slice := str[6:]
	fmt.Println("slice=", slice)

	// string 是不可变的 不可以通过以下方式修改
	// slice[7] = "6"
	// fmt.Println("slice=", slice)
	// fmt.Println("str=", str)

	// 修改string字符串 可以先转成[]byte或[]rune 后在转成string
	// 修改str 的字符串
	arr1 := []byte(str) // 转换成[]byte切片 不能处理中文
	// byte按字节处理  一个汉字是三个字节
	arr1[0] = 'z' // 修改字符串
	str = string(arr1)
	fmt.Println("arr1=", arr1) // [122 101 108 108 111 64 97 116 103 117 105 103 117 228 184 173 229 155 189]
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(string(arr1[i]))
	// }
	fmt.Println("str=", str)

	// 使用[]rune切片修改中文
	arr2 := []rune(str)
	arr2[0] = '福' // 修改字符串
	fmt.Printf("arr2= %c \n", arr2)
	str = string(arr2)
	fmt.Println("str=", str)
}
