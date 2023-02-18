package main

import (
	"container/list"
	"fmt"
)

/*
1 如何用栈结构实现队列结构？？？
    准备 push栈 && pop栈
	1 加入数据时，pop栈是否空
    2 如果pop栈数据没有拿完，不可以导入数据

2 如何用队列结构实现栈结构？？？经典的对列，非redis一样的双端队列
    设置两个队列A B
    加入数据：进入A队列中
	弹出数据：将 A 的数据导入 B中，只留最后一个，返回给用户
	再次加入数据：将数据放入有数据的队列中

前提：
图的宽度优先遍历 一般使用队列实现
图的深度优先遍历 一般使用栈实现
*/

func main() {
	queue := list.New() // 反向操作为队列
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	fmt.Println("Queue Peek:", queue.Front().Value.(int)) // 每次取值方向

	stack := list.New() // 同向操作为栈
	stack.PushFront(1)
	stack.PushFront(2)
	stack.PushFront(3)
	fmt.Println("Stack Peek:", stack.Front().Value.(int)) // 每次取值方向
}

/*
数组实现队列
    思路一 预留一个位置，不存放数据

    思路二 begin，end，添加 size 记录队列是否满了   #简单实用
		begin = 0 //push进元素，begin++
		end = 0   //pop出元素，end++
		size = 0
*/
