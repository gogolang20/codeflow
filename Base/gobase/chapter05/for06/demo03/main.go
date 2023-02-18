package main

import (
	"fmt"
	_ "math"
)

func main() {

	// 案例1
	// 三个班成绩情况，每个班5名学生 求各个班的平均分和所有班级的平均分 (成绩从键盘输入)
	// 班级 学生 总分数

	// 先做第一个班级 计算分数

	// 把代码做活
	// 定义班级的个数 和人数

	// 案例2
	// 统计三个班级的及格人数
	// 添加一个通过人数的变量 passCount

	var classNum int = 2
	var stuNum int = 4
	var totalSum float64 = 0.0
	var passCount int
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d个班 第%d个学生分数：\n", j, i)
			fmt.Scanln(&score)
			// 累计总分
			sum += score
			// 判断是否通过考试
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d个班级的平均分=%v \n", j, sum/float64(stuNum))
		// 算出各个班级的总分
		totalSum += sum
	}

	fmt.Printf("各个班级的总成绩%v 所有班级的平均分=%v \n",
		totalSum, totalSum/(float64(stuNum)*float64(classNum)))

	fmt.Printf("通过考试数量%v", passCount)
}
