package main

import "fmt"

// channel 有序通信
func main() {
	unBuffered()
	buffered()
	closed()
}

// 无缓冲的channel：接收在发送之前
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

// 关闭在接收之前（如有缓冲）
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
