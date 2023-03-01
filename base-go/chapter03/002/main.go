package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换方法一 fmt.Sprintf
	var num1 int = 98
	var num2 float32 = 25.314
	var b bool = true
	var char byte = 's'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	// 字符串转换方法二 strconv 函数
	var num3 int = 98
	var num4 float64 = 25.314
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	// 说明 "f" 代表转成的一种格式  10表示小数保留10位 64表示float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type is %T,str=%q\n", str, str)

	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type is %T,str=%q\n", str, str)
}
