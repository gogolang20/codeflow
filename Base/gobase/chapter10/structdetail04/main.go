package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	// leftUp, rightDown Point  同下方写法一样
	leftUp    Point
	rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	// r1 有四个整数
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// r1.leftUp.x 地址=0xc0420540a0 r1.leftUp.y 地址=0xc0420540a8
	// r1.rightDown.x 地址=0xc0420540b0 r1.rightDown.y 地址=0xc0420540b8

	// r2有两个 *Point类型 两个*Point类型本身地址也是连续的
	// 他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	fmt.Printf("r2.leftUp 本身地址=%p  r2.rightDown 本身地址=%p\n", &r2.leftUp, &r2.rightDown)
	// r2.leftUp 本身地址=0xc04203c1c0  r2.rightDown 本身地址=0xc04203c1c8

	// 指向地址不一定连续
	fmt.Printf("r2.leftUp 指向地址=%p  r2.rightDown 指向地址=%p", r2.leftUp, r2.rightDown)
}
