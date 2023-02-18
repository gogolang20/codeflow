package main

import "fmt"

// 资源地址
// https://www.bilibili.com/video/BV1EF411h7Xq?p=6

/*
#编译
go build -gcflags=all="-N -l" -ldflags='compressdwarf=false' -o eface eface.go
#反编译
go tool compile -S eface.go
gdb -tui ./eface
$ l
$ b eface.go:2 #断点位置
$ r
$ n
$ n
$ pt y #打印变量y
$ q #退出
*/

// nil 是否相等
func main() {
	var x *int = nil
	var y interface{} = x
	fmt.Println(x == y)
	fmt.Println(x == nil)
	fmt.Println(y == nil)

	var z interface{} = nil
	fmt.Println(z == nil)
}
