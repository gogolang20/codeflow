package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1:创建一个数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 表示黑子
	chessMap[2][3] = 2 // 表示蓝子
	// 2:输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 转成一个稀疏数组
	// 使用切片 因为不知道有多少个需要保存
	// 变量chessMap 如果发现有个元素的值不为0  就创建一个node结构体
	// 将其放入到对应的切片中

	// 创建一个切片 是ValNode 结构体
	var sparseArr []ValNode

	// 标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行和列，默认值
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	// 将要储存的值加入到切片中
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个  ValNode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				// 将值存入创建的 sparseArr 切片中
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是::::::::::::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n",
			i, valNode.row, valNode.col, valNode.val)
	}

	// 将这个稀疏输入存盘

	// 如何恢复原始的数组

	// 1:打开这个文件  恢复原始数组

	// 2:这里使用稀疏数组恢复

	// 先 创建一个原始的数组
	var chessMap2 [11][11]int

	// 遍历 sparseArr  [遍历文件的每一行]
	for i, valNode := range sparseArr {
		if i != 0 { // 跳过第一行记录的值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	// 遍历数组打印
	fmt.Println("恢复后的原始数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
