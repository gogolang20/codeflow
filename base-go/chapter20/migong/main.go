package main

import "fmt"

// 编写一个函数，完成老鼠找路
func SetWay(myMap *[8][7]int, i int, j int) bool {
	// 分析什么情况下找出出路
	if myMap[6][5] == 2 {
		return true
	} else {
		// 说明要继续找
		if myMap[i][j] == 0 { // 可以探测
			// 假设这个点可以通 ，但是需要探测上下左右
			myMap[i][j] = 2
			if SetWay(myMap, i-1, j) {
				return true
			} else if SetWay(myMap, i+1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else { // 当前点是死路
				myMap[i][j] = 3
				return false
			}
		} else { // 说明这个点不可以探测
			return false
		}
	}
}

func main() {
	// 先创建一个二位始数组 模拟迷宫
	// 规则
	// 1 如果元素的值为1 就是墙
	// 2 如果元素的值为0 说明还没有走过的点
	// 3 如果元素的值为2 表示是一个通路
	// 4 如果元素的值为3 表示走过的路，但是不通
	var myMap [8][7]int
	// 先设置墙  最上和最下
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	// 设置左右面的墙
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1

	// for i1, v1 := range myMap {
	// 	fmt.Println()
	// 	for _, v2 := range v1 {
	// 		fmt.Print(myMap[i1][v2])
	// 	}
	// }
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕")

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
