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

type Usb2 interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作!")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作!")
}

type Camera struct {
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
func (c Computer) Working(usb Usb2) {

	// 通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {

	// 测试
	// 创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点
	// computer 是绑定Computer结构体的方法
	// Computer 调用了Usb接口中 的方法
	// 至于方法的实现 ，取决与调用phone 中的写法
	computer.Working(phone)
	computer.Working(camera)

}
