package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 声明一个结构体
type Hero struct {
	Name string
	Age  int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是觉得使用什么标准进行排序
// 从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	// return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]//先吧i的值保留
	// hs[i] = hs[j]//用j的值将i的值覆盖
	// hs[j] = temp//将保留的i的值交换
	hs[i], hs[j] = hs[j], hs[i] // 与上面方法等价
}

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Student的切片 按Score 从大到小排列
type StuSlice []Student

// 绑定三个方法
// 以便使用sort.Sort的 Interface 排序接口
func (ss StuSlice) Len() int {
	return len(ss)
}

func (ss StuSlice) Less(i, j int) bool {
	return ss[i].Score > ss[j].Score
}

func (ss StuSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {

	// 数组/切片排序
	// 冒泡排序
	// 系统提供的方法 sort.Ints
	var intSlice = []int{0, -1, 99, 20, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 引出对结构体切片进行排序的需求
	// 冒泡排序
	// 系统提供的方法 Interface 接口 实现Len Less Swap
	// Len 是元素的个数 Less 是比较大小 Swap 是交换比较的值

	// 测试
	// 声明一个 HeroSlice 类型的变量 名叫 heros
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		// Hero结构体 添加内容
		hero := Hero{
			Name: fmt.Sprintf("big英雄_%d:", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	// 查看排序结果
	for _, v := range heroes {
		fmt.Println(v)
	}

	fmt.Println("排序后")
	// 调用 sort.Sort
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}

	// 练习部分
	var students StuSlice
	for i := 0; i < 20; i++ {
		stu := Student{
			Name:  fmt.Sprintf("小学生：%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Intn(100),
		}
		students = append(students, stu)
	}

	sort.Sort(students)
	for _, v := range students {
		fmt.Println(v)
	}

	// 交换
	n1 := 10
	n2 := 20
	n1, n2 = n2, n1
	fmt.Println("n1=", n1, "n2=", n2)

}
