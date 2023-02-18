package main

import (
	"fmt"
)

func main() {
	var adress string = "北京长城 110 hello world"
	fmt.Println(adress)

	// 字符串一旦赋值，字符串就不能修改了 在go中字符串是不可变的
	// var str = "hello"
	// str[0] = 'a' //这里不能修改str的内容

	// ""
	str2 := "abc\nabc"
	fmt.Println(str2)
	// 使用反引号 `` esc底下的按键 可以以字符串的原生形式输出

	// 字符串的拼接方式
	var str = "hello" + "world"
	str += " haha"
	fmt.Println(str)

	// 当一个拼接的操作很长的时候 拼接的 + 需要留着换行的上方

	// 变量的默认值
	var a int
	var b float32
	var c float64
	var isMarried bool
	var name string
	// %v代表按照变量的原值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v", a, b, c, isMarried, name)

}
