package main

import "fmt"

// channel 有序通信
func main() {
	unBuffered()
	buffered()
	closed()
}

// 无缓冲的channel：接收在发送之前
// 协程中 channel 要等到 a 赋值完毕后, 才能接收到 解除阻塞的信号
func unBuffered() {
	c := make(chan struct{})
	var a string
	go func() {
		a = "hello, world unBuffered"
		<-c
	}()
	c <- struct{}{}
	fmt.Println(a)
}

// 带缓冲的channel：发送在接收之前
// 给 a 赋值完成, 可以发送信号, 协程外 channel 解除阻塞
func buffered() {
	c := make(chan struct{}, 10)
	var a string
	go func() {
		a = "hello, world buffered"
		c <- struct{}{}
	}()
	<-c
	fmt.Println(a)
}

// 关闭在接收之前 (如有缓冲)
// 协程中先关闭 channel, 然后协程外 channel 会弹出默认值, 解除阻塞
func closed() {
	c := make(chan struct{}, 1)
	var a string
	go func() {
		a = "hello, world closed"
		close(c)
	}()
	<-c
	fmt.Println(a)
}
