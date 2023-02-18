package main

import (
	"fmt" // 前面加_ 表示暂时不使用 也不会报错
	"strconv"

	_ "unsafe"
)

func main() {
	// 演示string转成基本数据类型
	var str string = "true"
	var b bool
	// b ,_ = strconv.ParseBool(str)
	// 会返回两个值 只希望获取bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n", b, b)

	var str2 string = "1234569"
	var n1 int64
	var n2 int

	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type is %T n1=%v \n", n1, n1)
	fmt.Printf("n2 type is %T n2=%v \n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T n1=%v \n", f1, f1)
}
