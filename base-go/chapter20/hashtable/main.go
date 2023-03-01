package main

import (
	"fmt"
	"os"
)

// 单链表 散列链表
// 如何添加删除信息

// 定义 Emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 方法待定
func (emp *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员%d\n", emp.Id%7, emp.Id)
}

// 定义 EmpLink
// 这里的是不带表头的
type EmpLink struct {
	Head *Emp
}

// 方法待定
// 添加员工的方法 (有一个漏洞)
func (empLink *EmpLink) Insert(emp *Emp) {
	cur := empLink.Head // 辅助指针
	var pre *Emp = nil  // 辅助指针
	// 如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		empLink.Head = emp // 完成
		return
	}
	// 如果不是一个空链表 ！！！  给emp找到对应的位置并插入
	// 思路 让cur 和emp 比较，然后 pre 保存保持在cur 前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置了
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时，我们看下是否将emp 添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

// 显示当前列表的信息
func (empLink *EmpLink) ShowLink(no int) {
	if empLink.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	// 遍历当前的链表，并显示数据
	cur := empLink.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() // 换行处理
}

// 查询方法,根据id查找对应的雇员，如果没有就返回 nil
func (empLink *EmpLink) FindById(id int) *Emp {
	cur := empLink.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 定义 HashTable 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给HashTable 编写 Insert 雇员的方法
func (ht *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定添加到哪个 EmpLink
	linkNo := ht.HashFun(emp.Id)
	// 使用对应的链表添加
	ht.LinkArr[linkNo].Insert(emp) //
}

// 编写一个用于散列的方法
func (ht *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对应链表的下标
}

// 编写一个方法，完成查找
func (ht *HashTable) FindById(id int) *Emp {
	// 使用散列函数，确定雇员在哪个 EmpLink
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindById(id)
}

// 编写方法 显示 HashTable 的所有雇员
func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("系统菜单")
		fmt.Println("input ")
		fmt.Println("show")
		fmt.Println("find")
		fmt.Println("exit")
		fmt.Println("请输入选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id:")
			fmt.Scanln(&id)
			fmt.Println("输入雇员姓名:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("输入雇员id:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在", id)
			} else {
				// 编写一个方法 显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}

}
