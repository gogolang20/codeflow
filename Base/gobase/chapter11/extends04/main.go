package main

import (
	"fmt"
)

/*
继承
先抽取共有的属性(字段)
*/
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 将共有的方法绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n",
		stu.Name, stu.Age, stu.Score)
}

// 将共有的方法绑定到 *Student
func (stu *Student) SetScore(score float64) {
	// 条件判断
	stu.Score = score
}

// 添加一个方法
func (stu *Student) GetSum(n1 float64, n2 float64) float64 {
	return n1 + n2
}

type Pupil struct {
	Student // 嵌入了Student 的匿名结构体
}

// 特有的方法 保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}

// 如果还有其他学生考试  代码会重复导致冗余
// 不利于代码维护
// 不利于功能扩展
type Graduate struct {
	Student // 嵌入了Student 的匿名结构体
}

// 特有的方法 保留
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main() {
	// 测试  继承
	// 使用方法发生变化

	// 调用方式一
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 9
	pupil.testing()
	pupil.Student.SetScore(89)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.GetSum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "mary"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.Student.SetScore(89)
	graduate.Student.ShowInfo()
	fmt.Println("res=", graduate.Student.GetSum(6, 99))

}
