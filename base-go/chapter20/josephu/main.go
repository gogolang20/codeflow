package main

import (
	"fmt"
)

type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下个小孩的指针
}

// 编写一个函数 构成单向的环形链表
// num 表示小孩的个数
// *Boy 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  // 空结点
	curBoy := &Boy{} // 空结点
	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 分析构成循环链表，需要一个辅助指针【帮忙的】
		// 1 因为第一个小孩比较特殊
		if i == 1 { // 第一个小孩
			first = boy // 不能动
			curBoy = boy
			curBoy.Next = first // 形成环形
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

// 显示单向的环形链表
func ShowBoy(first *Boy) {
	// 如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩")
		return
	}
	// 先创建一个指针帮助遍历 【至少有一个小孩】
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		// 退出条件
		if curBoy.Next == first {
			break
		}
		// curBoy 移动到下一个
		curBoy = curBoy.Next
	}
}

// 分析思路
// 1 编写一个函数 playGame
func PlayGame(first *Boy, startNo int, countNum int) {
	// 1 空的链表单独出来
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	// 流一个 判断 startNo <= 小孩的总数
	// 2 需要定义辅助指针 帮助我们删除小孩
	tail := first
	// 3 让tail 指向环形链表的最后一个小孩 非常重要！！！
	// 因为 tail 在删除小孩时会使用到
	for {
		if tail.Next == first { // 说明tail 到了最后的小孩
			break
		}
		tail = tail.Next
	}
	// 4 让 first 移动到 startNo， 后面删除小孩 就以first 为准
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	// 5 开始数 countNum 然后就删除first 指向的小孩
	for {
		// 开始数 countNum - 1 次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈\n", first.No)
		// 删除 first 指向的小孩
		first = first.Next
		tail.Next = first
		// 判断如果 tail == first 说明只有一个小孩了
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈\n", first.No)
}

func main() {
	first := AddBoy(5)
	// 显示
	ShowBoy(first)
	PlayGame(first, 2, 3)

}
