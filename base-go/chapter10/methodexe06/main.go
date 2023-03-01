package main

import (
	"fmt"
)

// Circle 的j结构体
type Circle struct {
	radius float64
}

// 绑定Circle 结构体的 area 方法  返回一个值
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 为了提高效率 通常通过地址传递
func (c *Circle) area2() float64 {
	return 3.14 * (*c).radius * (*c).radius
	// return 3.14 * c.radius * c.radius 等价于上面的
}

func main() {

	// //声明一个Circle 结构体变量
	// var c Circle
	// //指定一个半径的值
	// c.radius = 4.0
	// //将字段传入方法  值赋给res
	// res := c.area()
	// fmt.Println("面积是:", res)

	// var c = Circle{radius: 5.0}
	var c Circle
	c.radius = 5.0

	res1 := c.area()
	res2 := (&c).area2()
	// 编译器底层优化  res := c.area2()
	fmt.Println("面积是:", res1)
	fmt.Println("面积是:", res2)

}
