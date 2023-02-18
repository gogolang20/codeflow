package main

import (
	"fmt"
	"unsafe"
)

// 变量使用的注意事项
func main() {

	// 该区域的数据值可以在同一类型范围内不断变化
	// var i int = 10
	// i := 10
	// i = 30
	// i = 50
	// fmt.Println("i=",i)
	// 变量在同一个作用域(函数或代码块)内不可以重名
	// i := 99 是不可以的 上面已有
	// var i int = 38 也是不可以的

	var i = 1
	var j = 2
	var r = i + j // +表示相加
	fmt.Println("r=", r)

	var str1 = "hello "
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res=", res)

	var n1 = 100
	// 查询n1 的数据类型
	fmt.Printf("n1 的类型是 %T\n", n1)

	// 如何在程序查看某个变量的占用字节大小和数据类型
	var n2 int64 = 10
	// unsafe.Sizeof()  查询n2占用的字节数
	fmt.Printf("n2 的类型是 %T  n2占用的字节数是 %d", n2, unsafe.Sizeof(n2))
}
