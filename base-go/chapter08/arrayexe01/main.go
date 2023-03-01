package main

import (
	"fmt"
)

func main() {
	// 三个班级 每隔班级5名学生
	var scores [3][5]float64
	// 循环输入成绩
	// len(scores)是 [3]的长度
	for i := 0; i < len(scores); i++ {
		// len(scores[i])是 [5]的长度
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班级 第%d同学的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	// fmt.Println(scores)//测试没有问题

	// 遍历输出成绩后的二维数组
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分数是%v,平均分是%v\n", i+1,
			sum, sum/float64(len(scores[i])))
	}

	fmt.Printf("所有班级的总分数是%v,所有班级平均分是%v\n",
		totalSum, totalSum/15.0)

}
