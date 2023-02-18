package main

import (
	"fmt"
)

// 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

// type Usb2 interface {
// 	Start()
// 	Stop()
// }

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作!")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作!")
}
func (p Phone) Call() {
	fmt.Println("手机 在打电话!")
}

type Camera struct {
	name string
}

// 让相机也实现 Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作!")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作!")
}

type Computer struct {
}

// 编写一个方法 working
// 方法接收一个Usb接口类型变量
// 所谓实现Usb接口 就是指实现了Usb接口声明所有办法
func (computer Computer) Working(usb Usb) {
	usb.Start()
	// 如果usb 是指向Phone结构体变量 则还需要调用Call方法
	// 类型断言 Attention！！！！！！！！！
	// phone1 这个名称可以该
	if phone1, ok := usb.(Phone); ok {
		phone1.Call()
	}
	usb.Stop()
}

func main() {

	// 定义一个Usb接口数组 可以
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	// fmt.Println(usbArr)

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

}
