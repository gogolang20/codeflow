package main

import "fmt" // fmt包中提供格式化 输出 输入的函数

func main() {
	// 演示转义字符的使用
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	fmt.Println("tom\\jack")
	fmt.Println("tom talk to jack\"byebye\"")
	// \r 回车 从当前行的最前面开始输入 覆盖掉以前的内容
	fmt.Println("tom talk to jack\"byebye\" \rbody")

	fmt.Println("姓名\t年龄\t籍贯\t地址\njhon\t12\t河北\t北京")
}
