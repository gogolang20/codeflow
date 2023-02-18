package main

import (
	"fmt"
)

func main() {
	/*
		养鸡场问题 6只鸡的体重  总体重和平均值
		使用数组解决
	*/

	// 声明一个数组
	var hens [6]float64
	// 给数组的每隔元素赋值
	hens[0] = 3.0 // 数组的第一个元素
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	// 遍历数组 求总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 求出平均体重
	avgWeight := fmt.Sprintf("%.3f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)

}
