package main

import (
	"errors"
	"fmt"
	"os"
)

// 单向数组队列：先进先出
// 使用结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 一个数组
	front   int    // 队列首位
	rear    int    // 队列最后一个
}

// 添加数据到队列
func (queue *Queue) AddQueue(val int) (err error) {

	// 先判断队列已满
	if queue.rear == queue.maxSize-1 { // rear 是队列尾部，包含最后一个元素
		return errors.New("queue full")
	}
	queue.rear++ // rear 后移
	queue.array[queue.rear] = val
	return
}

// 从队列中取出数据
func (queue *Queue) GetQueue() (val int, err error) {
	// 判断队列是否为空
	if queue.rear == queue.front {
		return -1, errors.New("queue empty")
	}
	queue.front++
	val = queue.array[queue.front]
	return val, err
}

// 显示队列 找到队首，遍历到队尾
func (queue *Queue) ShowQueue() {
	fmt.Println("当前队列如下：")
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("array[%d] = %d\t", i, queue.array[i])
	}
	fmt.Println()
}

func main() {

	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1:输入add 表示添加数据到队列")
		fmt.Println("2:输入get 表示从队列获取数据")
		fmt.Println("3:输入show 表示显示队列")
		fmt.Println("4:输入exit 退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入您要如队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("AddQueue ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
